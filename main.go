package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
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

func run(file string, timeLimit int) {
	problems, err := readProblemsCSV(file)
	if err != nil {
		fmt.Println("Failed to read CSV: ", err)
		os.Exit(1)
	}

	nCorrect := 0
	timer := time.NewTimer(time.Duration(timeLimit) * time.Second)
	// Block until the goroutine gets a message from the channel (timer done)
	// Exit the program when timer runs out.

	ansCh := make(chan string)
problemLoop:
	for i, problem := range problems {
		fmt.Printf("Problem #%d: %s = ", i+1, problem.q)

		go func() {
			var ans string
			fmt.Scanf("%s\n", &ans)
			ansCh <- ans
		}()

		select {
		case <-timer.C:
			fmt.Println()
			break problemLoop
		case ans := <-ansCh:
			if ans == problem.a {
				nCorrect++
			}
		}
	}
	fmt.Printf("You scored %d out of %d.\n", nCorrect, len(problems))
}

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	timeLimit := flag.Int("limit", 30, "the time limit for the quiz in seconds.")
	flag.Parse()
	run(*csvFilename, *timeLimit)
}
