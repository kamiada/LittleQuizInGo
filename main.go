package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func main() {
	csvFilename := flag.String("csv", "problems.csv", "csv file in format, question - answer")
	flag.Parse()
	file, err := os.Open(*csvFilename) //* is a pointer to a string - fileName string in this case
	if err != nil {
		exit(fmt.Sprintf("Failed to open csv file: %s", *csvFilename))
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("was unable to open the csv file and read it")
	}
	problems := parseLines(lines)
	//curly {} braces in Go means struct, no object like in JS
	for index, prob := range problems {
		fmt.Printf("Problem #%d: %s = \n", index+1, prob.question)
		var userAnswer string
		fmt.Scanf("%s\n", &userAnswer) //expects only numbers, not string values answers
		//check if answer is correct
		if userAnswer == prob.answer {
			fmt.Println("correct!")
		}
	}
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines)) //make initialises object. Gets type, length of it(int)

	//loop in Go
	for i, line := range lines {
		ret[i] = problem{
			question: line[0],
			answer:   line[1],
		}
	}
	return ret
}

type problem struct {
	question string
	answer   string
}

func exit(msg string) {
	fmt.Printf(msg)
	os.Exit(1) //exit code 1 means that app had some error of some sort
}
