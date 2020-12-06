package day6

import "strings"

// SumGroupAnswersAnyone computes the sum off all per-group unique answers in declarations
func SumGroupAnswersAnyone() int {
	count := 0
	for _, declaration := range getDeclarations() {
		singleLineDecl := toSingleLine(declaration)
		dedupedDecl := dedupeDeclaration(singleLineDecl)
		count += getNumberOfUniqueAnswers(dedupedDecl)
	}
	return count
}

// SumGroupAnswersEveryone computes the sum off all per-group common answers in declarations
func SumGroupAnswersEveryone() int {
	count := 0
	for _, groupDecl := range getDeclarations() {
		groupDeclarations := strings.Split(groupDecl, "\n")
		count += getNumberOfCommonAnswers(groupDeclarations)
	}
	return count
}