package main

import (
	"context"
	"fmt"

	"github.com/TonChan8028/ai-project-orchestrator/internal/agent"
	"github.com/TonChan8028/ai-project-orchestrator/internal/agent_llm"
	"github.com/TonChan8028/ai-project-orchestrator/internal/domain"
	"github.com/TonChan8028/ai-project-orchestrator/internal/llm"
	"github.com/TonChan8028/ai-project-orchestrator/internal/orchestrator"
)

func main() {
	ctx := context.Background()

	fmt.Println("=== Step2-2: PM Agent LLM版 動作確認 ===")

	// ダミーLLM（正しいJSONを返す）
	client := llm.DummyClient{
		Reply: `{
			"Functional": [
				{
					"ID": "F-1",
					"Description": "ユーザーの入力から要件を構造化する",
					"Rationale": "設計とタスク分解の入力にするため"
				}
			],
			"NonFunctional": [
				{
					"ID": "NF-1",
					"Description": "出力は検証可能な構造化形式とする",
					"Rationale": "LLMの幻覚を検知できるようにするため"
				}
			],
			"Constraints": [
				"エージェント同士は直接会話しない"
			],
			"OutOfScope": [
				"自動実装"
			]
		}`,
	}

	// ★ PM だけ LLM版に差し替え
	pm := agent_llm.NewPMAgentLLM(client)

	orch := orchestrator.New(
		pm,                    // ← LLM版
		agent.ArchitectStub{}, // 既存Stub
		agent.PlannerStub{},
		agent.RiskStub{},
		agent.ReviewerStub{},
	)

	overview := domain.ProjectOverview{
		OriginalText: `
AIエージェント設計を学習するための
マルチエージェント構成の学習用プロジェクトを作りたい。
`,
	}

	result, err := orch.Run(ctx, overview)
	if err != nil {
		panic(err)
	}

	fmt.Println("CanProceed:", result.CanProceed)
	fmt.Println("Functional Requirements:", len(result.State.Requirements.Functional))
	fmt.Println("NonFunctional Requirements:", len(result.State.Requirements.NonFunctional))
}
