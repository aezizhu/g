Providers and Authentication
===========================

Author: AZ <Aezi.zhu@icloud.com>

Overview
--------

`LuciCodex` supports multiple providers for planning: Gemini (API), Gemini CLI (external), OpenAI, and Anthropic.

Selection
---------

- CLI flag: `-provider gemini|gemini-cli|openai|anthropic`
- Env: `LUCICODEX_PROVIDER` (legacy `G_PROVIDER` still read)

Gemini (API)
------------

- Configure `GEMINI_API_KEY` (or UCI/file). Uses HTTPS API.

Gemini CLI (External)
---------------------

- Install `@google/gemini-cli` or another CLI that prints text.
- Configure path via `LUCICODEX_EXTERNAL_GEMINI` (legacy `G_EXTERNAL_GEMINI`) (default `/usr/bin/gemini`).
- `lucicodex` invokes it and attempts to parse a JSON plan from stdout.
- For login, use the CLI’s built-in OAuth or device code flow.

OpenAI
------

- Set `OPENAI_API_KEY`.
- Default model: `gpt-4o-mini` (override with `-model`).

Anthropic
---------

- Set `ANTHROPIC_API_KEY`.
- Default model: `claude-3-5-sonnet-20240620` (override with `-model`).

Security
--------

- External CLIs run as subprocesses with timeouts; ensure they are trusted.
- All plans are still validated by the policy engine before execution.


