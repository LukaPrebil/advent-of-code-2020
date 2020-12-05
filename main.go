package main

import (
	"fmt"

	"gitlab.com/LukaPG/advent-of-code-2020/day1"
	"gitlab.com/LukaPG/advent-of-code-2020/day2"
	"gitlab.com/LukaPG/advent-of-code-2020/day3"
	"gitlab.com/LukaPG/advent-of-code-2020/day4"
	"gitlab.com/LukaPG/advent-of-code-2020/day5"
)

func main() {
	fmt.Printf("Solution for 1st day:\n\tPart 1: %d\n\tPart 2: %d\n", day1.Expenses(), day1.Expenses3())
	fmt.Printf("Solution for 2nd day:\n\tPart 1: %d\n\tPart 2: %d\n", day2.CountValid(false), day2.CountValid(true))
	fmt.Printf("Solution for 3rd day:\n\tPart 1: %d\n\tPart 2: %d\n", day3.CountTrees(), day3.MinimizeStops())
	fmt.Printf("Solution for 4th day:\n\tPart 1: %d\n\tPart 2: %d\n", day4.ValidatePassports(), day4.ValidatePassportsStrong())
	fmt.Printf("Solution for 5th day:\n\tPart 1: %d\n\tPart 2: %d\n", day5.GetMaxSeatID(), day5.GetMySeatID())
}
