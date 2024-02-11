package utils

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

func throwError(err error, message string) {
	if err != nil {
		log.Fatal(message)
	}
}

type HandlerOption func(h *handler)

// NOTE: Go's design better suits returning
// [data,error] tuple rather than throwing error
func ParseJSON(filename string) Story {
	file, err := os.Open("data/" + filename)

	throwError(err, "Error while opening file")

	d := json.NewDecoder(file)

	var story Story

	if err := d.Decode(&story); err != nil {
		throwError(err, "Failed decoding json")
	}
	// converts data in []Byte instead of remaining a struct
	// res, err := json.MarshalIndent(story, "", "	")
	// if err != nil {
	// 	throwError(err, "Failed indenting json")
	// }

	return story
}

// NOTE: Dave cheney's way of passing functional options
// https://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis
func WithTemplate(t *template.Template) HandlerOption {
	return func(h *handler) {
		h.t = t
	}
}

func WithPathFn(fn func(r *http.Request) string) HandlerOption {
	return func(h *handler) {
		h.pathFn = fn
	}
}

// Cleaning and parsing the path in a URL
func defaultPathFn(r *http.Request) string {

	// NOTE: Mini useless instant legacy code router
	path := strings.TrimSpace(r.URL.Path)

	if path == "" || path == "/" {
		path = "/intro"
	}

	// NOTE: Python like slicing
	path = path[1:]
	return path
}

// Handling and parsing templates to be rendered on the server
var tmpl *template.Template

// NOTE: How the fuck this invoked w/o manually invoked???
func init() {
	tmpl = template.Must(template.ParseFiles("templates/index.html"))
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := h.pathFn(r)
	if chapter, ok := h.s[path]; ok {

		err := tmpl.Execute(w, chapter)

		if err != nil {

			// Writing For debugging purposes
			log.Printf("%v", err)
			// Writing for the end user/client to see
			http.Error(w, "Something went wrong...", http.StatusInternalServerError)
		}
		return
	}

	http.Error(w, "Chapter not found", http.StatusNotFound)

}

// NOTE: Factory function, very common pattern in
// GO to act as a contructor
func NewHandler(s Story, opts ...HandlerOption) handler {
	h := handler{s, tmpl, defaultPathFn}

	for _, opt := range opts {
		opt(&h)
	}

	return h
}

// NOTE: Usually standard types are preferred to non exported
// types as they do not make it into documentation.
type handler struct {
	s      Story
	t      *template.Template
	pathFn func(r *http.Request) string
}
