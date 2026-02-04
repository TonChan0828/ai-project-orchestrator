package domain

type Risk struct {
	Type        RiskType
	Description string
	Mitigation  string
	Detection   string
}

type RiskType string

const (
	RiskTechnical   RiskType = "technical"
	RiskRequirement RiskType = "requirement"
	RiskOperational RiskType = "operational"
)

type OpenQuestion struct {
	Question string
	Reason   string
}
