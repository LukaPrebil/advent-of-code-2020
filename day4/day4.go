package day4

import (
	"regexp"
	"strconv"
)

// ValidatePassports counts the number of valid passports
func ValidatePassports() int {
	count := 0
	for _, pass := range ParsePassports() {
		if pass.Validate() {
			count++
		}
	}
	return count
}

// ValidatePassportsStrong counts the number of valid passports with valid data
func ValidatePassportsStrong() int {
	count := 0
	for _, pass := range ParsePassports() {
		if pass.ValidateStrong() {
			count++
		}
	}
	return count
}

type validatable interface {
	Validate() bool
	ValidateStrong() bool
}

// Validate validates that all fields apart from CID are present
func (pass Passport) Validate() bool {
	return pass.byr != 0 &&
		pass.iyr != 0 &&
		pass.eyr != 0 &&
		pass.hgt != "" &&
		pass.hcl != "" &&
		pass.ecl != "" &&
		pass.pid != ""
}

// ValidateStrong validates that all fields apart from CID are present and checks values of data
func (pass Passport) ValidateStrong() bool {
	yearsOk := 1920 <= pass.byr && pass.byr <= 2002 &&
		2010 <= pass.iyr && pass.iyr <= 2020 &&
		2020 <= pass.eyr && pass.eyr <= 2030

	otherDataPresent := pass.hgt != "" && pass.hcl != "" && pass.ecl != "" && pass.pid != ""

	return yearsOk && otherDataPresent && checkHeight(pass.hgt) && checkHairColour(pass.hcl) && checkEyeColour(pass.ecl) && len(pass.pid) == 9
}

func checkHeight(height string) bool {
	if height[len(height)-2:] == "cm" {
		hgt, err := strconv.Atoi(height[:len(height)-2])
		handleError(err)
		return 150 <= hgt && hgt <= 193
	} else if height[len(height)-2:] == "in" {
		hgt, err := strconv.Atoi(height[:len(height)-2])
		handleError(err)
		return 59 <= hgt && hgt <= 76
	} else {
		return false
	}
}

func checkHairColour(colour string) bool {
	re := regexp.MustCompile("^#[0-9a-f]{6}$")
	return re.Match([]byte(colour))
}

func checkEyeColour(colour string) bool {
	re := regexp.MustCompile("^amb|blu|brn|gry|grn|hzl|oth")
	return re.Match([]byte(colour))
}
