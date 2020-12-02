package main

import (
	"fmt"

	"gitlab.com/LukaPG/advent-of-code-2020/day1"
	"gitlab.com/LukaPG/advent-of-code-2020/day2"
)

func main() {
	fmt.Printf("Solution for first day:\n\tPart 1: %d\n\tPart 2: %d\n", day1.Expenses(), day1.Expenses3())
	fmt.Printf("Solution for second day:\n\tPart 1: %d\n\tPart 2: %d\n", day2.CountValid(false), day2.CountValid(true))
}
