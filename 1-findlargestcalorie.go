package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	fmt.Println("Advent of Code Day 1")
	fmt.Println("Attempting to read provided file now")
	fmt.Println("")
	b, err := ioutil.ReadFile("./1-puzzle/adventofcode.com_2022_day_1_input.txt")

	if err != nil {
		fmt.Println(err)
	}

	str := string(b)
	fmt.Println(str)
}
