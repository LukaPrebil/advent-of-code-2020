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

func getNumberOfCommonAnswers(declarationPerMember []string) int {
	commonDeclarations := mapset.NewSet()
	for i, decl := range declarationPerMember {
		if i == 0 { // Setup the initial condition to be the set of first answers
			commonDeclarations = commonDeclarations.Union(dedupeDeclaration(decl))
		} else { // Keep intersecting the set with other answers to get the common ones
			commonDeclarations = commonDeclarations.Intersect(dedupeDeclaration(decl))
		}
	}
	return commonDeclarations.Cardinality()
}

func getNumberOfUniqueAnswers(declaration mapset.Set) int {
	return declaration.Cardinality()
}