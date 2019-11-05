package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/svwielga4/cyoa"
)

func main() {
	// create flag to accept
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
	fmt.Println(adventure)
}
