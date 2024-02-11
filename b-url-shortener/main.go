package main

import (
	"fmt"
	"net/http"
	"urlshort/utils"
	"flag"
)

func main() {
	mux := defaultMux()

	yamlFile := flag.String("yml", "default.yml", "Enter the name of the yaml file located in data folder") 

	flag.Parse()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := utils.MapHandler(pathsToUrls, mux)

	yamlPaths := utils.LoadYaml(*yamlFile)  	
	
	// Build the YAMLHandler using the mapHandler as the
	// fallback
	yamlHandler, err := utils.YAMLHandler(yamlPaths, mapHandler)
	if err != nil {
		panic(err)
	}
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", yamlHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}