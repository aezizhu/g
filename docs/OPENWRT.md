OpenWrt Notes
=============

Author: aezizhu

Targets
-------

Cross-compile for the router's architecture (e.g., mips, mipsel, arm, aarch64):

```bash
go run ./scripts/build-openwrt.sh
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

Feed Setup (optional)
---------------------

Use this repository as an OpenWrt feed:

```bash
# in your OpenWrt buildroot
echo "src-git g https://github.com/aezizhu/g.git;main" >> feeds.conf
./scripts/feeds update g
./scripts/feeds install -a -p g
make package/g/compile V=s
make package/luci-app-g/compile V=s
```

Prebuilt Releases
-----------------

Download `.ipk` from the Releases page and install:

```bash
opkg install g_*.ipk
opkg install luci-app-g_*.ipk
```


