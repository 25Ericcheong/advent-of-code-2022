package a_day_one

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func Trebuchet() {
	log.Println("Starting Trebuchet challenge")
	file, err := os.Open("./a_day_one/trebuchet_input.txt")

	if err != nil {
		log.Fatal("error when trying to open trebuchet input file: ", err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		firstDigit, lastDigit := getDigit(scanner.Text())
		fmt.Printf("Total %d\n", firstDigit+lastDigit)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal("error when trying to close file: ", err)
		}
	}(file)
}

func getDigit(text string) (firstDigit int, secondDigit int) {
	textLength := len(text)
	first := 0
	second := 0

	for i := 0; i < textLength; i++ {
		num, err := strconv.Atoi(text[i : i+1])
		if err == nil {
			first = num
			break
		}
	}

	for i := textLength - 1; i > -1; i-- {
		num, err := strconv.Atoi(text[i : i+1])

		if err == nil {
			second = num
			break
		}
	}
	return first * 10, second
}
