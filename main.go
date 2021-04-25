package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {

	csvFileName := flag.String("csv", "question.csv", "a csv file in a format of 'Q&A'")
	timeLimit := flag.Int("limit", 30, "the time limit for the quiz in second")

	flag.Parse()

	file, err := os.Open(*csvFileName)
	if err != nil {
		exit(fmt.Sprintf("failed to open file %s\n", *csvFileName))
	}
	defer file.Close()

	r := csv.NewReader(file)
	Lines, err := r.ReadAll()
	if err != nil {
		exit(fmt.Sprintf("failed to parse csv file %s\n", *csvFileName))
	}

	question := parsLines(Lines)

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	correct := 0

	for i, v := range question {
		fmt.Printf("question #%d : %s\n", i+1, v.q)
		ch := make(chan string)
		go func(ch chan string) {
			var answer string
			fmt.Scanf("%s\n", &answer)
			ch <- answer

		}(ch)

		select {
		case <-timer.C:
			fmt.Printf("\nYour score is %d of %d\n", correct, len(question))
			return
		case answer := <-ch:
			if answer == v.a {
				correct++
			}
		}

	}
	fmt.Printf("\nYour score is %d of %d\n", correct, len(question))

}

type questions struct {
	q string
	a string
}

func parsLines(Lines [][]string) []questions {
	ret := make([]questions, len(Lines))
	for i, line := range Lines {
		ret[i] = questions{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}

	return ret

}

func exit(msg string) {
	fmt.Printf(msg)
	os.Exit(1)
}
