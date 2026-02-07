package orchestrator

import "github.com/TonChan8028/ai-project-orchestrator/internal/domain"

// Resultはorchestratorの最終アウトプット
// 成功失敗ではなく進行可能かを示す
type Result struct {
	State         domain.ProjectState
	CanProceed    bool
	OpenQuestion  []domain.OpenQuestion
	CriticalIssue []domain.ReviewFinding
}
