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

- **Natural Language Interface**: Translate plain English to safe OpenWrt commands
- **LuCI Web UI**: Full-featured web interface for configuration and execution
- **Multiple LLM Providers**: Support for Gemini, OpenAI, Anthropic, and external CLI
- **Policy-Based Safety**: Allow/deny validation with shell-free execution
- **Dry-Run Mode**: Preview commands before execution with human approval
- **OpenWrt Integration**: Native support for uci, ubus, fw4, opkg, and diagnostics
- **Secure by Default**: No shell execution, execution locking, rate limiting
- **UCI Configuration**: Store API keys and settings in OpenWrt's native config system

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

LuCI Web Interface
------------------

g includes a comprehensive LuCI web interface for OpenWrt routers, allowing you to configure API keys and interact with your router using natural language directly from the web UI.

### Features

- **Configuration Page**: Easy setup for all LLM providers (Gemini, OpenAI, Anthropic)
- **Interactive Interface**: Enter natural language requests and see AI-generated commands
- **Plan Review**: Review generated commands before execution with safety warnings
- **Dry-Run Mode**: Test commands without executing them
- **Execution Controls**: Execute approved commands directly from the web UI
- **Safety Features**: Built-in execution locking, rate limiting, and input validation

### Installation

1. Install the g package on your OpenWrt router
2. Install luci-app-g package
3. Access the web interface at System → g Assistant

### Configuration via LuCI

Navigate to System → g Assistant → Configuration to set up:

- **LLM Provider**: Choose between Gemini, OpenAI, Anthropic, or External Gemini CLI
- **API Keys**: Enter your provider-specific API keys (stored securely in UCI)
- **Model**: Specify the model to use (or leave empty for provider default)
- **Safety Settings**: Configure dry-run defaults, command timeouts, and limits

### Using the Web Interface

1. Go to System → g Assistant → Run
2. Enter your natural language request (e.g., "Show me the current network configuration")
3. Click "Generate Plan" to see the AI-generated commands
4. Review the plan, including any warnings
5. If satisfied, click "Execute Commands" (or use dry-run mode to just see the commands)

### UCI Configuration

The LuCI interface stores configuration in `/etc/config/g`:

```
config api
	option provider 'gemini'
	option key 'your-gemini-api-key'
	option model 'gemini-1.5-flash'
	option endpoint 'https://generativelanguage.googleapis.com/v1beta'
	option openai_key ''
	option anthropic_key ''

config settings
	option dry_run '1'
	option confirm_each '0'
	option timeout '30'
	option max_commands '10'
	option log_file '/tmp/g.log'
```

You can also configure via UCI commands:

```bash
# Set Gemini API key
uci set g.@api[0].key='your-api-key-here'
uci set g.@api[0].provider='gemini'
uci commit g

# Set OpenAI API key
uci set g.@api[0].provider='openai'
uci set g.@api[0].openai_key='your-openai-key'
uci commit g

# Configure safety settings
uci set g.@settings[0].dry_run='1'
uci set g.@settings[0].timeout='30'
uci commit g
```

### Security Features

- **Execution Locking**: Prevents concurrent command execution
- **Input Validation**: Limits prompt size and validates input
- **Rate Limiting**: Prevents abuse via execution locks
- **Secure Storage**: API keys stored in UCI with restricted permissions (600)
- **RPCD ACL**: Minimal permissions granted to the web interface
- **POST-only Endpoints**: All command execution uses POST with JSON bodies
- **No Shell Execution**: Uses nixio.spawn with argv arrays (no shell injection)

OpenWrt Notes
-------------

- Cross-compile with appropriate target (see `docs/OPENWRT.md`).
- UCI storage for API key is supported via `g.@api[0].key`.
- Ensure required tools exist in `PATH`: `uci`, `ubus`, `fw4`, `opkg`.
- The LuCI web interface provides a user-friendly way to configure and use g.
- Binary size is approximately 8-9MB for MIPS/ARM targets.

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


