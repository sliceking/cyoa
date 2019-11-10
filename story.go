package cyoa

import (
	"encoding/json"
	"html/template"
	"io"
	"log"
	"net/http"
	"strings"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").Parse(defaultHandlerTmpl))
}

var defaultHandlerTmpl = `
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <meta http-equiv="X-UA-Compatible" content="ie=edge">
  <title>Document</title>
</head>
<body>
  <h1>{{.Title}}</h1>

  {{range .Paragraphs}}
    <p>{{.}}</p>
  {{end}}

  <ul>
    {{range .Options}}
      <li><a href="{{.Chapter}}">{{.Text}}</a></li>
    {{end}}
  </ul>
</body>
</html>`

// HandlerOption is a config option that modifies the http handler
type HandlerOption func(h *handler)

// WithTemplate is a handler option that allows the customization of templates
func WithTemplate(t *template.Template) HandlerOption {
	return func(h *handler) {
		h.t = t
	}
}

// WithPathFunc is a handler option that alllows the customization of routes
func WithPathFunc(fn func(r *http.Request) string) HandlerOption {
	return func(h *handler) {
		h.pathFunc = fn
	}
}

func defaultPathFunc(r *http.Request) string {
	path := strings.TrimSpace(r.URL.Path)
	if path == "" || path == "/" {
		path = "/intro"
	}

	return path[1:]
}

// NewHandler creates a handler by taking in a story and handler options
func NewHandler(s Story, opts ...HandlerOption) http.Handler {
	h := handler{s, tpl, defaultPathFunc}
	for _, opt := range opts {
		opt(&h)
	}
	return h
}

type handler struct {
	s        Story
	t        *template.Template
	pathFunc func(r *http.Request) string
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := h.pathFunc(r)
	if chapter, ok := h.s[path]; ok {
		err := tpl.Execute(w, chapter)
		if err != nil {
			log.Printf("%v", err)
			http.Error(w, "Something went wrong", http.StatusInternalServerError)
		}
		return
	}
	http.Error(w, "Chapter not found", http.StatusNotFound)
}

// NewStory takes in a reader or file and returns a story
func NewStory(file io.Reader) (Story, error) {
	var adventure Story
	r := json.NewDecoder(file)
	err := r.Decode(&adventure)
	if err != nil {
		return nil, err
	}

	return adventure, nil
}

// Story is map comprised of chapters
type Story map[string]Chapter

// Chapter is a part of the story, options are links to different chapters
type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option `json:"options"`
}

// Option is where you can go to from the current chapter
type Option struct {
	Text    string `json:"text"`
	Chapter string `json:"arc"`
}
