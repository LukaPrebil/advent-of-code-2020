package day1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// Expenses calculates solution for day 1. Find two numbers summing up to 2020, multiply them.
func Expenses() int {
	numbers := getNumbers()
	for idx, number := range numbers {
		for _, candidate := range numbers[idx+1:] {
			if number+candidate == 2020 {
				fmt.Printf("Numbers that sum to 2020 are %d and %d\n", number, candidate)
				return number * candidate
			}
		}
	}
	log.Fatal("Could not find two numbers summing to 2020. :( Bye.")
	return -1
}

// getNumbers parses the input file and appends each line as an integer to the result slice
func getNumbers() []int {

	// Open input file
	file, err := os.Open("./day1/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Setup file scanner to read lines
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	numbers := make([]int, 0)

	// Parse each line into integer and append to numbers
	for scanner.Scan() {
		number, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatalf("Line is not a number: %s", err)
		}
		numbers = append(numbers, number)
	}

	return numbers
}
