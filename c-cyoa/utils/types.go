package utils


type Option struct {
	Text string `json:"text"`
	Chapter string `json:"arc"`
}

// Struct tags give metadata about fields
type Chapter struct {
	Title   string   `json:"title"`
	Paragraphs   []string `json:"story"`
	Options []Option `json:"options"`
}

type Story map[string]Chapter

