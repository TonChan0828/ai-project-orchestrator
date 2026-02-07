package orchestrator

import (
	"context"
	"errors"

	"github.com/TonChan8028/ai-project-orchestrator/internal/agent"
	"github.com/TonChan8028/ai-project-orchestrator/internal/domain"
)

type Orchestrator struct {
	PM        agent.PMAgent
	Architect agent.ArchitectAgent
	Planner   agent.PlannerAgent
	Risk      agent.RiskAgent
	Reviewer  agent.ReviewerAgent
}

func New(
	pm agent.PMAgent,
	arch agent.ArchitectAgent,
	planner agent.PlannerAgent,
	risk agent.RiskAgent,
	reviewer agent.ReviewerAgent,
) *Orchestrator {
	return &Orchestrator{
		PM:        pm,
		Architect: arch,
		Planner:   planner,
		Risk:      risk,
		Reviewer:  reviewer,
	}
}

func (o *Orchestrator) Run(ctx context.Context, overview domain.ProjectOverview) (*Result, error) {
	state := domain.ProjectState{
		Overview: &overview,
	}

	// 1. PM Agent
	req, err := o.PM.Run(ctx, overview)
	if err != nil {
		return nil, err
	}
	state.Requirements = &req

	// 2. Architect Agent
	if !state.ReadyForArchitecture() {
		return nil, errors.New("architecture precondition not met")
	}
	arch, err := o.Architect.Run(ctx, req)
	if err != nil {
		return nil, err
	}
	state.Architecture = &arch

	// 3. Planner Agent
	if !state.ReadyForPlanning() {
		return nil, errors.New("planning precondition not met")
	}
	plan, err := o.Planner.Run(ctx, req, arch)
	if err != nil {
		return nil, err
	}
	state.Plan = &plan

	// 4. Risk Agent
	risks, questions, err := o.Risk.Run(ctx, state)
	if err != nil {
		return nil, err
	}
	state.Risks = risks
	state.OpenQuestions = questions

	// 5. Reviewer Agent
	findings, err := o.Reviewer.Run(ctx, state)
	if err != nil {
		return nil, err
	}
	state.ReviewFindings = findings

	// ゲート判定
	critical := filerCritical(findings)
	if len(critical) > 0 {
		return &Result{
			State:         state,
			CanProceed:    false,
			OpenQuestion:  state.OpenQuestions,
			CriticalIssue: critical,
		}, nil
	}

	// 進行可能
	return &Result{
		State:      state,
		CanProceed: true,
	}, nil
}

func filerCritical(findings []domain.ReviewFinding) []domain.ReviewFinding {
	var critical []domain.ReviewFinding
	for _, f := range findings {
		if f.Severity == domain.SeverityCritical {
			critical = append(critical, f)
		}
	}
	return critical
}
