package domain

type Architecture struct {
	Components []Component
	DataFlows  []DataFlow
}

type Component struct {
	Name           string
	Responsibility string
	DependsOn      []string
}

type DataFlow struct {
	From        string
	To          string
	Description string
}
