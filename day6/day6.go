package day6

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
