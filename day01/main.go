package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0
	max := 0
	max1 := 0
	max2 := 0
	for scanner.Scan() {
		if scanner.Text() == "\n" || scanner.Text() == "" {
			// new elf
			s, minM := minMax(max, max1, max2)
			if sum >= minM {
				if s == "m1" {
					max = sum
				} else if s == "m2" {
					max1 = sum
				} else if s == "m3" {
					max2 = sum
				}
			}
			sum = 0
			continue
		}
		cals, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		sum += cals
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(max, max1, max2, max+max1+max2)
}

func minMax(m1 int, m2 int, m3 int) (string, int) {
	if m1 <= m2 && m1 <= m3 {
		return "m1", m1
	} else if m2 <= m1 && m2 <= m3 {
		return "m2", m2
	} else if m3 <= m1 && m3 <= m2 {
		return "m3", m3
	}

	return "", 0
}
