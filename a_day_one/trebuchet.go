package a_day_one

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func Trebuchet() {
	log.Println("Starting Trebuchet challenge")
	file, err := os.Open("./a_day_one/trebuchet_input.txt")

	if err != nil {
		log.Fatal("error when trying to open trebuchet input: ", err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal("error when trying to close file: ", err)
		}
	}(file)
}
