module("luci.controller.g", package.seeall)

function index()
    entry({"admin", "system", "g"}, firstchild(), _("g Assistant"), 60).dependent = false
    entry({"admin", "system", "g", "overview"}, template("g/overview"), _("Overview"), 1)
    entry({"admin", "system", "g", "run"}, call("action_run")).leaf = true
end

function action_run()
    local http = require "luci.http"
    local util = require "luci.util"
    local json = require "luci.jsonc"

    local body = http.formvalue("q") or ""
    if body == "" then
        http.status(400, "Bad Request")
        http.write_json({ error = "missing q" })
        return
    }

    local cmd = {"/usr/bin/g", "-dry-run=false", "-approve", body}
    local output = util.exec(table.concat(cmd, " ")) or ""
    http.prepare_content("application/json")
    http.write_json({ ok = true, output = output })
end


