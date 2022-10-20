package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// An interface puzzle that has a function solve()

type puzzle interface {
	solve()
}

// a function solve() that returns true or false

func (r *riddle) solve() bool {
	fmt.Println(r.question) // Printing the riddle question
	reader := bufio.NewReader(os.Stdin)
	// ReadString will block until the delimiter is entered
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		return false
	}
	// remove the delimeter from the string
	userAnswer := strings.TrimSuffix(input, "\n")
	var riddleAnswer = r.questionAnswer // Setting riddle answer to the correct answer for the riddle
	if riddleAnswer == userAnswer {     // This checks the inputed user answer against questionAnswer in riddle
		return true
	} else {
		return false
	}

}

// A struct riddle that has a question and a correct answer
type riddle struct {
	question       string
	questionAnswer string
	puzzle

	//A riddle satisfies the puzzle interface by displaying the question to the console output
	//and reading in the user’s answer, returning true if the user’s answer is the correct answer
}

func main() {
	QA := riddle{question: "Why did the vampire take art class? ", questionAnswer: "He wanted to learn how to draw blood"}
	fmt.Print(QA.solve())
}
