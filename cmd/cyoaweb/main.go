package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/svwielga4/cyoa"
)

func main() {
	// create flag to accept
	port := flag.Int("port", 3000, "the port to start your CYOA app on")
	f := flag.String("file", "gopher.json", "The name of the file with a create your own adventure story")
	flag.Parse()

	// open file
	file, err := os.Open(*f)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// parse json
	adventure, err := cyoa.NewStory(file)
	if err != nil {
		panic(err)
	}

	// load a custom template
	tpl := template.Must(template.New("").Parse(customTpl))
	// create a handler with custom options
	h := cyoa.NewHandler(
		adventure,
		cyoa.WithTemplate(tpl),
		cyoa.WithPathFunc(pathFunc),
	)
	// create a mux to control routing a little better
	mux := http.NewServeMux()
	// all requests with `/story/` go to the custom handler
	mux.Handle("/story/", h)
	// all requests without go to a default handler
	mux.Handle("/", cyoa.NewHandler(adventure))
	fmt.Printf("starting the server on port: %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), mux))
}

// custom path function as an example
func pathFunc(r *http.Request) string {
	path := strings.TrimSpace(r.URL.Path)
	if path == "/story" || path == "/story/" {
		path = "/story/intro"
	}

	return path[len("/story/"):]
}

var customTpl = `
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
      <li><a href="/story/{{.Chapter}}">{{.Text}}</a></li>
    {{end}}
  </ul>
</body>
</html>`
