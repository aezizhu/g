g - Natural Language CLI for OpenWrt
====================================

Author: aezizhu

<p align="center">
  <a href="#"><img alt="Build" src="https://img.shields.io/badge/build-passing-brightgreen"></a>
  <a href="#"><img alt="License" src="https://img.shields.io/badge/license-MIT-blue"></a>
  <a href="#"><img alt="Go Version" src="https://img.shields.io/badge/Go-1.21+-1f425f"></a>
  <a href="#"><img alt="OpenWrt" src="https://img.shields.io/badge/OpenWrt-supported-00a0e9"></a>
  <a href="https://github.com/aezizhu/g/actions/workflows/build.yml"><img alt="CI" src="https://github.com/aezizhu/g/actions/workflows/build.yml/badge.svg"></a>
</p>

Overview
--------

g is a secure, extensible command-line utility that translates natural language requests into audited shell commands on headless Linux systems, with first-class support for OpenWrt. It combines deterministic planning, strict policy enforcement, and human-in-the-loop confirmations to make system administration safer and more intuitive.

Key Features
------------

- Natural-language to commands with structured plans
- Policy-based allow/deny validation and shell-free execution
- Dry-run, interactive approval, and full audit-friendly output
- OpenWrt focus: uci, ubus, fw4, opkg, diagnostics
- Provider-agnostic design, Gemini HTTP integration included

Quick Start
-----------

1. Build:

```bash
cd g
go build ./cmd/g
```

2. Configure API key using one of the following (precedence: env > UCI > file):

- Environment: `export GEMINI_API_KEY=...`
- OpenWrt UCI: `uci set g.@api[0]=api; uci set g.@api[0].key=...; uci commit g`
- JSON file: `/etc/g/config.json` or `$HOME/.config/g/config.json`

3. Run a dry-run request:

```bash
./g -dry-run "open port 22 on firewall for lan"
```

4. Approve and execute:

```bash
./g -dry-run=false -approve "restart wifi"
```

Safety Model
------------

- No shell expansion or pipelines; argv-only execution
- Policy engine with allowlist and denylist regexes
- Non-root by default; explicit elevation only when required
- Per-command timeouts and minimal environment
- Human confirmation unless `-approve` is set

Configuration
-------------

Config precedence is file < UCI < env. The config schema:

```json
{
  "author": "aezizhu",
  "api_key": "...",
  "endpoint": "https://generativelanguage.googleapis.com/v1beta",
  "model": "gemini-1.5-flash",
  "dry_run": true,
  "auto_approve": false,
  "timeout_seconds": 30,
  "max_commands": 10,
  "allowlist": ["^uci(\\s|$)", "^ubus(\\s|$)"],
  "denylist": ["^rm -rf /"],
  "log_file": "/tmp/g.log"
}
```

CLI Flags
---------

- `-config`: path to JSON config file
- `-model`: model name (default: gemini-1.5-flash)
- `-dry-run`: only print plan (default: true)
- `-approve`: auto-approve plan
- `-timeout`: per-command timeout (default: 30s)
- `-max-commands`: max commands to run (default: 10)
- `-log-file`: execution log path (informational)
- `-version`: print version

Elevation
---------

- Some operations require root. If a plan marks a command with `needs_root: true`, the executor will prefix the argv with `elevate_command` when configured (e.g., `doas -n` or `sudo -n`).
- `-version`: print version

Development
-----------

- Go 1.21+
- Code layout:

```
cmd/g           # CLI entrypoint
internal/config # config loader (env, UCI, file)
internal/llm    # provider clients (Gemini)
internal/plan   # planner schema and instructions
internal/policy # allow/deny validation
internal/executor # argv-only runner with timeouts
internal/ui     # CLI I/O helpers
```

OpenWrt Notes
-------------

- Cross-compile with appropriate target (see `docs/OPENWRT.md`).
- UCI storage for API key is supported via `g.@api[0].key`.
- Ensure required tools exist in `PATH`: `uci`, `ubus`, `fw4`, `opkg`.

Security Considerations
-----------------------

- Keep allowlist narrowly scoped and review regularly.
- Avoid blanket `-approve` in unattended environments.
- Route logs to persistent storage if needed.

License
-------
SEO Topics and Keywords
-----------------------

- OpenWrt natural language CLI, router automation, LLM-assisted administration, firewall management, UCI automation, ubus integration, fw4 control, opkg package management, secure command execution, policy engine, dry-run approval workflow, headless Linux orchestration, embedded systems operations, network diagnostics automation, infrastructure as conversation.

About This Project
------------------

`g` aims to make router administration safer and faster by combining deterministic execution with human-readable intent. It focuses on OpenWrt first, with a provider-agnostic design and strong safety defaults.


MIT


