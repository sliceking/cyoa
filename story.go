package cyoa

import (
	"encoding/json"
	"io"
)

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
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []Option `json:"options"`
}

// Option is where you can go from a chapter
type Option struct {
	Text    string `json:"text"`
	Chapter string `json:"arc"`
}
