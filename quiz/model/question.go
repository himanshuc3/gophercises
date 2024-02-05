package model 

import (
	"fmt"
)


type Question struct{
	Text string
	Ans string
}

func (q Question) String() string {
	return fmt.Sprintf("%v :: %v", q.Text, q.Ans)
}

type QuestionList []Question

type Quiz struct {
	Qs QuestionList
	Score int
	CurrentQ int
}


func CreateQuiz(questionList QuestionList) Quiz{
	return Quiz{questionList, 0,0}
}

func (quiz Quiz) getNextQuestion() string {
	return quiz.Qs[quiz.CurrentQ].Text
}

func (quiz Quiz) StartQuiz() {
	quiz.resetScore()
	// length := len(QuestionList)
	var userInput string
	for i, question := range quiz.Qs{
		fmt.Printf("Problem #%v: %v = ",i+1, question.Text)
		fmt.Scanf("%v", &userInput)

		if question.validate(userInput){
			quiz.Score += 1
		}

	}
	quiz.DisplayScore()
	//quizCompleted <- true
}

func (quiz Quiz) DisplayScore(){

	fmt.Printf("\nYour score is %v out of %v \n", quiz.Score, len(quiz.Qs))
}

func (q Question) validate(ans string) bool{
	return q.Ans == ans 
} 

func (quiz Quiz) resetScore() {
	quiz.Score = 0;
}

func (this QuestionList) String() string {
 	var output string

 	for _, problem := range this {
 		output = output + problem.Text + "::" + problem.Ans + "\n"
 	}
 	return output
}

