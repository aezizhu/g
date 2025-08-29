# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [0.1.0] - 2025-01-01

### Added
- Core CLI (`g`) with natural language to command translation
- Multi-provider support: Gemini (API), Gemini CLI (external), OpenAI, Anthropic
- Safety features: dry-run mode, policy engine with allow/deny lists, human confirmation
- OpenWrt integration: UCI config, uci/ubus/fw4/opkg command recognition, environment facts
- LuCI web interface (`luci-app-g`) for browser-based access
- Command execution with timeouts, privilege elevation, audit logging
- JSON output mode and per-step confirmation
- Cross-compilation for OpenWrt targets (mips, mipsle, arm, aarch64, x86_64)
- GitHub Actions CI/CD with release automation
- OpenWrt package metadata and feed configuration
- Comprehensive documentation with security model and usage guides

### Security
- Shell-free execution (argv arrays only)
- Policy-based command validation
- Minimal environment and process timeouts
- Non-root execution by default with explicit elevation
- Audit trail with structured logging

[0.1.0]: https://github.com/aezizhu/g/releases/tag/v0.1.0
