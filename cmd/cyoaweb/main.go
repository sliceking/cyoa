package main

import (
	"encoding/json"
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

	r := json.NewDecoder(file)
	var adventure cyoa.Story
	err = r.Decode(&adventure)
	if err != nil {
		panic(err)
	}
	fmt.Println(adventure)
}
