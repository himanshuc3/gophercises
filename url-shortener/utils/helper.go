package utils

import (
	yaml "gopkg.in/yaml.v2"
	"os"
	"log"
)




func LoadYaml(filename string) []PathUrl{

	// NOTE: Returns []byte instead of *os.File
	file, err := os.ReadFile("data/" + filename)
	
	if err != nil {
		log.Fatal("Error while opening the problems file", err)
	}

	yamlPaths, err := parseYaml(file)

	if err != nil {

		log.Fatal("Error while parsing csv", err)
	}
	return yamlPaths
}

func parseYaml(data []byte) ([]PathUrl, error) {
	var pathUrls []PathUrl
	if err := yaml.Unmarshal(data,&pathUrls); err != nil {
		return nil, err
	}
	return pathUrls, nil	
}
