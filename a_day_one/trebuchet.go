package a_day_one

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Trebuchet() {
	fmt.Println("starting Trebuchet challenge")
	fmt.Println("")
	file, err := os.Open("./a_day_one/trebuchet_input.txt")

	if err != nil {
		log.Fatal("error when trying to open trebuchet input file: ", err)
	}

	scanner := bufio.NewScanner(file)
	total := 0
	fmt.Println("going through provided trebuchet input file now")
	for scanner.Scan() {
		firstDigit, lastDigit := getDigit(scanner.Text())
		total += firstDigit + lastDigit
	}
	fmt.Printf("finish calculation, total found to be %d\n", total)

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
	spelledNum := ""
	spelledNumToNum := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	for i := 0; i < textLength; i++ {
		spelledNum += text[i : i+1]
		numInput, err := strconv.Atoi(text[i : i+1])
		if err == nil {
			first = numInput
			break
		}

		if len(spelledNum) > 2 {
			for spelledNumKey, numSpelled := range spelledNumToNum {
				if strings.Contains(spelledNum, spelledNumKey) {
					first = numSpelled
					goto exitForFirstDigit
				}
			}
		}
	}
exitForFirstDigit:

	spelledNum = ""

	for i := textLength - 1; i > -1; i-- {
		spelledNum = text[i:i+1] + spelledNum
		num, err := strconv.Atoi(text[i : i+1])

		if err == nil {
			second = num
			break
		}

		if len(spelledNum) > 2 {
			for spelledNumKey, numSpelled := range spelledNumToNum {
				if strings.Contains(spelledNum, spelledNumKey) {
					second = numSpelled
					goto exitForSecondDigit
				}
			}
		}
	}
exitForSecondDigit:

	return first * 10, second
}
