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

func TestCommonAnswers(t *testing.T) {
	input := []string {"abcd", "abc", "bcd"}
	expectedResult := 2

	result := getNumberOfCommonAnswers(input)

	if expectedResult != result {
		t.Errorf("Incorrectly calculated common answers, expected: %d, but got %d", expectedResult, result)
	}
}

func TestCommonAnswersWithDuplicates(t *testing.T) {
	input := []string {"aabcd", "abbc", "bccdd"}
	expectedResult := 2

	result := getNumberOfCommonAnswers(input)

	if expectedResult != result {
		t.Errorf("Incorrectly calculated common answers, expected: %d, but got %d", expectedResult, result)
	}
}

func TestNoCommonAnswers(t *testing.T) {
	input := []string {"abcd", "abc", "d"}
	expectedResult := 0

	result := getNumberOfCommonAnswers(input)

	if expectedResult != result {
		t.Errorf("Incorrectly calculated common answers, expected: %d, but got %d", expectedResult, result)
	}
}

func TestCommonAnswersOneAnswer(t *testing.T) {
	input := []string {"abcd"}
	expectedResult := 4

	result := getNumberOfCommonAnswers(input)

	if expectedResult != result {
		t.Errorf("Incorrectly calculated common answers, expected: %d, but got %d", expectedResult, result)
	}
}

func TestCommonAnswersNoAnswers(t *testing.T) {
	input := []string {}
	expectedResult := 0

	result := getNumberOfCommonAnswers(input)

	if expectedResult != result {
		t.Errorf("Incorrectly calculated common answers, expected: %d, but got %d", expectedResult, result)
	}
}



// --------- Utils ---------

// setsDiffer performs Difference on both sets, and checks that the resulting set Cardinality is not 0 (there is something in the difference set)
func setsDiffer(first, second mapset.Set) bool {
	return first.Difference(second).Cardinality() != 0
}