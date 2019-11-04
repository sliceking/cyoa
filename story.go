package cyoa

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
