package main

import (
	"context"
	"fmt"

	"github.com/TonChan8028/ai-project-orchestrator/internal/agent"
	"github.com/TonChan8028/ai-project-orchestrator/internal/domain"
	"github.com/TonChan8028/ai-project-orchestrator/internal/llm"
	"github.com/TonChan8028/ai-project-orchestrator/internal/orchestrator"
)

func main() {
	ctx := context.Background()

	fmt.Println("=== Case1: 正常系（すべて進行可能） ===")
	runHappyPath(ctx)

	fmt.Println("\n=== Case2: LLMが壊れたJSONを返す ===")
	runBrokenJSON(ctx)

	fmt.Println("\n=== Case3: Reviewerがcriticalを返す ===")
	runCriticalReview(ctx)
}

func runHappyPath(ctx context.Context) {
	client := llm.DummyClient{
		Reply: `{
			"Functional":[
				{"ID":"F-1","Description":"要件を構造化する","Rationale":"設計と計画の入力にするため"}
			],
			"NonFunctional":[
				{"ID":"NF-1","Description":"構造化出力","Rationale":"検証可能性を確保するため"}
			],
			"Constraints":["エージェント間会話禁止"],
			"OutOfScope":["自動実装"]
		}`,
	}

	// LLM版PM Agent（最小）
	pm := agentLLMPM(client)

	orch := orchestrator.New(
		pm,
		agent.ArchitectStub{},
		agent.PlannerStub{},
		agent.RiskStub{},
		agent.ReviewerStub{}, // majorのみ → 通過
	)

	overview := domain.ProjectOverview{
		OriginalText: "AIエージェント設計を学びたい",
	}

	result, err := orch.Run(ctx, overview)
	if err != nil {
		panic(err)
	}

	fmt.Println("CanProceed:", result.CanProceed)
	fmt.Println("Requirements count:", len(result.State.Requirements.Functional))
}

func runBrokenJSON(ctx context.Context) {
	client := llm.DummyClient{
		Reply: `
		以下が要件です！

		{
			"Functional":[
				{"ID":"F-1","Description":"途中で終わる"
		`,
	}

	pm := agentLLMPM(client)

	orch := orchestrator.New(
		pm,
		agent.ArchitectStub{},
		agent.PlannerStub{},
		agent.RiskStub{},
		agent.ReviewerStub{},
	)

	overview := domain.ProjectOverview{
		OriginalText: "JSONが壊れるケース",
	}

	result, err := orch.Run(ctx, overview)

	fmt.Println("error:", err)
	fmt.Println("result is nil:", result == nil)
}

type CriticalReviewerStub struct{}

func (c CriticalReviewerStub) Run(ctx context.Context, state domain.ProjectState) ([]domain.ReviewFinding, error) {
	return []domain.ReviewFinding{
		{
			Severity:    domain.SeverityCritical,
			Description: "要件と設計が矛盾している",
			Target:      domain.TargetArchitecture,
		},
	}, nil
}

func runCriticalReview(ctx context.Context) {
	pm := agent.PMStub{}

	orch := orchestrator.New(
		pm,
		agent.ArchitectStub{},
		agent.PlannerStub{},
		agent.RiskStub{},
		CriticalReviewerStub{}, // ← 差し替え
	)

	overview := domain.ProjectOverview{
		OriginalText: "Reviewerが止めるケース",
	}

	result, err := orch.Run(ctx, overview)
	if err != nil {
		panic(err)
	}

	fmt.Println("CanProceed:", result.CanProceed)
	fmt.Println("CriticalIssue:", len(result.CriticalIssue))
}

func agentLLMPM(client llm.Client) agent.PMAgent {
	return pmLLM{client: client}
}

type pmLLM struct {
	client llm.Client
}

func (p pmLLM) Run(ctx context.Context, in domain.ProjectOverview) (domain.Requirements, error) {
	system := `
あなたはPM Agentです。
出力は JSON のみ。
以下の型に必ず一致させてください。
説明文は禁止。
`

	return llm.GenerateJSON[domain.Requirements](
		ctx,
		p.client,
		system,
		in.OriginalText,
		"trace-pm-001",
	)
}
