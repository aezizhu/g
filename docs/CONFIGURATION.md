Configuration
=============

Author: aezizhu

Sources and Precedence
----------------------

1. Defaults (compiled)
2. JSON file: `/etc/g/config.json` or `$HOME/.config/g/config.json`
3. OpenWrt UCI: `g.@api[0].*`
4. Environment variables

Environment Variables
---------------------

- `GEMINI_API_KEY`: API key (required unless set in UCI or file)
- `GEMINI_ENDPOINT`: Override the base API endpoint
- `G_MODEL`: Override the model name
- `G_LOG_FILE`: Override log path
- `G_ELEVATE`: Elevation command prefix (e.g., `doas -n`) when `needs_root` is set
- `G_PROVIDER`: Provider name (default `gemini`)

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
  "log_file": "/tmp/g.log"
  ,"elevate_command": "doas -n"
}
```

OpenWrt UCI
-----------

```bash
uci set g.@api[0]=api
uci set g.@api[0].key=YOUR_KEY
uci set g.@api[0].model=gemini-1.5-flash
uci commit g
```


