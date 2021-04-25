package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func main() {

	csvFileName := flag.String("csv", "question.csv", "a csv file in a format of 'Q&A'")
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
	correct := 0
	for i, p := range question {
		fmt.Printf("question #%d : %s\n", i+1, p.q)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == p.a {
			correct++
		}
	}
	fmt.Printf("Your score is %d of %d\n", correct, len(question))

}

type questions struct {
	q string
	a string
}

func parsLines(Lines [][]string) []questions {
	ret := make([]questions, len(Lines))
	for i, line := range Lines {
		ret[i] = questions{
			line[0],
			line[1],
		}
	}

	return ret

}

func exit(msg string) {
	fmt.Printf(msg)
	os.Exit(1)
}
