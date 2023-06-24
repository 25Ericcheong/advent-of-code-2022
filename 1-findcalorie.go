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
	file, errFile := os.Open(filePath)

	if errFile != nil {
		log.Fatalf("Error found when trying to open file: %s", errFile)
	}
	defer file.Close()

	var lines, errScan = getData(file)

	if errScan != nil {
		log.Fatalf("Error while reading through each line of file: %s", errScan)
	}

	fmt.Println("Each line found in find should now be in an array")

	var maxCalorie int
	var sum int
	for _, data := range lines {
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

	fmt.Printf("Program has finished execution. Largest calorie is found to be %d \n", maxCalorie)
}

func getData(file *os.File) ([]string, error) {
	var lines []string
	scanner := bufio.NewScanner(file)

	// note that this is similar to a while loop
	// loop continues as long as .Scan() is true
	for scanner.Scan() {
		// calling Scan, generates next line of data found in file
		lines = append(lines, scanner.Text())
	}

	if scanner.Err() != nil {
		return lines, scanner.Err()
	}

	fmt.Println("Each line found in find should now be in an array")
	return lines, nil
}
