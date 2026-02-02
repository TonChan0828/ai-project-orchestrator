# 目的 / 非目的（MVPスコープ）
## 目的（MVP）

本システムの目的は、
ユーザーが入力した「作りたいプロジェクト概要」をもとに、
人間が次に取るべき行動が明確になる設計成果物を生成することである。
具体的には以下を生成対象とする。
- 要件定義（機能要件・非機能要件・前提・制約）
- アーキテクチャ設計（構成要素・責務・データフロー）
- タスク分解（WBS / チケット粒度）
- リスク・不確実性の洗い出し
- 未確定事項に対する質問
- 上記全体に対する矛盾・漏れの指摘
重要なのは 「決定」ではなく「判断材料の提示」 である。

## 非目的
- プロジェクトの自動実行・自動実装
- 外部API・DB・チケット管理ツールとの連携
- プロンプトチューニングによる品質向上
- エージェント同士の自由会話
- 正解」を1つに決めること
  
# 用語定義（ProjectState / Agent / Orchestrator など）
### ProjectState

プロジェクトに関する すべての中間成果物を集約した状態オブジェクト。
このシステムにおける 唯一の真実（Single Source of Truth）。
LLMの出力はすべて ProjectState に格納され、
LLMの発言＝事実 とは扱わない。

### Agent
特定の役割を持ち、
入力を受け取り、構造化された出力を返す純粋関数的存在。
状態を保持しない
他のAgentと直接通信しない
判断の最終責任を持たない

### Orchestrator
システムの中央制御装置。
ProjectState の生成・更新を管理
Agent の実行順序を制御
出力の妥当性チェックとゲート判定を行う
「次に人間がやること」をまとめる

# 全体アーキテクチャ（中央制御・会話禁止の理由）
### 中央制御（Orchestrator）を置く理由
LLMは以下の性質を持つ。
- 一貫性が保証されない
- 出力の正当性を自己検証できない
- 状態管理が苦手

そのため、
判断と状態管理を LLM から切り離し、
Orchestrator が担う構造を採用する。

### エージェント同士を会話させない理由
エージェント同士が直接会話すると：
- 誰がどの根拠で決めたか追えなくなる
- 再現性が失われる
- デバッグ不能になる

そのため、
すべてのやり取りは ProjectState を介して行う。

# ProjectState（フィールド一覧と責務）
ProjectState は「途中成果物の箱」であり、
未完成・矛盾・不確実性を含むことを前提とする。

### 主なフィールドと責務
#### ProjectOverview
- ユーザー入力原文
- 前提条件の明文化
- 用語の暫定定義

#### Requirements
- 機能要件
- 非機能要件
- 制約条件
- スコープ外事項

#### Architecture
- コンポーネント構成
- 責務分割
- データフロー（文章レベル）

#### API / DataModel
- 想定エンドポイント
- 認証方針
- 主要エンティティと関係

#### Plan
- Epic / Story / Task
- 依存関係
- 優先度

#### Risks
- 技術的リスク
- 要件不確実性
- 運用リスク
- 対策案・検知方法

#### OpenQuestions
- 判断不能な事項
- ユーザーに確認すべき質問

#### ReviewFindings
- 矛盾点
- 抜け漏れ
- 曖昧な表現

# エージェント定義（責務・入力・出力）
### PM Agent
入力：ProjectOverview

出力：Requirements

責務：要求の構造化、前提と用語の整理

### Architect Agent

入力：Requirements

出力：Architecture / API / DataModel

責務：構造的に破綻しない設計案の提示

### Planner Agent

入力：Requirements + Architecture

出力：Plan

責務：実装可能な粒度への分解

### Risk Agent

入力：ProjectState

出力：Risks / OpenQuestions

責務：不確実性の可視化

### Reviewer Agent

入力：ProjectState

出力：ReviewFindings

責務：整合性チェック（判断はしない）

# 実行フロー（順序とゲート）
1. PM Agent → Requirements生成
2. Architect Agent → 設計生成
3. Planner Agent → タスク分解
4. Risk Agent → リスク抽出
5. Reviewer Agent → 矛盾・漏れ検出
6. Orchestrator → 最終レポート生成

Reviewer の結果は ゲートとして扱い、
重大な指摘があれば「未確定」として人間に返す。

# 品質担保（構造化、バリデーション、レビュー）
- 出力はすべて 構造化データ前提
- 「正しさ」ではなく 検証可能性 を重視
- 品質はプロンプトではなく 工程分離で担保する

# 失敗設計（LLM失敗＝状態として扱う）
LLMの失敗は例外ではない。
- JSON崩れ
- 矛盾の多発
- 内容が薄い

これらは ProjectState に「失敗状態」として記録し、
OpenQuestions として人間に返す。

# 将来拡張（外部ツール、検索、RAG、チケット連携）
- Web検索・RAG Agent
- コード生成 Agent
- チケット管理ツール連携
- 実装フェーズ専用 Orchestrator