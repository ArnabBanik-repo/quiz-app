package data

import (
	"fmt"
	"strings"
)

type Question struct {
	title   string
	options []string
	answer  string
	Score   uint
}

func CreateQuestion(title string, options []string, answer string, score uint) Question {
  return Question{title, options, answer, score}
}

func (q *Question) PrintQuestion(number uint) {
	fmt.Printf("Q:%v. %v\t(%v)\n\n", number, q.title, q.Score)
	for i, v := range q.options {
		fmt.Printf("%v: %v\n", i, v)
	}
	fmt.Print("\nAnswer: ")
}

func (q *Question) VerifyAnswer(answer string) uint {
	c := strings.ToLower(q.answer)
	a := strings.Trim(answer, " ")
	a = strings.ToLower(answer)

	if c == a {
		return q.Score
	}
	return 0
}
