module("luci.controller.g", package.seeall)

function index()
    entry({"admin", "system", "g"}, firstchild(), _("g Assistant"), 60).dependent = false
    entry({"admin", "system", "g", "overview"}, template("g/overview"), _("Overview"), 1)
    entry({"admin", "system", "g", "run"}, call("action_run")).leaf = true
    entry({"admin", "system", "g", "metrics"}, call("action_metrics")).leaf = true
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

    local dry_run = http.formvalue("dry_run")
    local cmd = {"/usr/bin/g"}
    
    if dry_run then
        table.insert(cmd, "-dry-run=true")
    else
        table.insert(cmd, "-dry-run=false")
        table.insert(cmd, "-approve")
    end
    
    table.insert(cmd, body)
    
    local output = util.exec(table.concat(cmd, " ")) or ""
    http.prepare_content("application/json")
    http.write_json({ ok = true, output = output })
end

function action_metrics()
    local http = require "luci.http"
    local util = require "luci.util"
    local json = require "luci.jsonc"

    -- Try to get metrics from g CLI
    local output = util.exec("/usr/bin/g -auth status 2>/dev/null") or "{}"
    
    -- Parse or provide default metrics
    local metrics = {
        total_requests = 0,
        success_rate = 0.0,
        average_duration = 0,
        top_provider = "unknown",
        top_command = "unknown"
    }
    
    -- Try to load from file
    local f = io.open("/tmp/g-metrics.json", "r")
    if f then
        local content = f:read("*all")
        f:close()
        local parsed = json.parse(content)
        if parsed then
            metrics = parsed
        end
    end
    
    http.prepare_content("application/json")
    http.write_json(metrics)
end


