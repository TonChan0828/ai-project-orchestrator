package agent

import (
	"context"

	"github.com/TonChan8028/ai-project-orchestrator/internal/domain"
)

type ArchitectStub struct{}

func (a ArchitectStub) Run(ctx context.Context, in domain.Requirements) (domain.Architecture, error) {
	return domain.Architecture{
		Components: []domain.Component{
			{Name: "orchestrator", Responsibility: "状態管理と実行制御", DependsOn: []string{"PMAgent", "Architect Agent", "Planner Agent", "Risk Agent", "Reviewer Agent"}},
		},
		DataFlows: []domain.DataFlow{
			{From: "User", To: "PM Agent", Description: "プロジェクト概要入力"},
			{From: "Orchestrator", To: "PM Agent", Description: "概要を渡して要件抽出"},
			{From: "PM Agent", To: "Orchestrator", Description: "Requirementsを返却しProjectStateに格納"},
		},
	}, nil
}
