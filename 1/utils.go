package main

import (
	"log"
	"encoding/csv"
	"os"
	"math-quiz/model"
)

func ParseCSV(filename string) [][]string{
	
	file, err := os.Open("./data/" + filename)
	defer file.Close()
	
	if err != nil {
		log.Fatal("Error while opening the problems file", err)
	}

	reader := csv.NewReader(file)

	records,err := reader.ReadAll()

	if err != nil {

		log.Fatal("Error while parsing csv", err)
	}
	return records
}

func GetQuestionList(records [][]string) model.QuestionList {
	
	var ql model.QuestionList
	for _,problem := range records {

		var question model.Question
		question.Text = problem[0]
		question.Ans = problem[1]
		ql = append(ql,question)
	}
	return ql

}