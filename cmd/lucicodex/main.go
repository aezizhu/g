package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/aezizhu/g/internal/config"
	"github.com/aezizhu/g/internal/executor"
	"github.com/aezizhu/g/internal/llm"
	"github.com/aezizhu/g/internal/logging"
	"github.com/aezizhu/g/internal/openwrt"
	"github.com/aezizhu/g/internal/plan"
	"github.com/aezizhu/g/internal/policy"
	"github.com/aezizhu/g/internal/repl"
	"github.com/aezizhu/g/internal/ui"
	"github.com/aezizhu/g/internal/wizard"
)

const version = "0.3.0"

func main() {
	var (
		configPath  = flag.String("config", "", "path to JSON config file")
		model       = flag.String("model", "", "model name")
		provider    = flag.String("provider", "", "provider name (gemini, openai, anthropic, gemini-cli)")
		dryRun      = flag.Bool("dry-run", true, "only print plan, do not execute")
		approve     = flag.Bool("approve", false, "auto-approve plan without confirmation")
		confirmEach = flag.Bool("confirm-each", false, "confirm each command before execution")
		timeout     = flag.Int("timeout", 0, "per-command timeout in seconds")
		maxCommands = flag.Int("max-commands", 0, "maximum number of commands to execute")
		logFile     = flag.String("log-file", "", "log file path")
		showVersion = flag.Bool("version", false, "print version and exit")
		jsonOutput  = flag.Bool("json", false, "emit JSON output for plan and results")
		facts       = flag.Bool("facts", true, "include environment facts in prompt")
		interactive = flag.Bool("interactive", false, "start interactive REPL mode")
		setup       = flag.Bool("setup", false, "run setup wizard")
	)

	flag.Parse()

	if *showVersion {
		fmt.Printf("LuCICodex version %s\n", version)
		os.Exit(0)
	}

	cfg, err := config.Load(*configPath)
	if err != nil {
		if !*setup {
			fmt.Fprintf(os.Stderr, "Configuration error: %v\n", err)
			fmt.Fprintf(os.Stderr, "Run with -setup to configure LuCICodex\n")
			os.Exit(1)
		}
		cfg = config.Config{}
	}

	if *model != "" {
		cfg.Model = *model
	}
	if *provider != "" {
		cfg.Provider = *provider
	}
	if *timeout > 0 {
		cfg.TimeoutSeconds = *timeout
	}
	if *maxCommands > 0 {
		cfg.MaxCommands = *maxCommands
	}
	if *logFile != "" {
		cfg.LogFile = *logFile
	}
	cfg.DryRun = *dryRun
	cfg.AutoApprove = *approve

	if *setup {
		w := wizard.New(os.Stdin, os.Stdout)
		if err := w.Run(); err != nil {
			fmt.Fprintf(os.Stderr, "Setup error: %v\n", err)
			os.Exit(1)
		}
		os.Exit(0)
	}

	if *interactive {
		r := repl.New(cfg)
		ctx := context.Background()
		if err := r.Run(ctx, os.Stdin, os.Stdout); err != nil {
			fmt.Fprintf(os.Stderr, "REPL error: %v\n", err)
			os.Exit(1)
		}
		os.Exit(0)
	}

	args := flag.Args()
	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "Usage: lucicodex [flags] <prompt>\n")
		fmt.Fprintf(os.Stderr, "Run 'lucicodex -h' for help\n")
		os.Exit(1)
	}

	prompt := args[0]
	ctx := context.Background()

	llmProvider := llm.NewProvider(cfg)
	policyEngine := policy.New(cfg)
	execEngine := executor.New(cfg)
	logger := logging.New(cfg.LogFile)

	instruction := plan.BuildInstructionWithLimit(cfg.MaxCommands)
	if *facts {
		factsCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
		defer cancel()
		envFacts := openwrt.CollectFacts(factsCtx)
		if envFacts != "" {
			instruction += "\n\nEnvironment facts (read-only):\n" + envFacts
		}
	}

	fullPrompt := instruction + "\n\nUser request: " + prompt

	// Generate plan
	planCtx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	p, err := llmProvider.GeneratePlan(planCtx, fullPrompt)
	if err != nil {
		fmt.Fprintf(os.Stderr, "LLM error: %v\n", err)
		os.Exit(1)
	}

	if len(p.Commands) == 0 {
		fmt.Println("No commands proposed.")
		os.Exit(0)
	}

	if cfg.MaxCommands > 0 && len(p.Commands) > cfg.MaxCommands {
		p.Commands = p.Commands[:cfg.MaxCommands]
	}

	// Validate plan
	if err := policyEngine.ValidatePlan(p); err != nil {
		fmt.Fprintf(os.Stderr, "Plan rejected by policy: %v\n", err)
		os.Exit(1)
	}

	if *jsonOutput {
		if err := ui.PrintPlanJSON(os.Stdout, p); err != nil {
			fmt.Fprintf(os.Stderr, "JSON output error: %v\n", err)
			os.Exit(1)
		}
	} else {
		ui.PrintPlan(os.Stdout, p)
	}

	logger.Plan(prompt, p)

	if cfg.DryRun {
		fmt.Println("\nDry run mode - no execution")
		os.Exit(0)
	}

	if !cfg.AutoApprove {
		reader := bufio.NewReader(os.Stdin)
		ok, err := ui.Confirm(reader, os.Stdout, "Execute these commands?")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Confirmation error: %v\n", err)
			os.Exit(1)
		}
		if !ok {
			fmt.Println("Cancelled")
			os.Exit(0)
		}
	}

	var results executor.Results
	if *confirmEach {
		reader := bufio.NewReader(os.Stdin)
		for i, cmd := range p.Commands {
			fmt.Printf("\nExecute command %d: %s\n", i+1, executor.FormatCommand(cmd.Command))
			ok, err := ui.Confirm(reader, os.Stdout, "Proceed?")
			if err != nil || !ok {
				fmt.Println("Skipped")
				continue
			}
			result := execEngine.RunCommand(ctx, i, cmd)
			results.Items = append(results.Items, result)
			if result.Err != nil {
				results.Failed++
			}
		}
	} else {
		results = execEngine.RunPlan(ctx, p)
	}

	if *jsonOutput {
		if err := ui.PrintResultsJSON(os.Stdout, results); err != nil {
			fmt.Fprintf(os.Stderr, "JSON output error: %v\n", err)
			os.Exit(1)
		}
	} else {
		ui.PrintResults(os.Stdout, results)
	}

	items := make([]logging.ResultItem, 0, len(results.Items))
	for _, it := range results.Items {
		errStr := ""
		if it.Err != nil {
			errStr = it.Err.Error()
		}
		items = append(items, logging.ResultItem{
			Index:   it.Index,
			Command: it.Command,
			Output:  it.Output,
			Error:   errStr,
			Elapsed: it.Elapsed,
		})
	}
	logger.Results(items)

	if results.Failed > 0 {
		os.Exit(1)
	}
}
