package main

import (
	"fmt"

	"github.com/LukaPrebil/advent-of-code-2020/day1"
	"github.com/LukaPrebil/advent-of-code-2020/day2"
	"github.com/LukaPrebil/advent-of-code-2020/day3"
	"github.com/LukaPrebil/advent-of-code-2020/day4"
	"github.com/LukaPrebil/advent-of-code-2020/day5"
	"github.com/LukaPrebil/advent-of-code-2020/day6"
)

func main() {
	fmt.Printf("Solution for 1st day:\n\tPart 1: %d\n\tPart 2: %d\n", day1.Expenses(), day1.Expenses3())
	fmt.Printf("Solution for 2nd day:\n\tPart 1: %d\n\tPart 2: %d\n", day2.CountValid(false), day2.CountValid(true))
	fmt.Printf("Solution for 3rd day:\n\tPart 1: %d\n\tPart 2: %d\n", day3.CountTrees(), day3.MinimizeStops())
	fmt.Printf("Solution for 4th day:\n\tPart 1: %d\n\tPart 2: %d\n", day4.ValidatePassports(), day4.ValidatePassportsStrong())
	fmt.Printf("Solution for 5th day:\n\tPart 1: %d\n\tPart 2: %d\n", day5.GetMaxSeatID(), day5.GetMySeatID())
	fmt.Printf("Solution for 6th day:\n\tPart 1: %d\n\tPart 2: %d\n", day6.SumGroupAnswersAnyone(), day6.SumGroupAnswersEveryone())
}
