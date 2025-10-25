package llm

import (
    "bytes"
    "context"
    "encoding/json"
    "errors"
    "fmt"
    "io"
    "net/http"
    "time"

    "github.com/aezizhu/g/internal/config"
    "github.com/aezizhu/g/internal/plan"
)

type AnthropicClient struct {
    httpClient *http.Client
    cfg        config.Config
}

func NewAnthropicClient(cfg config.Config) *AnthropicClient {
    return &AnthropicClient{httpClient: &http.Client{Timeout: 30 * time.Second}, cfg: cfg}
}

type anthropicMessage struct {
    Role    string `json:"role"`
    Content string `json:"content"`
}

type anthropicReq struct {
    Model     string              `json:"model"`
    Messages  []anthropicMessage  `json:"messages"`
    MaxTokens int                 `json:"max_tokens"`
}

type anthropicResp struct { Content []struct{ Text string `json:"text"` } `json:"content"` }

func (c *AnthropicClient) GeneratePlan(ctx context.Context, prompt string) (plan.Plan, error) {
    var zero plan.Plan
    if c.cfg.AnthropicAPIKey == "" {
        return zero, errors.New("missing ANTHROPIC_API_KEY")
    }
    model := c.cfg.Model
    if model == "" {
        model = "claude-3-5-sonnet-20240620"
    }
    body := anthropicReq{Model: model, MaxTokens: 2048}
    body.Messages = []anthropicMessage{{Role: "user", Content: prompt}}
    b, _ := json.Marshal(body)
    req, _ := http.NewRequestWithContext(ctx, http.MethodPost, "https://api.anthropic.com/v1/messages", bytes.NewReader(b))
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("x-api-key", c.cfg.AnthropicAPIKey)
    req.Header.Set("anthropic-version", "2023-06-01")
    resp, err := c.httpClient.Do(req)
    if err != nil { return zero, err }
    defer resp.Body.Close()
    if resp.StatusCode < 200 || resp.StatusCode >= 300 {
        data, _ := io.ReadAll(resp.Body)
        return zero, fmt.Errorf("anthropic http %d: %s", resp.StatusCode, string(data))
    }
    var ar anthropicResp
    if err := json.NewDecoder(resp.Body).Decode(&ar); err != nil { return zero, err }
    if len(ar.Content) == 0 { return zero, errors.New("empty response") }
    text := ar.Content[0].Text
    return plan.TryUnmarshalPlan(text)
}


