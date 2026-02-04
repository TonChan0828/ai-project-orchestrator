package domain

type ProjectOverview struct {
	OriginalText string
	Assumptions  []string
	Glossary     []Term
}

type Term struct {
	Name        string
	Description string
}
