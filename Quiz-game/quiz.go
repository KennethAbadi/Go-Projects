package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

type question struct {
	question string
	answer   string
}

type quiz struct {
	answered  int
	correct   int
	questions []question
}

func errormsg(message string, err error) {
	if err != nil {
		log.Fatalln(message, ":", err)
	}
	//Provides error when file isnt read.

}

func readfile(filepath string) *quiz {
	csvFile, err := os.Open(filepath)
	errormsg("Error opening CSV file", err)
	defer csvFile.Close()

	reader := csv.NewReader(bufio.NewReader(csvFile))

	var quiz quiz

	for {
		line, err := reader.Read()

		if err == io.EOF {
			break
		}

		errormsg("Error", err)
		question := question{line[0], line[1]}

		quiz.questions = append(quiz.questions, question)
	}

	return &quiz
}

func (quiz *quiz) run() {

	for _, question := range quiz.questions {
		fmt.Println(question.question)
		answerchan := make(chan string)

		go func() {
			scanner.Scan()
			answer := scanner.Text()
			answerchan <- answer
		}()

		select {
		case answer := <-answerchan:
			if answer == question.answer {
				quiz.correct++
			}
			quiz.answered++
		}
	}
	return
}

func (quiz *quiz) output() {
	fmt.Printf(
		"You answered %v questions out of %v, you got %v questions correct",
		quiz.answered,
		len(quiz.questions),
		quiz.correct,
	)
}

var (
	scanner     = bufio.NewScanner(os.Stdin)
	filePathPtr = flag.String("file", "./problems.csv", "Path to csv file containing quiz.")
)

func main() {
	quiz := readfile(*filePathPtr)
	quiz.run()
	quiz.output()
}
