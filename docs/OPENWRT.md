OpenWrt Notes
=============

Author: aezizhu

Targets
-------

Cross-compile for the router's architecture (e.g., mips, mipsel, arm, aarch64):

```bash
GOOS=linux GOARCH=mipsle go build -trimpath -ldflags "-s -w" ./cmd/g
```

UCI Configuration
-----------------

```bash
uci set g.@api[0]=api
uci set g.@api[0].key=YOUR_KEY
uci set g.@api[0].model=gemini-1.5-flash
uci commit g
```

Dependencies
------------

Ensure the following tools are installed and in `PATH`:

- `uci`, `ubus`, `fw4`, `opkg`
- Diagnostics: `logread`, `dmesg`, `ip`, `ifstatus`

Deployment
----------

Copy the `g` binary to `/usr/bin/g` and ensure it is executable.


