package llm

import (
	"context"
	"encoding/json"
	"strings"
	"time"
)

// GenerateJSONはLLMにJSONのみを返させ、型Tへデコードする
func GenerateJSON[T any](
	ctx context.Context,
	client Client,
	system string,
	user string,
	traceID string,
) (T, error) {
	var zero T

	const maxAttempts = 3
	sys := system

	for attempt := 1; attempt <= maxAttempts; attempt++ {
		req := Request{
			System:      sys,
			User:        user,
			MaxTokens:   1200,
			Temperature: 0.1,
			TraceID:     traceID,
		}

		res, err := client.Generate(ctx, req)
		if err != nil {
			// 通信などは一時的とみなす
			if attempt < maxAttempts {
				time.Sleep(backoff(attempt))
				continue
			}
			return zero, &Error{Kind: ErrTransient, Msg: "llm generate failed", Cause: err}
		}

		jsonText, ok := extractJSON(res.Text)
		if !ok {
			if attempt < maxAttempts {
				sys = tightenSystem(sys, "JSONが検出できませんでした。JSONのみを返してください。")
				continue
			}
		}

		var out T
		if err := json.Unmarshal([]byte(jsonText), &out); err != nil {
			if attempt < maxAttempts {
				sys = tightenSystem(sys, "JSONの形式が不正でした。説明文は禁止。JSONのみ。")
				continue
			}
			return zero, &Error{Kind: ErrInvalid, Msg: "json unmarshal failed", Cause: err}
		}

		return out, nil
	}
	return zero, &Error{Kind: ErrUnknown, Msg: "unexpected"}
}

func backoff(attempt int) time.Duration {
	switch attempt {
	case 1:
		return 200 * time.Millisecond
	case 2:
		return 500 * time.Millisecond
	default:
		return 800 * time.Millisecond
	}
}

// extractJSONはテキストから最初のJSONオブジェクト/配列を抜き出す
func extractJSON(s string) (string, bool) {
	startObj := strings.Index(s, "{")
	startArr := strings.Index(s, "[")

	start := -1
	if startObj >= 0 && startArr >= 0 {
		if startObj < startArr {
			start = startObj
		} else {
			start = startArr
		}
	} else if startObj >= 0 {
		start = startObj
	} else if startArr >= 0 {
		start = startArr
	}

	if start < 0 {
		return "", false
	}
	// 対応する末尾まで切る
	sub := s[start:]
	return sub, true
}

func tightenSystem(prev, hint string) string {
	if strings.Contains(prev, hint) {
		return prev
	}
	return prev + "\n\n" + hint
}
