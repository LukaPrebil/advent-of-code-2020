package day5

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type seat struct {
	row, column int
}

func parseSeat(boardingPass string) seat {
	rowBinary, colBinary := boardingPass[:7], boardingPass[7:]

	// Convert to binary number representation
	rowBinary = strings.ReplaceAll(rowBinary, "B", "1")
	rowBinary = strings.ReplaceAll(rowBinary, "F", "0")
	colBinary = strings.ReplaceAll(colBinary, "R", "1")
	colBinary = strings.ReplaceAll(colBinary, "L", "0")

	rowNumber, _ := strconv.ParseInt(rowBinary, 2, 8)
	colNumber, _ := strconv.ParseInt(colBinary, 2, 8)

	return seat{int(rowNumber), int(colNumber)}
}

func parseFile() []seat {
	content, err := ioutil.ReadFile("./day5/input.txt")
	handleError(err)

	seats := make([]seat, 0)
	for _, boardingPass := range strings.Split(string(content), "\n") {
		seats = append(seats, parseSeat(boardingPass))
	}
	return seats
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}