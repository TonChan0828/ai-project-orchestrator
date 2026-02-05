package agent

import (
	"context"

	"github.com/TonChan8028/ai-project-orchestrator/internal/domain"
)

type ReviewerStub struct{}

func (r ReviewerStub) Run(ctx context.Context, state domain.ProjectState) ([]domain.ReviewFinding, error) {
	// スタブなので「重大指摘が出る」例も入れてゲート動作を確認できるようにする
	findings := []domain.ReviewFinding{
		{
			Severity:    domain.SeverityMajor,
			Description: "非機能要件に性能・セキュリティ要件が不足している可能性",
			Target:      domain.TargetRequirements,
		},
	}
	return findings, nil
}
