# g - Natural Language Assistant for OpenWrt

**Control your OpenWrt router with plain English commands**

Author: aezizhu

<p align="center">
  <a href="#"><img alt="Build" src="https://img.shields.io/badge/build-passing-brightgreen"></a>
  <a href="#license"><img alt="License" src="https://img.shields.io/badge/license-Dual-blue"></a>
  <a href="#"><img alt="Go Version" src="https://img.shields.io/badge/Go-1.21+-1f425f"></a>
  <a href="#"><img alt="OpenWrt" src="https://img.shields.io/badge/OpenWrt-supported-00a0e9"></a>
  <a href="https://github.com/aezizhu/g/actions/workflows/build.yml"><img alt="CI" src="https://github.com/aezizhu/g/actions/workflows/build.yml/badge.svg"></a>
</p>

---

## What is g?

**g** is an intelligent assistant that lets you manage your OpenWrt router using natural language instead of memorizing complex commands. Simply tell g what you want to do in plain English, and it will translate your request into safe, audited commands that you can review before execution.

**Example:** Instead of remembering `uci set wireless.radio0.disabled=0 && uci commit wireless && wifi reload`, just say: *"turn on the wifi"*

---

## Table of Contents

- [Why Use g?](#why-use-g)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation on OpenWrt](#installation-on-openwrt)
  - [Getting Your API Key](#getting-your-api-key)
- [Using g on Your Router](#using-g-on-your-router)
  - [Method 1: Web Interface (Recommended)](#method-1-web-interface-recommended)
  - [Method 2: Command Line (SSH)](#method-2-command-line-ssh)
- [Configuration Guide](#configuration-guide)
  - [Choosing Your AI Provider](#choosing-your-ai-provider)
  - [Configuring via Web Interface](#configuring-via-web-interface)
  - [Configuring via Command Line](#configuring-via-command-line)
- [Common Use Cases](#common-use-cases)
- [Safety Features](#safety-features)
- [Troubleshooting](#troubleshooting)
- [Advanced Usage](#advanced-usage)
- [License](#license)
- [Support](#support)

---

## Why Use g?

### For Home Users
- **No command memorization**: Manage your router in plain English
- **Safe by default**: All commands are reviewed before execution
- **Easy web interface**: No need to SSH into your router
- **Learn as you go**: See the actual commands g generates

### For Power Users
- **Faster administration**: Natural language is quicker than looking up syntax
- **Policy-based safety**: Customize what commands are allowed
- **Multiple AI providers**: Choose between Gemini, OpenAI, or Anthropic
- **Audit trail**: Full logging of all operations

---

## Getting Started

### Prerequisites

Before installing g, you need:

1. **An OpenWrt router** (version 21.02 or later recommended)
2. **Internet connection** on your router
3. **At least 10MB free storage** space
4. **An API key** from one of these providers:
   - Google Gemini (recommended for beginners - free tier available)
   - OpenAI (GPT-4/GPT-3.5)
   - Anthropic (Claude)

### Installation on OpenWrt

#### Step 1: Download the Package

SSH into your router and download the g package for your architecture:

```bash
# For MIPS routers (most common)
cd /tmp
wget https://github.com/aezizhu/g/releases/latest/download/g-mips.ipk

# For ARM routers
wget https://github.com/aezizhu/g/releases/latest/download/g-arm.ipk

# For x86_64 routers
wget https://github.com/aezizhu/g/releases/latest/download/g-amd64.ipk
```

#### Step 2: Install the Package

```bash
opkg update
opkg install /tmp/g-*.ipk
```

#### Step 3: Install the Web Interface (Optional but Recommended)

```bash
opkg install luci-app-g
```

#### Step 4: Verify Installation

```bash
g -version
```

You should see: `g version 0.2.1`

### Getting Your API Key

#### Option 1: Google Gemini (Recommended for Beginners)

1. Visit https://makersuite.google.com/app/apikey
2. Click "Create API Key"
3. Copy your API key (starts with `AIza...`)
4. **Free tier**: 60 requests per minute

#### Option 2: OpenAI

1. Visit https://platform.openai.com/api-keys
2. Click "Create new secret key"
3. Copy your API key (starts with `sk-...`)
4. **Note**: Requires payment method on file

#### Option 3: Anthropic

1. Visit https://console.anthropic.com/settings/keys
2. Click "Create Key"
3. Copy your API key (starts with `sk-ant-...`)
4. **Note**: Requires payment method on file

---

## Using g on Your Router

### Method 1: Web Interface (Recommended)

This is the easiest way to use g, especially if you're not comfortable with command line.

#### Step 1: Access the Web Interface

1. Open your router's web interface (usually http://192.168.1.1)
2. Log in with your admin credentials
3. Navigate to **System → g Assistant**

#### Step 2: Configure Your API Key

1. Click on the **Configuration** tab
2. Select your AI provider (Gemini, OpenAI, or Anthropic)
3. Enter your API key in the appropriate field
4. Click **Save & Apply**

#### Step 3: Use the Assistant

1. Click on the **Run** tab
2. Type your request in plain English, for example:
   - "Show me the current network configuration"
   - "Restart the wifi"
   - "Open port 8080 for my web server"
3. Click **Generate Plan**
4. Review the commands that g suggests
5. If they look correct, click **Execute Commands**

**That's it!** You're now using natural language to control your router.

### Method 2: Command Line (SSH)

If you prefer using SSH or want to automate tasks, you can use g from the command line.

#### Step 1: Configure Your API Key

```bash
# Set your Gemini API key
uci set g.@api[0].provider='gemini'
uci set g.@api[0].key='YOUR-API-KEY-HERE'
uci commit g
```

#### Step 2: Test g with a Dry Run

```bash
g "show me the wifi status"
```

This will show you what commands g would run, but won't execute them yet.

#### Step 3: Execute Commands

If the commands look correct, run with approval:

```bash
g -approve "restart the wifi"
```

Or use interactive mode to confirm each command:

```bash
g -confirm-each "update the package list and install htop"
```

---

## Configuration Guide

### Choosing Your AI Provider

g supports multiple AI providers. Here's how to choose:

| Provider | Best For | Cost | Speed |
|----------|----------|------|-------|
| **Gemini** | Beginners, home users | Free tier available | Fast |
| **OpenAI** | Advanced users, complex tasks | Pay per use | Very fast |
| **Anthropic** | Privacy-conscious users | Pay per use | Fast |

### Configuring via Web Interface

1. Go to **System → g Assistant → Configuration**
2. Configure these settings:

**API Settings:**
- **Provider**: Choose your AI provider
- **API Key**: Enter your key (stored securely)
- **Model**: Leave empty for default, or specify (e.g., `gemini-1.5-flash`, `gpt-4`, `claude-3-sonnet`)
- **Endpoint**: Leave default unless using custom endpoint

**Safety Settings:**
- **Dry Run by Default**: Keep enabled (recommended) - shows commands before running
- **Confirm Each Command**: Enable for extra safety
- **Command Timeout**: How long to wait for each command (default: 30 seconds)
- **Max Commands**: Maximum commands per request (default: 10)
- **Log File**: Where to save execution logs (default: `/tmp/g.log`)

3. Click **Save & Apply**

### Configuring via Command Line

All settings are stored in `/etc/config/g` using OpenWrt's UCI system:

```bash
# Configure Gemini
uci set g.@api[0].provider='gemini'
uci set g.@api[0].key='YOUR-GEMINI-KEY'
uci set g.@api[0].model='gemini-1.5-flash'

# Configure OpenAI
uci set g.@api[0].provider='openai'
uci set g.@api[0].openai_key='YOUR-OPENAI-KEY'
uci set g.@api[0].model='gpt-4'

# Configure Anthropic
uci set g.@api[0].provider='anthropic'
uci set g.@api[0].anthropic_key='YOUR-ANTHROPIC-KEY'
uci set g.@api[0].model='claude-3-sonnet-20240229'

# Safety settings
uci set g.@settings[0].dry_run='1'          # 1=enabled, 0=disabled
uci set g.@settings[0].confirm_each='0'     # 1=confirm each, 0=confirm once
uci set g.@settings[0].timeout='30'         # seconds
uci set g.@settings[0].max_commands='10'    # max commands per request

# Apply changes
uci commit g
```

---

## Common Use Cases

### Network Management

```bash
# Check network status
g "show me all network interfaces and their status"

# Restart network
g -approve "restart the network"

# Configure static IP
g "set lan interface to static ip 192.168.1.1"
```

### WiFi Management

```bash
# Check WiFi status
g "show me the wifi status"

# Change WiFi password
g "change the wifi password to MyNewPassword123"

# Enable/disable WiFi
g -approve "turn off the wifi"
g -approve "turn on the wifi"

# Restart WiFi
g -approve "restart wifi"
```

### Firewall Management

```bash
# Check firewall rules
g "show me the current firewall rules"

# Open a port
g "open port 8080 for tcp traffic from lan"

# Block an IP
g "block ip address 192.168.1.100"
```

### Package Management

```bash
# Update package list
g "update the package list"

# Install a package
g "install the htop package"

# List installed packages
g "show me all installed packages"
```

### System Monitoring

```bash
# Check system status
g "show me system information and uptime"

# Check memory usage
g "show me memory usage"

# Check disk space
g "show me disk space usage"

# View system logs
g "show me the last 20 lines of system log"
```

### Diagnostics

```bash
# Ping test
g "ping google.com 5 times"

# DNS test
g "check if dns is working"

# Check internet connectivity
g "test internet connection"
```

---

## Safety Features

g is designed with safety as the top priority:

### 1. Dry-Run Mode (Default)
By default, g shows you what it would do without actually doing it. You must explicitly approve execution.

### 2. Command Review
Every command is shown to you before execution. You can see exactly what will run on your system.

### 3. Policy Engine
g has built-in rules about what commands are allowed:

**Allowed by default:**
- `uci` (configuration)
- `ubus` (system bus)
- `fw4` (firewall)
- `opkg` (package manager)
- `ip`, `ifconfig` (network info)
- `cat`, `grep`, `tail` (read files)
- `logread`, `dmesg` (logs)

**Blocked by default:**
- `rm -rf /` (dangerous deletions)
- `mkfs` (filesystem formatting)
- `dd` (disk operations)
- Fork bombs and other malicious patterns

### 4. No Shell Execution
g never uses shell expansion or pipes. Commands are executed directly with exact arguments, preventing injection attacks.

### 5. Execution Locking
Only one g command can run at a time, preventing conflicts and race conditions.

### 6. Timeouts
Every command has a timeout (default 30 seconds) to prevent hanging.

### 7. Audit Logging
All commands and their results are logged to `/tmp/g.log` for review.

---

## Troubleshooting

### "API key not configured"

**Solution:** Make sure you've set your API key:

```bash
# Via UCI
uci set g.@api[0].key='YOUR-KEY-HERE'
uci commit g

# Or via environment variable
export GEMINI_API_KEY='YOUR-KEY-HERE'
```

### "execution in progress"

**Solution:** Another g command is running. Wait for it to finish, or remove the lock:

```bash
rm /var/lock/g.lock
```

### "command not found: g"

**Solution:** Make sure g is installed and in your PATH:

```bash
which g
# Should show: /usr/bin/g

# If not found, reinstall:
opkg update
opkg install g
```

### Commands are not executing

**Solution:** Make sure you're not in dry-run mode:

```bash
# Use -approve flag
g -approve "your command here"

# Or disable dry-run in config
uci set g.@settings[0].dry_run='0'
uci commit g
```

### "prompt too long (max 4096 chars)"

**Solution:** Your request is too long. Break it into smaller requests or be more concise.

### Web interface not showing up

**Solution:** Make sure luci-app-g is installed:

```bash
opkg update
opkg install luci-app-g
/etc/init.d/uhttpd restart
```

Then clear your browser cache and reload.

---

## Advanced Usage

### Interactive Mode (REPL)

Start an interactive session where you can have a conversation with g:

```bash
g -interactive
```

### JSON Output

Get structured output for scripting:

```bash
g -json "show network status" | jq .
```

### Custom Configuration File

Use a custom config file instead of UCI:

```bash
g -config /etc/g/custom-config.json "your command"
```

### Environment Variables

Override settings with environment variables:

```bash
export GEMINI_API_KEY='your-key'
export G_PROVIDER='gemini'
export G_MODEL='gemini-1.5-flash'
g "your command"
```

### Command-Line Flags

```bash
g -help
```

Available flags:
- `-approve`: Auto-approve plan without confirmation
- `-dry-run`: Only show plan, don't execute (default: true)
- `-confirm-each`: Confirm each command individually
- `-json`: Output in JSON format
- `-interactive`: Start interactive REPL mode
- `-timeout=30`: Set command timeout in seconds
- `-max-commands=10`: Set max commands per request
- `-model=name`: Override model name
- `-config=path`: Use custom config file
- `-log-file=path`: Set log file path
- `-version`: Show version

### Customizing the Policy

Edit the allowlist and denylist in `/etc/config/g` or your config file:

```json
{
  "allowlist": [
    "^uci(\\s|$)",
    "^custom-command(\\s|$)"
  ],
  "denylist": [
    "^dangerous-command(\\s|$)"
  ]
}
```

---

## License

**Dual License:**

- **Free for individual/personal use** - Use g on your home router at no cost
- **Commercial use requires a license** - Contact aezi.zhu@icloud.com for commercial licensing

See [LICENSE](LICENSE) file for full details.

---

## Support

### Getting Help

- **Documentation**: You're reading it!
- **Issues**: https://github.com/aezizhu/g/issues
- **Discussions**: https://github.com/aezizhu/g/discussions

### Commercial Support

For commercial licensing, enterprise support, or custom development:
- Email: aezi.zhu@icloud.com
- Include "g Commercial License" in the subject line

### Contributing

Contributions are welcome! Please read our contributing guidelines before submitting pull requests.

---

## About This Project

**g** was created to make OpenWrt router administration accessible to everyone, not just networking experts. By combining the power of modern AI with strict safety controls, g lets you manage your router using natural language while maintaining security and transparency.

The project focuses on OpenWrt first, with a provider-agnostic design and strong safety defaults. Every command is audited, every action is logged, and you're always in control.

---

**Made with ❤️ for the OpenWrt community**


