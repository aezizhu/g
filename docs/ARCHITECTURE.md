Architecture
============

Author: AZ <Aezi.zhu@icloud.com>

Overview
--------

The system converts natural language input into a structured command plan, validates it with a policy engine, displays it to the user, and executes it without invoking a shell. The design prioritizes safety, determinism, and maintainability.

Components
----------

- CLI (`cmd/lucicodex`): Parses flags, loads config, orchestrates request/plan/execute.
- Config (`internal/config`): Loads defaults, JSON file, UCI (OpenWrt), and env.
- Planner (`internal/plan`): Defines the plan schema and instruction prefix.
- LLM Client (`internal/llm`): Calls provider HTTP API (Gemini) and parses plan.
- Policy (`internal/policy`): Allow/Deny checks, shell metacharacter checks.
- Executor (`internal/executor`): Runs argv-only commands with timeouts and minimal env.
- UI (`internal/ui`): Renders plans and results, prompts for confirmation.

Data Flow
---------

```
User -> CLI -> Instruction + Prompt -> LLM -> JSON Plan -> Policy -> UI -> Exec -> Results
```

Safety Principles
-----------------

- No shell execution; only execve-style argv.
- Explicit allowlist/denylist checks.
- Per-command timeouts and minimal environment.
- Human-in-the-loop confirmation by default.

Extensibility
-------------

- Providers: add new clients under `internal/llm/` implementing a `GeneratePlan`-like method.
- Policies: extend allow/deny lists via config or add advanced validators.
- OpenWrt tools: add wrappers under `internal/openwrt/` to enrich prompts.


