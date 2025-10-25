LuCI Integration
================

Author: AZ <Aezi.zhu@icloud.com>

Overview
--------

`luci-app-lucicodex` provides a simple web UI to submit a request to `lucicodex` and show the command output. It is intended for trusted administrators; it calls `lucicodex` with `-approve`.

Paths
-----

- Controller: `luci.controller.lucicodex`
- View: `lucicodex/overview`
- Menu: System → LuCICodex

API
---

- POST `admin/system/lucicodex/execute` with JSON body `{ "prompt": "..." }`
- Response JSON: `{ ok: boolean, output: string }`

Security Notes
--------------

- The backend executes `/usr/bin/lucicodex -dry-run=false -approve <q>`.
- Ensure allowlist/denylist in `lucicodex` config are strict.
- Consider restricting access to LuCI or this endpoint to admin users only.


