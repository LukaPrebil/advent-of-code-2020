package day6

import (
	"testing"
	"github.com/deckarep/golang-set"
)

func TestConvertToSingleLine(t *testing.T) {
	input := "foo\nbar\nbaz"
	expectedResult := "foobarbaz"

	result := toSingleLine(input)

	if result != expectedResult {
		t.Errorf("Did not convert to single line, expected: %s, but got %s", expectedResult, result)
	}
}

func TestConvertToSingleLineEmpty(t *testing.T) {
	input := ""
	expectedResult := ""

	result := toSingleLine(input)

	if result != expectedResult {
		t.Errorf("Did not convert to single line, expected: %s, but got %s", expectedResult, result)
	}
}

func TestConvertToSingleLineAlreadySingleLine(t *testing.T) {
	input := "foo"
	expectedResult := "foo"

	result := toSingleLine(input)

	if result != expectedResult {
		t.Errorf("Did not convert to single line, expected: %s, but got %s", expectedResult, result)
	}
}

func TestDedupe(t *testing.T) {
	input := "foo"
	expectedResult := mapset.NewSetWith('f', 'o')

	result := dedupeDeclaration(input)

	if setsDiffer(expectedResult, result) {
		t.Errorf("Did not deduplicate string, expected: %s, but got %s", expectedResult, result)
	}
}

func TestDedupeDeduped(t *testing.T) {
	input := "fo"
	expectedResult := mapset.NewSetWith('f', 'o')

	result := dedupeDeclaration(input)

	if setsDiffer(expectedResult, result) {
		t.Errorf("Did weird thing with deduplicated string, expected: %s, but got %s", expectedResult, result)
	}
}

func TestDedupeEmpty(t *testing.T) {
	input := ""
	expectedResult := mapset.NewSet()

	result := dedupeDeclaration(input)

	if setsDiffer(expectedResult, result) {
		t.Errorf("Did weird thing with deduplicated string, expected: %s, but got %s", expectedResult, result)
	}
}

// setsDiffer performs Difference on both sets, and checks that the resulting set Cardinality is not 0 (there is something in the difference set)
func setsDiffer(first, second mapset.Set) bool {
	return first.Difference(second).Cardinality() != 0
}