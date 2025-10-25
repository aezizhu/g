Configuration
=============

Author: AZ <Aezi.zhu@icloud.com>

Sources and Precedence
----------------------

1. Defaults (compiled)
2. JSON file: `/etc/lucicodex/config.json` or `$HOME/.config/lucicodex/config.json`
3. OpenWrt UCI: `lucicodex.@api[0].*`
4. Environment variables

Environment Variables
---------------------

- `GEMINI_API_KEY`: API key (required unless set in UCI or file)
- `GEMINI_ENDPOINT`: Override the base API endpoint
Additionally supported:
- `LUCICODEX_MODEL`: Override the model name
- `LUCICODEX_LOG_FILE`: Override log path
- `LUCICODEX_ELEVATE`: Elevation command prefix (e.g., `doas -n`) when `needs_root` is set
- `LUCICODEX_PROVIDER`: Provider name (default `gemini`)

Sample JSON
-----------

```json
{
  "author": "aezizhu",
  "api_key": "...",
  "endpoint": "https://generativelanguage.googleapis.com/v1beta",
  "model": "gemini-1.5-flash",
  "provider": "gemini",
  "dry_run": true,
  "auto_approve": false,
  "timeout_seconds": 30,
  "max_commands": 10,
  "allowlist": ["^uci(\\s|$)", "^ubus(\\s|$)"],
  "denylist": ["^rm -rf /"],
  "log_file": "/tmp/lucicodex.log"
  ,"elevate_command": "doas -n"
}
```

OpenWrt UCI
-----------

```bash
uci set lucicodex.@api[0]=api
uci set lucicodex.@api[0].key=YOUR_KEY
uci set lucicodex.@api[0].model=gemini-1.5-flash
uci commit lucicodex
```


