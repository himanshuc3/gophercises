package main

import (
	"os"
	"time"
	"math-quiz/model"
	"flag"
)

const (
	QUIZ_TIME int64 = 10
)


func main(){

	filename := flag.String("filename", "problems.csv", "name of the file")
	timelimit := flag.Int64("timelimit",QUIZ_TIME,"time limit in seconds")
	flag.Parse()	

	records := ParseCSV(*filename)


	var quiz model.Quiz = model.CreateQuiz(GetQuestionList(records))
	
	// Confused on how to get Channels to work with timers and premature evacuation
	//quizTimer := time.NewTimer(time.Second*QUIZ_TIME)
	//var quizCompleted chan bool


	go func(){
		quiz.StartQuiz()
		print("YAY!!!! Premature evacuation")
		os.Exit(0)
	}()	
	// stops the main thread execution and so that async goroutines/threads 
	// can execute 
	//	initiateTime()
	// select{
	// case input := <-quizCompleted:
	// 	fmt.Println("yay completed before time")
	// 	print(input)
	// 	os.Exit(0)
	// case <-quizTimer.C:
	// 	quiz.DisplayScore()
	// 	fmt.Println("Timer completed")
	// }

	// TO-DO: Replace with channel implementation of the same
	time.Sleep(time.Second*time.Duration(*timelimit))

	quiz.DisplayScore()


}