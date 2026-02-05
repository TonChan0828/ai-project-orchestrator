package agent

import (
	"context"

	"github.com/TonChan8028/ai-project-orchestrator/internal/domain"
)

type PlannerStub struct{}

func (p PlannerStub) Run(ctx context.Context, req domain.Requirements, arch domain.Architecture) (domain.Plan, error) {
	return domain.Plan{
		Epics: []domain.Epic{
			{
				Name: "MVP骨格",
				Stories: []domain.Story{
					{
						Title: "ProhectStateとAgent契約を確定する",
						Tasks: []domain.Task{
							{Title: "domain型定義", DependsOn: nil, Priority: domain.PriorityHigh},
							{Title: "agent interface定義", DependsOn: []string{"domain型定義"}, Priority: domain.PriorityHigh},
						},
					},
				},
			},
		},
	}, nil
}
