Usage Guide
===========

Author: AZ <Aezi.zhu@icloud.com>

Basics
------

```bash
lucicodex "restart wifi"
```

This performs a dry run by default, showing a plan and commands.

Approving Execution
-------------------

```bash
lucicodex -dry-run=false -approve "open port 22 for lan"
```

Flags
-----

- `-config` path to JSON config
- `-model` model name
- `-provider` provider name (default: gemini)
- `-dry-run` show plan only (default true)
- `-approve` auto-confirm
- `-confirm-each` confirm each step before execution
- `-timeout` per-command timeout
- `-max-commands` limit
- `-max-retries` maximum retry attempts for failed commands (-1 = use config)
- `-auto-retry` automatically retry failed commands with AI-generated fixes (default true)
- `-log-file` log path hint
- `-version` version info
- `-json` emit JSON for plan/results
- `-facts` include environment facts in prompt (default true)
- `-interactive` start interactive REPL mode
- `-setup` run setup wizard

Interactive Mode
----------------

Start an interactive session:

```bash
lucicodex -interactive
```

Commands in interactive mode:
- `help` - show available commands
- `history` - show command history
- `!<number>` - re-run command from history
- `set key=value` - change settings (dry-run, auto-approve, provider, model)
- `status` - show current configuration
- `clear` - clear history
- `exit` or `quit` - exit interactive mode

Setup Wizard
------------

For first-time setup:

```bash
lucicodex -setup
```

The wizard will guide you through provider selection, credential configuration, and security settings.

Automatic Error Recovery
-------------------------

LuciCodex can automatically detect and fix errors:

```bash
# Will automatically retry if command fails
lucicodex -dry-run=false -approve "show me the last 20 lines of system log"
```

When a command fails, LuciCodex will:
1. Detect the error and capture the output
2. Send the error to the AI for analysis
3. Generate a fix plan automatically
4. Execute the fix and verify success
5. Retry up to `max-retries` times (default: 2)

Example error recovery scenarios:
- **File not found**: AI suggests alternative paths or installation
- **Permission denied**: AI adjusts commands to use proper elevation
- **Command syntax error**: AI corrects the syntax
- **Missing tools**: AI suggests package installation

To disable auto-retry:
```bash
lucicodex -auto-retry=false "your command"
```

Troubleshooting
---------------

- Ensure `GEMINI_API_KEY` is set or UCI config exists.
- Confirm that required OpenWrt tools are installed and in `PATH`.


