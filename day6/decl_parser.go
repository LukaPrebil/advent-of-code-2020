package day6

import (
	"io/ioutil"
	"strings"
	"github.com/deckarep/golang-set"
)

func getDeclarations() []string {
	content, _ := ioutil.ReadFile("./day6/input.txt")

	return strings.Split(string(content), "\n\n")
}

func toSingleLine(s string) string {
	return strings.ReplaceAll(s, "\n", "")
}

func dedupeDeclaration(s string) mapset.Set {
	declSet := mapset.NewSet()
	for _, decl := range s {
		declSet.Add(decl)
	}
	return declSet
}

func getNumberOfUniqueAnswers(declaration mapset.Set) int {
	return declaration.Cardinality()
}