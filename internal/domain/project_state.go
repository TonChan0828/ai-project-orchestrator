package domain

type ProjectState struct {
	Overview       *ProjectOverview
	Requirements   *Requirements
	Architecture   *Architecture
	Plan           *Plan
	Risks          *Risk
	OpenQuestions  []OpenQuestion
	ReviewFindings []ReviewFinding
}

// precondition helpers
// Orchestratorが「次に何を実行できるか」を判断するための関数

func (ps *ProjectState) ReadyForArchitecture() bool {
	return ps.Requirements != nil
}

func (ps *ProjectState) ReadyForPlanning() bool {
	return ps.Requirements != nil && ps.Architecture != nil
}

func (ps *ProjectState) ReadyForReview() bool {
	return ps.Requirements != nil && ps.Architecture != nil
}
