Development Guide
=================

Author: aezizhu

Prerequisites
-------------

- Go 1.21+

Build
-----

```bash
go build ./cmd/lucicodex
```

Run (dry run by default)
------------------------

```bash
export GEMINI_API_KEY=YOUR_KEY
./lucicodex "restart wifi"
```

Executing
---------

```bash
./lucicodex -dry-run=false -approve "open port 22 for lan"
```

Project Layout
--------------

See `README.md` and `docs/ARCHITECTURE.md`.

Testing
-------

```bash
go test ./...
```

Cross-Compilation
-----------------

See `docs/OPENWRT.md` or use scripts under `scripts/`.


