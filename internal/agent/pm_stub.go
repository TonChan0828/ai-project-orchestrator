package agent

import (
	"context"

	"github.com/TonChan8028/ai-project-orchestrator/internal/domain"
)

type PMStub struct{}

func (p PMStub) Run(ctx context.Context, in domain.ProjectOverview) (domain.Requirements, error) {
	return domain.Requirements{
		Functional: []domain.RequirementItem{
			{ID: "F-1", Description: "ユーザの入力から要件を構造化する", Rationale: "次工程(設計/タスク分解)を可能とするため"},
		},
		NonFunctional: []domain.RequirementItem{
			{ID: "NF-1", Description: "出力は構造化され検証可能である", Rationale: "LLMののブレをシステムで抑えるため"},
		},
		Constraints: []string{"LLMは判断主体にならない", "エージェント同士は直接会話しない"},
		OutOfScore:  []string{"自動実装", "外部ツール連携"},
	}, nil
}
