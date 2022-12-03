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
		"X": 1, // r
		"Y": 2, // p
		"Z": 3, // s
	}
	// Determine the move needed for input outcome
	rock := map[string]string{
		"X": "Z", // lose
		"Y": "X", // draw
		"Z": "Y", // win
	}
	paper := map[string]string{
		"X": "X", // lose
		"Y": "Y", // draw
		"Z": "Z", // win
	}
	scissors := map[string]string{
		"X": "Y", // paper lose
		"Y": "Z", // scis draw
		"Z": "X", // rock win
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
		if a[0] == "A" {
			totScore += scores[rock[a[1]]]
		} else if a[0] == "B" {
			totScore += scores[paper[a[1]]]
		} else if a[0] == "C" {
			totScore += scores[scissors[a[1]]]
		}

		if a[1] == "Y" {
			totScore += 3

		} else if a[1] == "Z" {
			totScore += 6
		}
		fmt.Println(a, totScore)
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
