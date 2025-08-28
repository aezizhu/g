LuCI Integration
================

Author: aezizhu

Overview
--------

`luci-app-g` provides a simple web UI to submit a request to `g` and show the command output. It is intended for trusted administrators; it calls `g` with `-approve`.

Paths
-----

- Controller: `luci.controller.g`
- View: `g/overview`
- Menu: System → g Assistant

API
---

- POST `admin/system/g/run` with form field `q` (the natural language request)
- Response JSON: `{ ok: boolean, output: string }`

Security Notes
--------------

- The backend executes `/usr/bin/g -dry-run=false -approve <q>`.
- Ensure allowlist/denylist in `g` config are strict.
- Consider restricting access to LuCI or this endpoint to admin users only.


