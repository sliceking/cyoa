package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

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

	h := cyoa.NewHandler(adventure, nil)
	fmt.Printf("starting the server on port: %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))
}
