package domain

type Plan struct {
	Epics []Epic
}

type Epic struct {
	Name    string
	Stories []Story
}

type Story struct {
	Title string
	Tasks []Task
}

type Task struct {
	Title     string
	DependsOn []string
}

type Priority string

const (
	PriorityHigh   Priority = "high"
	PriorityMedium Priority = "medium"
	PriorityLow    Priority = "low"
)
