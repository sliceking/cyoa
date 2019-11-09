package cyoa

import (
	"encoding/json"
	"html/template"
	"io"
	"log"
	"net/http"
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

// HandlerOption is a config
type HandlerOption func(h *handler)

func WithTemplate(t *template.Template) HandlerOption {
	return func(h *handler) {
		h.t = t
	}
}

func NewHandler(s Story, opts ...HandlerOption) http.Handler {
	h := handler{s, tpl}
	for _, opt := range opts {
		opt(&h)
	}
	return h
}

type handler struct {
	s Story
	t *template.Template
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "" || r.URL.Path == "/" {
		err := tpl.Execute(w, h.s["intro"])
		if err != nil {
			panic(err)
		}
		return
	}

	path := r.URL.Path[1:]
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

// Story is comprised of chapters
type Story map[string]Chapter

// Chapter is a part of the story
type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option `json:"options"`
}

// Option is where you can go from a chapter
type Option struct {
	Text    string `json:"text"`
	Chapter string `json:"arc"`
}
