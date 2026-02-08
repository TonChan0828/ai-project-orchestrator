package llm

type Request struct {
	System      string
	User        string
	MaxTokens   int
	Temperature float64
	TraceID     string
}

type Response struct {
	Text string
	Raw  any
}
