package agent

import (
	"context"

	"github.com/TonChan8028/ai-project-orchestrator/internal/domain"
)

type RiskStub struct{}

func (r RiskStub) Run(ctx context.Context, state domain.ProjectState) ([]domain.Risk, []domain.OpenQuestion, error) {
	risks := []domain.Risk{
		{
			Type:        domain.RiskRequirement,
			Description: "ユーザ入力が曖昧で要件が確定できない可能性",
			Mitigation:  "OpenQuestionを生成し、未確定を明示する",
			Detection:   "Requirementsに曖昧語(例：適切に、いい感じに)が多い",
		},
	}
	qs := []domain.OpenQuestion{
		{Question: "対象ユーザは誰ですか？(社内/個人/一般)", Reason: "非機能要件とUI設計が変わるため"},
	}

	return risks, qs, nil
}
