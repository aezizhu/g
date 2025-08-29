Usage Guide
===========

Author: aezizhu

Basics
------

```bash
./g "restart wifi"
```

This performs a dry run by default, showing a plan and commands.

Approving Execution
-------------------

```bash
./g -dry-run=false -approve "open port 22 for lan"
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
./g -interactive
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
./g -setup
```

The wizard will guide you through provider selection, credential configuration, and security settings.

Troubleshooting
---------------

- Ensure `GEMINI_API_KEY` is set or UCI config exists.
- Confirm that required OpenWrt tools are installed and in `PATH`.


