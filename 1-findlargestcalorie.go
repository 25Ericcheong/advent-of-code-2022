package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

const filePath = "./1-puzzle/adventofcode.com_2022_day_1_input.txt"

func main() {
	fmt.Println("Advent of Code Day 1")
	fmt.Println("Attempting to open file now")
	file, err := os.Open(filePath)

	if err != nil {
		fmt.Printf("Error found when trying to open file: %s", err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)

	// note that this is similar to a while loop
	// loop continues as long as .Scan() is true
	for scanner.Scan() {
		// calling Scan, generates next line of data found in file
		lines = append(lines, scanner.Text())
	}

	if scanner.Err() != nil {
		log.Fatalf("Error found while reading lines from file: %s", scanner.Err())
	}

	fmt.Println("Each line found in find should now be in an array")

	var maxCalorie int
	var sum int
	for index, data := range lines {
		fmt.Println(sum)
		fmt.Println(data)
		if index == 30 {
			break
		}

		if data == "" {
			if sum > maxCalorie {
				maxCalorie = sum
			}
			sum = 0

		} else {

			num, err := strconv.Atoi(data)
			if err != nil {
				log.Fatalf("Error found when trying to parse string to int: %s", err)
			}

			sum += num
		}
	}

	fmt.Println("Each line found in find should now be in an array")
}
