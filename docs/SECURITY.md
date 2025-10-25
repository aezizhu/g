Security Model
==============

Author: AZ <Aezi.zhu@icloud.com>

Threat Model
------------

- Accidental harmful commands due to misinterpretation
- Malicious prompt injections
- Privilege escalation or environment abuse

Controls
--------

- Shell-free execution: commands are argv arrays, no pipelines or redirections
- Allowlist and denylist regexes checked against entire command line
- Minimal environment: only `PATH` preserved
- Per-command timeouts; SIGTERM then SIGKILL on deadline
- Interactive confirmation by default
- Non-root by default; explicit elevation is required when needed

Best Practices
--------------

- Keep allowlist minimal and review regularly
- Avoid unattended `-approve` in production
- Send logs to persistent storage if required for audit


