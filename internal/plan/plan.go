package plan

import (
    "encoding/json"
    "fmt"
    "strings"
)

// PlannedCommand represents a single command to execute safely without shell interpolation.
type PlannedCommand struct {
    Command     []string `json:"command"`
    Description string   `json:"description,omitempty"`
    NeedsRoot   bool     `json:"needs_root,omitempty"`
}

// Plan is the structured response expected from the model.
type Plan struct {
    Summary  string           `json:"summary,omitempty"`
    Commands []PlannedCommand `json:"commands"`
    Warnings []string         `json:"warnings,omitempty"`
}

// BuildInstruction returns the instruction prefix to reliably elicit a JSON plan.
func BuildInstruction(cfg interface{}) string {
    // Keep instruction concise and deterministic.
    b := &strings.Builder{}
    b.WriteString("You are a router command planner.\n")
    b.WriteString("Output only strict JSON that conforms to this schema:\n")
    b.WriteString("{\n  \"summary\": string,\n  \"commands\": [ { \"command\": [string, ...], \"description\": string, \"needs_root\": bool } ],\n  \"warnings\": [string]\n}\n")
    b.WriteString("Rules:\n")
    b.WriteString("- Use explicit argv arrays; do not return shell pipelines or redirections.\n")
    b.WriteString("- Prefer OpenWrt tools: uci, ubus, fw4, opkg, logread, dmesg.\n")
    b.WriteString("- Limit commands to safe, idempotent operations when possible.\n")
    b.WriteString("- Keep the commands minimal and directly actionable.\n")
    return b.String()
}

// BuildInstructionWithLimit adds a hint for maximum number of commands.
func BuildInstructionWithLimit(maxCommands int) string {
    base := BuildInstruction(nil)
    if maxCommands > 0 {
        return base + "\nDo not return more than " + fmt.Sprint(maxCommands) + " commands."
    }
    return base
}

// TryUnmarshalPlan attempts to decode a JSON string to Plan.
func TryUnmarshalPlan(s string) (Plan, error) {
    var p Plan
    err := json.Unmarshal([]byte(s), &p)
    return p, err
}


