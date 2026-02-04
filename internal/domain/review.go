package domain

type ReviewFinding struct {
	Severity    Severity
	Description string
	Target      ReviewTarget
}

type Severity string

const (
	SeverityCritical Severity = "critical"
	SeverityMajor    Severity = "major"
	SeverityMinor    Severity = "minor"
)

type ReviewTarget string

const (
	TargetRequirements ReviewTarget = "requirements"
	TargetArchitecture ReviewTarget = "architecture"
	TargetPlan         ReviewTarget = "plan"
)
