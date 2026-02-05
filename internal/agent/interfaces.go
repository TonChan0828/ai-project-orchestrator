package agent

import (
	"context"

	"github.com/TonChan8028/ai-project-orchestrator/internal/domain"
)

// PMAgent: overview->requirements
type PMAgent interface {
	Run(ctx context.Context, in domain.ProjectOverview) (domain.Requirements, error)
}

// ArchitectAgent: requirements->architecture
type ArchitectAAgent interface {
	Run(ctx context.Context, in domain.Requirements) (domain.Architecture, error)
}

// PlannerAgent:(requirements+architecture)->plan
type PlannerAgent interface {
	Run(ctx context.Context, req domain.Requirements, arch domain.Architecture) (domain.Plan, error)
}

// RiskAgent: state->risks+open questions
type RiskAgent interface {
	Run(ctx context.Context, state domain.ProjectState) ([]domain.Risk, []domain.OpenQuestion, error)
}

// ReviewerAgent: state->findings
type ReviewerAgent interface {
	Run(ctx context.Context, state domain.ProjectState) ([]domain.ReviewFinding, error)
}
