package day2

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

// CountValid counts the amount of valid passwords in the input
func CountValid(isSecondPart bool) int {
	passwordLines := parsePasswords(isSecondPart)
	valid := 0
	for _, line := range passwordLines {
		if line.rule.Validate(line.password) {
			valid++
		}
	}
	return valid
}

// PasswordRuleCount defines conditions for when a password is valid
type PasswordRuleCount struct {
	character string
	min, max  int
}

// PasswordRuleLocation defines conditions for when a password is valid
type PasswordRuleLocation struct {
	character byte
	first, second  int
}

// PasswordLine represents direct data parsed from a single line of input
type PasswordLine struct {
	password string
	rule     PasswordValidator
}

// PasswordValidator validates a password given a rule
type PasswordValidator interface {
	Validate(password string) bool
}

// Validate checks that password contains between min and max occurences of character
func (rule PasswordRuleCount) Validate(password string) bool {
	occurences := strings.Count(password, rule.character)
	return rule.min <= occurences && occurences <= rule.max
}

// Validate checks that firstPosition XOR secondPosition contains the character
func (rule PasswordRuleLocation) Validate(password string) bool {
	// -1 to account for 1 based indexing
	first := password[rule.first-1] == rule.character
	second := password[rule.second-1] == rule.character

	// X ^ Y == (X || Y) && !(X && Y)
	return (first || second) && !(first && second)
}

// parsePasswords parses the input file and appends each line as an integer to the result slice
func parsePasswords(isSecondPart bool) []PasswordLine {
	// Open input file
	file, err := os.Open("./day2/input.txt")
	handleError(err)
	defer file.Close()

	// Setup file scanner to read lines
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	passwordLines := make([]PasswordLine, 0)

	// Parse each line into rule and password
	for scanner.Scan() {
		splits := strings.SplitN(scanner.Text(), " ", 3)
		limits := strings.SplitN(splits[0], "-", 2)

		// Rule minimum occurences
		min, err := strconv.Atoi(limits[0])
		handleError(err)

		// Rule maximum occurences
		max, err2 := strconv.Atoi(limits[1])
		handleError(err2)

		// What character the rule looks for ("a:" -> "a")
		character := splits[1][0]

		// The actual password
		password := splits[2]

		if isSecondPart {
			passwordLines = append(passwordLines, PasswordLine{password, PasswordRuleLocation{character, min, max}})
		} else {
			passwordLines = append(passwordLines, PasswordLine{password, PasswordRuleCount{string(character), min, max}})
		}
	}

	return passwordLines
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
