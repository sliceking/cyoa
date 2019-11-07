package cyoa

import (
	"encoding/json"
	"html/template"
	"io"
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

  <p>{{.Story}}</p>
  <ul>
    {{range .Options}}
      <li><a href="{{.Chapter}}">{{.Text}}</a></li>
    {{end}}
  </ul>
</body>
</html>`

func NewHandler(s Story) http.Handler {
	return handler{s}
}

type handler struct {
	s Story
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := tpl.Execute(w, h.s["intro"])
	if err != nil {
		panic(err)
	}
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
