package agent_llm

import (
	"context"

	"github.com/TonChan8028/ai-project-orchestrator/internal/domain"
	"github.com/TonChan8028/ai-project-orchestrator/internal/llm"
)

// PMAgentLLMはLLMを使ったPM Agent実装
// Agent契約（I/O）はStep1から一切変更しない
type PMAgentLLM struct {
	Client llm.Client
}

func NewPMAgentLLM(client llm.Client) PMAgentLLM {
	return PMAgentLLM{Client: client}
}

// Run implements agent.PMAgent
func (p PMAgentLLM) Run(
	ctx context.Context,
	in domain.ProjectOverview,
) (domain.Requirements, error) {
	system := `
	あなたはPM Agentです。
	以下の制約を必ず守ってください。

	- 出力は JSON のみ
	- 説明文・前置き・コメントは禁止
	- 指定された型に必ず一致させること
	- すべての RequirementItem に Rationale を含めること

	JSONスキーマ（厳守）:
	{
	"Functional": [
		{ "ID": "string", "Description": "string", "Rationale": "string" }
	],
	"NonFunctional": [
		{ "ID": "string", "Description": "string", "Rationale": "string" }
	],
	"Constraints": ["string"],
	"OutOfScope": ["string"]
	}
	`

	return llm.GenerateJSON[domain.Requirements](
		ctx,
		p.Client,
		system,
		in.OriginalText,
		"trace-pm-llm",
	)
}
