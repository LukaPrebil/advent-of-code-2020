package day5

import (
	"testing"
)

func TestParseSeatCorrectlyParsesSimpleSeat(t *testing.T) {
	boardingPass := "FBFBBFFRLR"

	seat := parseSeat(boardingPass)

	if seat.row != 44 {
		t.Errorf("Row was incorrect, expected %d but got %d", 44, seat.row)
	}
	if seat.column != 5 {
		t.Errorf("Column was incorrect, expected %d but got %d", 5, seat.column)
	}
}

func TestParseSeatCorrectlyParsesZerothSeat(t *testing.T) {
	boardingPass := "FFFFFFFLLL"

	seat := parseSeat(boardingPass)

	if seat.row != 0 {
		t.Errorf("Row was incorrect, expected %d but got %d", 0, seat.row)
	}
	if seat.column != 0 {
		t.Errorf("Column was incorrect, expected %d but got %d", 0, seat.column)
	}
}

func TestParseSeatCorrectlyParsesLastSeat(t *testing.T) {
	boardingPass := "BBBBBBBRRR"

	seat := parseSeat(boardingPass)

	if seat.row != 127 {
		t.Errorf("Row was incorrect, expected %d but got %d", 127, seat.row)
	}
	if seat.column != 7 {
		t.Errorf("Column was incorrect, expected %d but got %d", 7, seat.column)
	}
}