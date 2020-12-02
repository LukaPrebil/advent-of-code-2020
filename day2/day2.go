package day2

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

// CountValid counts the amount of valid passwords in the input
func CountValid() int {
	passwordLines := parsePasswords()
	valid := 0
	for _, line := range passwordLines {
		if line.rule.Validate(line.password) {
			valid++
		}
	}
	return valid
}

// PasswordRule defines conditions for when a password is valid
type PasswordRule struct {
	character string
	min, max  int
}

// PasswordLine represents direct data parsed from a single line of input
type PasswordLine struct {
	password string
	rule     PasswordRule
}

// PasswordValidator validates a password given a rule
type PasswordValidator interface {
	Validate(password string) bool
}

// Validate defines validation of a password based on a PasswordRule
func (rule PasswordRule) Validate(password string) bool {
	occurences := strings.Count(password, rule.character)
	return rule.min <= occurences && occurences <= rule.max
}

// parsePasswords parses the input file and appends each line as an integer to the result slice
func parsePasswords() []PasswordLine {
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
		character := string(splits[1][0])

		// The actual password
		password := splits[2]

		passwordLines = append(passwordLines, PasswordLine{password, PasswordRule{character, min, max}})
	}

	return passwordLines
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
