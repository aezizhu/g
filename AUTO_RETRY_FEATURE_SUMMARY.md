# Automatic Error Recovery Feature - Implementation Summary

## Overview

LuciCodex now includes **automatic error detection and retry** functionality. When commands fail, the AI will automatically analyze the error, generate a fix, and retry execution - making it truly self-healing.

## What Was Implemented

### 1. **Configuration System** ?
Added new configuration options to control retry behavior:

```go
// Config additions
MaxRetries  int  `json:"max_retries"`   // Default: 2
AutoRetry   bool `json:"auto_retry"`    // Default: true
```

**Command-line flags:**
- `-auto-retry` - Enable/disable automatic retry (default: true)
- `-max-retries=N` - Maximum retry attempts (default: 2)

**Environment variables:**
- `LUCICODEX_AUTO_RETRY=1` - Enable auto-retry
- `LUCICODEX_MAX_RETRIES=3` - Set max retries

**Configuration file:**
```json
{
  "max_retries": 2,
  "auto_retry": true
}
```

### 2. **Enhanced LLM Provider Interface** ?
Added `GenerateErrorFix()` method to all LLM providers:

```go
type Provider interface {
    GeneratePlan(ctx context.Context, prompt string) (plan.Plan, error)
    GenerateErrorFix(ctx context.Context, originalCommand string, errorOutput string, attempt int) (plan.Plan, error)
}
```

Implemented for:
- ? Gemini (Google)
- ? OpenAI (GPT-4)
- ? Anthropic (Claude)
- ? External Gemini CLI

### 3. **Intelligent Error Analysis** ?
The AI analyzes errors and provides contextual fixes:

**Error Detection Capabilities:**
- File not found ? Suggests alternative paths or package installation
- Permission denied ? Adjusts commands to use root privileges
- Command syntax errors ? Corrects the syntax
- Missing tools ? Recommends package installation
- OpenWrt-specific issues ? Uses proper OpenWrt commands

### 4. **Enhanced Command Instructions** ?
Updated the AI's system prompt with comprehensive OpenWrt knowledge:

**Added Command Reference:**
- Network: `uci show network`, `ip addr`, `ifconfig`, `ifstatus`
- WiFi: `wifi status`, `wifi up/down`, `/etc/init.d/network restart`
- Firewall: `fw4 print`, `uci show firewall`
- Packages: `opkg update`, `opkg install`, `opkg list-installed`
- System: `ubus call system board`, `free`, `df -h`
- Logs: `logread`, `dmesg`, `tail /var/log/messages`
- DNS: `nslookup`, `cat /etc/resolv.conf`

**Common OpenWrt Paths:**
- `/etc/config/` (UCI configuration)
- `/var/log/` (logs)
- `/sys/class/net/` (network interfaces)
- `/tmp/` (temporary files)

### 5. **Expanded Security Policy** ?
Added more allowed commands to the default allowlist:

```go
// New allowed commands
`^wifi(\s|$)`        // WiFi management
`^ping(\s|$)`        // Network diagnostics
`^nslookup(\s|$)`    // DNS queries
`^ifconfig(\s|$)`    // Network interface config
`^route(\s|$)`       // Routing table
`^iptables(\s|$)`    // Firewall rules
`^/etc/init\.d/`     // Service management
```

### 6. **Automatic Retry Logic** ?
Implemented intelligent retry workflow in `main.go`:

**Retry Process:**
1. Execute initial command
2. Detect failure (non-zero exit code)
3. Capture error output
4. Send to AI for analysis
5. Generate fix plan
6. Validate fix against security policy
7. Execute fix
8. Verify success
9. Repeat up to `max_retries` times

**User-Friendly Output:**
```
??  Command failed: logread
Error: exec: "logread": executable file not found in $PATH
Output: 
?? Attempting automatic fix (attempt 1/2)...

?? Fix plan: Use dmesg as alternative to logread
  ? dmesg | tail -n 20

? Fix successful!
```

### 7. **Updated Documentation** ?
- Enhanced `USAGE.md` with auto-retry examples
- Updated `README.md` with error recovery section
- Added safety feature #8: Automatic Error Recovery
- Included practical examples for all common use cases

## How It Works

### Example 1: Missing Command

```bash
$ lucicodex -dry-run=false -approve "show me the last 20 lines of system log"
```

**Scenario:** `logread` not found

1. **Initial attempt:** `["logread", "-n", "20"]`
2. **Error detected:** "executable file not found in $PATH"
3. **AI analysis:** Identifies missing logread utility
4. **Fix generated:** `["dmesg"]` or `["tail", "-n", "20", "/var/log/messages"]`
5. **Retry successful:** Shows system logs using alternative method

### Example 2: Permission Denied

```bash
$ lucicodex -dry-run=false -approve "restart the network"
```

**Scenario:** Insufficient permissions

1. **Initial attempt:** `["/etc/init.d/network", "restart"]`
2. **Error detected:** "permission denied"
3. **AI analysis:** Identifies permission issue
4. **Fix generated:** Same command with `needs_root: true`
5. **Retry successful:** Command runs with proper elevation

