package main

import (
	"cyoa/utils"
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := flag.Int("port", 3000, "which port do you want the application to run on?")
	filename := flag.String("filename", "story.json", "Enter story in format similar to example file")

	flag.Parse()

	story := utils.ParseJSON(*filename)

	h := utils.NewHandler(story)

	fmt.Printf("Starting the server on port: %d\n", *port)
	// the handler should have ServeHTTP method attached essentially
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))

}
