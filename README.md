# Advent of Code
Fun puzzles to warm my brain up for learning. Also, where I record my thought process.

(1) [First puzzle](https://adventofcode.com/2022/day/1) was a nice introductory puzzle.
- I can calculate the sum of any elf I come across first; store that temporarily and replace the sum (after a blank) if it is larger than my previous sum value. If it is not larger, continue on to the next elf for sum of calories for the next elf.
- Learned to read entire file and acquire binary data via `ioutil.ReadFile` but would be a pain to manage everything in a single string for the purpose of this problem
- Next, I needed to read the file line by line by first opening it and receiving an object with the type - `os.File` which then gets passed into the `bufio.NewScanner` method which is responsible for going through each line of the file. Cool thing is, `.Scan()` triggers the scanner object to generate the next line for data extraction.
