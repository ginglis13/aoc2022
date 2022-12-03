package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	scores := map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
		"X": 1,
		"Y": 2,
		"Z": 3,
	}
	// W, D, L for win draw lose
	rock := map[string]int{
		"A": 3,
		"B": 0,
		"C": 6,
	}
	paper := map[string]int{
		"A": 6,
		"B": 3,
		"C": 0,
	}
	scissors := map[string]int{
		"A": 0,
		"B": 6,
		"C": 3,
	}
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	totScore := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		a := strings.Split(line, " ")
		totScore += scores[a[1]]
		if a[1] == "X" {
			totScore += rock[a[0]]
		} else if a[1] == "Y" {
			totScore += paper[a[0]]
		} else if a[1] == "Z" {
			totScore += scissors[a[0]]
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(totScore)
}

// score = shape selected + outcome
// 1 Rock
// 2 Paper
// 3 Scissors
// 0 L
// 3 Draw
// 6 W
