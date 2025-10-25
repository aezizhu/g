OpenWrt Notes
=============

Author: AZ <Aezi.zhu@icloud.com>

Targets
-------

Cross-compile for the router's architecture (e.g., mips, mipsel, arm, aarch64):

```bash
go run ./scripts/build-openwrt.sh
```

UCI Configuration
-----------------

```bash
uci set lucicodex.@api[0]=api
uci set lucicodex.@api[0].key=YOUR_KEY
uci set lucicodex.@api[0].model=gemini-1.5-flash
uci commit lucicodex
```

Dependencies
------------

Ensure the following tools are installed and in `PATH`:

- `uci`, `ubus`, `fw4`, `opkg`
- Diagnostics: `logread`, `dmesg`, `ip`, `ifstatus`

Deployment
----------

Copy the `lucicodex` binary to `/usr/bin/lucicodex` and ensure it is executable. A compatibility symlink to `/usr/bin/g` may also be created.

Feed Setup (optional)
---------------------

Use this repository as an OpenWrt feed:

```bash
# in your OpenWrt buildroot
echo "src-git lucicodex https://github.com/aezizhu/LuciCodex.git;main" >> feeds.conf
./scripts/feeds update lucicodex
./scripts/feeds install -a -p lucicodex
make package/lucicodex/compile V=s
make package/luci-app-lucicodex/compile V=s
```

Prebuilt Releases
-----------------

Download `.ipk` from the Releases page and install:

```bash
opkg install lucicodex_*.ipk
opkg install luci-app-lucicodex_*.ipk
```


