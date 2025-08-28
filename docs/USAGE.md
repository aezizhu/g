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

Troubleshooting
---------------

- Ensure `GEMINI_API_KEY` is set or UCI config exists.
- Confirm that required OpenWrt tools are installed and in `PATH`.


