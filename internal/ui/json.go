package ui

import (
    "encoding/json"
    "io"

    "github.com/aezizhu/g/internal/executor"
    "github.com/aezizhu/g/internal/plan"
)

func PrintPlanJSON(w io.Writer, p plan.Plan) error {
    enc := json.NewEncoder(w)
    enc.SetIndent("", "  ")
    return enc.Encode(p)
}

func PrintResultsJSON(w io.Writer, res executor.Results) error {
    enc := json.NewEncoder(w)
    enc.SetIndent("", "  ")
    return enc.Encode(res)
}