### Example 3: Wrong Path

```bash
$ lucicodex -dry-run=false -approve "show wifi configuration"
```

**Scenario:** Looking in wrong config location

1. **Initial attempt:** `["cat", "/etc/wifi.conf"]`
2. **Error detected:** "No such file or directory"
3. **AI analysis:** Identifies incorrect path
4. **Fix generated:** `["uci", "show", "wireless"]`
5. **Retry successful:** Shows proper OpenWrt wireless config

## Testing

All tests pass:
```
? internal/config
? internal/executor
? internal/llm
? internal/plan
? internal/policy
? Binary compilation successful
? Version check: LuciCodex version 0.3.0
```

## Usage Examples

### Basic Usage (Auto-Retry Enabled by Default)

```bash
# Will automatically fix and retry on errors
lucicodex -dry-run=false -approve "show me system logs"
```

### Disable Auto-Retry

```bash
# Fail immediately without retrying
lucicodex -auto-retry=false -dry-run=false -approve "restart wifi"
```

### Custom Retry Limit

```bash
# Allow up to 5 retry attempts
lucicodex -max-retries=5 -dry-run=false -approve "configure firewall"
```

### Interactive Mode with Retry

```bash
# Retry is enabled in interactive mode too
lucicodex -interactive -dry-run=false
> show system status
> restart network
```

## Configuration Options

### Via Command Line
```bash
lucicodex -auto-retry=true -max-retries=3 "your command"
```

### Via Environment Variables
```bash
export LUCICODEX_AUTO_RETRY=1
export LUCICODEX_MAX_RETRIES=3
lucicodex "your command"
```

### Via Config File
```json
{
  "auto_retry": true,
  "max_retries": 3,
  "provider": "gemini",
  "api_key": "your-key"
}
```

### Via UCI (OpenWrt)
```bash
# Future enhancement - not yet implemented in UCI config
```

## Benefits

1. **User-Friendly**: No need to know exact OpenWrt command syntax
2. **Self-Healing**: Automatically recovers from common errors
3. **Educational**: Shows both the error and the fix
4. **Safe**: All fixes are validated by the same security policy
5. **Efficient**: Reduces back-and-forth iterations
6. **Smart**: Learns from OpenWrt-specific patterns

## Common Use Cases Now Working

All the use cases you mentioned are now properly supported with auto-retry:

### Network Management ?
```bash
lucicodex "show me all network interfaces and their status"
lucicodex -approve "restart the network"
lucicodex "set lan interface to static ip 192.168.1.1"
```

### WiFi Management ?
```bash
lucicodex "show me the wifi status"
lucicodex "change the wifi password to MyNewPassword123"
lucicodex -approve "turn off the wifi"
lucicodex -approve "turn on the wifi"
lucicodex -approve "restart wifi"
```

### Firewall Management ?
```bash
lucicodex "show me the current firewall rules"
lucicodex "open port 8080 for tcp traffic from lan"
lucicodex "block ip address 192.168.1.100"
```

### Package Management ?
```bash
lucicodex "update the package list"
lucicodex "install the htop package"
lucicodex "show me all installed packages"
```

### System Monitoring ?
```bash
lucicodex "show me system information and uptime"
lucicodex "show me memory usage"
lucicodex "show me disk space usage"
lucicodex "show me the last 20 lines of system log"  # ? Auto-fixes if logread missing
```

### Diagnostics ?
```bash
lucicodex "ping google.com 5 times"
lucicodex "check if dns is working"
lucicodex "test internet connection"
```

## Next Steps

The automatic error recovery system is now fully functional and ready for use. Users can:

1. **Try it out** with any command - errors will be automatically handled
2. **Customize retry behavior** with flags or config
3. **Monitor logs** to see what fixes were applied
4. **Disable if needed** with `-auto-retry=false`

## Files Modified

1. `internal/config/config.go` - Added MaxRetries and AutoRetry config
2. `internal/llm/provider.go` - Added GenerateErrorFix interface
3. `internal/llm/gemini.go` - Implemented error fix generation
4. `internal/llm/openai.go` - Implemented error fix generation
5. `internal/llm/anthropic.go` - Implemented error fix generation
6. `internal/llm/gemini_external.go` - Implemented error fix generation
7. `internal/plan/plan.go` - Enhanced instructions with OpenWrt commands
8. `cmd/lucicodex/main.go` - Added retry loop logic
9. `docs/USAGE.md` - Added auto-retry documentation
10. `README.md` - Updated with error recovery examples

## Conclusion

LuciCodex is now truly **self-healing**. When you type:

```bash
lucicodex "show me the latest 20 lines of the system log"
```

The AI will:
1. Generate the appropriate command
2. Execute it
3. If it fails, detect and analyze the error
4. Generate a fix automatically
5. Retry with the corrected approach
6. Succeed and show you the logs

**No more trial and error. No more googling OpenWrt commands. Just tell LuciCodex what you want, and it figures out how to do it.**
