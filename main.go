package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

type problem struct {
	q string
	a string
}

func parseLines(lines [][]string) []problem {
	res := make([]problem, len(lines))
	for i, line := range lines {
		res[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}
	return res
}

func readProblemsCSV(file string) ([]problem, error) {
	fmt.Println("Reading file: ", file)
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	csvr := csv.NewReader(f)
	lines, err := csvr.ReadAll()
	if err != nil {
		fmt.Println("Failed to parse CSV file.")
		os.Exit(1)
	}
	res := parseLines(lines)
	return res, err
}

func run(file string) {
	problems, err := readProblemsCSV(file)

	if err != nil {
		fmt.Println("Failed to read CSV: ", err)
		os.Exit(1)
	}

	nCorrect := 0
	var ans string
	for i, problem := range problems {
		fmt.Printf("Problem #%d: %s = ", i+1, problem.q)

		fmt.Scanf("%s\n", &ans)
		if ans == problem.a {
			nCorrect++
		}
	}
	fmt.Printf("You scored %d out of %d.\n", nCorrect, len(problems))
}

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	run(*csvFilename)
}
