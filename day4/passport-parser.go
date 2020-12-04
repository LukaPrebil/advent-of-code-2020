package day4

import (
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)
// ParsePassports converts file of passports into slice of Passport objects
func ParsePassports() []Passport {
	passports := make([]Passport, 0)
	re := regexp.MustCompile("\\s+")
	for _, passport := range getPassports() {
		parts := re.Split(passport, -1)
		passport := buildPassport(parts)
		passports = append(passports, passport)
	}
	return passports
}

// Passport contains all the data on a passport
type Passport struct {
	byr, iyr, eyr           int
	hgt, hcl, ecl, pid, cid string
}

func buildPassport(parts []string) Passport {
	var passport Passport
	for _, part := range parts {
		field := part[:3]
		switch field {
		case "byr":
			byr, err := strconv.Atoi(part[4:])
			handleError(err)
			passport.byr = byr
			break
		case "iyr":
			iyr, err := strconv.Atoi(part[4:])
			handleError(err)
			passport.iyr = iyr
			break
		case "eyr":
			eyr, err := strconv.Atoi(part[4:])
			handleError(err)
			passport.eyr = eyr
			break
		case "hgt":
			passport.hgt = part[4:]
			break
		case "hcl":
			passport.hcl = part[4:]
			break
		case "ecl":
			passport.ecl = part[4:]
			break
		case "pid":
			passport.pid = part[4:]
			break
		case "cid":
			passport.cid = part[4:]
			break
		}
	}
	return passport
}

func getPassports() []string {
	content, err := ioutil.ReadFile("./day4/input.txt")
	handleError(err)

	return strings.Split(string(content), "\n\n")
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
