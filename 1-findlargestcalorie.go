package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Advent of Code Day 1")
	fmt.Println("Attempting to read provided file now")
	fmt.Println("")
	file, err := os.Open("./1-puzzle/adventofcode.com_2022_day_1_input.txt")

	if err != nil {
		fmt.Printf("Error found when trying to open file: %s", err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if scanner.Err() != nil {
		log.Fatalf("Error found while reading lines from file: %s", scanner.Err())
	}

	fmt.Println("Lines should now be in array as shown below")
	fmt.Println(lines)
}
