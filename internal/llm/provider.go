package llm

import (
    "context"

    "github.com/aezizhu/g/internal/config"
    "github.com/aezizhu/g/internal/plan"
)

// Provider is the interface implemented by LLM clients that can produce plans.
type Provider interface {
    GeneratePlan(ctx context.Context, prompt string) (plan.Plan, error)
}

// NewProvider returns a Provider based on configuration.
func NewProvider(cfg config.Config) Provider {
    switch cfg.Provider {
    case "openai":
        return NewOpenAIClient(cfg)
    case "anthropic":
        return NewAnthropicClient(cfg)
    case "gemini-cli":
        return NewExternalGeminiClient(cfg)
    default:
        return NewGeminiClient(cfg)
    }
}


