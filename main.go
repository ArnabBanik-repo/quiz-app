package main

import (
	"quizApp/data"
	"time"

	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var questions []data.Question

func greeting() {
	fmt.Println("Welcome to the quiz")
	fmt.Println()
}

func goodbye(user_score uint, total_score uint) {
	fmt.Printf("\nYou scored: %v/%v\n", user_score, total_score)
	fmt.Println("See ya!")
  os.Exit(0)
}

func readQuestions(csvFileName *string) ([]data.Question, error) {


	file, err := os.Open(*csvFileName)
	if err != nil {
		return nil, err
	}

	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		return nil, err
	}
	if len(lines) == 0 {
		return nil, err
	}

	noptions := len(lines[0]) - 3

	for _, v := range lines {
		t := make([]string, noptions)
		for j := 0; j < noptions; j += 1 {
			t[j] = v[j+1]
		}

		score, err := strconv.ParseUint(v[noptions+2], 10, 32)
		if err != nil {
			fmt.Println(err)
			fmt.Printf("Skipping question %v", v)
			continue
		}
		q := data.CreateQuestion(v[0], t, v[noptions+1], uint(score))
		questions = append(questions, q)
	}

	return questions, nil
}

func playGame(questions []data.Question, done chan bool, user_score *uint, total_score *uint) {

	reader := bufio.NewReader(os.Stdin)

	for i, question := range questions {
		question.PrintQuestion(uint(i + 1))
		answer, _ := reader.ReadString('\n')
		answer = strings.Replace(answer, "\n", "", -1)
		*user_score += question.VerifyAnswer(answer)
		*total_score += question.Score
	}
	done <- true
	return
}

func main() {
	csvFileName := flag.String("csv", "problems.csv", "The question paper in a .csv formatted as 'question,(options),answer,score'")
  timeLimit := flag.Int("limit", 5, "The time limit in seconds")
  flag.Parse()

	questions, err := readQuestions(csvFileName)
	if err != nil {
		fmt.Println(err)
		return
	}

	greeting()

  var user_score, total_score uint
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
  quizChan := make(chan bool)
	go playGame(questions, quizChan, &user_score, &total_score)

	select {
	case <-quizChan:
    goodbye(user_score, total_score)
	case <-timer.C:
    fmt.Print("\nTime Up!")
    goodbye(user_score, total_score)
	}
}
