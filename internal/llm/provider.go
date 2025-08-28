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
    // For now only Gemini is implemented.
    return NewGeminiClient(cfg)
}


