package llm

import "context"

// DummyClientは固定のJSONを返すテスト用
type DummyClient struct {
	Reply string
}

func (d DummyClient) Generate(ctx context.Context, req Request) (Response, error) {
	return Response{
		Text: d.Reply,
	}, nil
}
