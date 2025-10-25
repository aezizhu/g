Policy Guide
============

Author: AZ <Aezi.zhu@icloud.com>

Overview
--------

The policy engine permits or rejects proposed commands before execution. It relies on allowlist and denylist regular expressions matched against the full command line string.

Recommended Defaults
--------------------

- Allow: `uci`, `ubus`, `fw4`, `opkg` (read and controlled writes), diagnostics like `logread`, `dmesg`, `ip`.
- Deny: destructive tools (`rm -rf /`, `mkfs`, `dd`) and fork bombs.

Examples
--------

```json
{
  "allowlist": ["^uci(\\s|$)", "^ubus(\\s|$)", "^fw4(\\s|$)", "^opkg(\\s|$)(update|install|remove|list|info)"],
  "denylist": ["^rm\\s+-rf\\s+/", "^mkfs(\\s|$)", "^dd(\\s|$)"]
}
```

Extending Policy
----------------

- Edit the JSON config or UCI entries.
- Keep patterns precise. Favor exact subcommands for tools that can be destructive.

Testing Policy
--------------

```bash
go test ./internal/policy
```


