package domain

type Requirements struct {
	Functional    []RequirementItem
	NonFunctional []RequirementItem
	Constraints   []string
	OutOfScore    []string
}

type RequirementItem struct {
	ID          string
	Description string
	Rationale   string
}
