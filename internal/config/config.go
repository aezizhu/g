package config

import (
    "encoding/json"
    "errors"
    "fmt"
    "os"
    "os/exec"
    "path/filepath"
    "strings"
)

type Config struct {
    Author         string   `json:"author"`
    APIKey         string   `json:"api_key"`
    Endpoint       string   `json:"endpoint"`
    Model          string   `json:"model"`
    Provider       string   `json:"provider"`
    DryRun         bool     `json:"dry_run"`
    AutoApprove    bool     `json:"auto_approve"`
    TimeoutSeconds int      `json:"timeout_seconds"`
    MaxCommands    int      `json:"max_commands"`
    Allowlist      []string `json:"allowlist"`
    Denylist       []string `json:"denylist"`
    LogFile        string   `json:"log_file"`
    ElevateCommand string   `json:"elevate_command"`
}

func defaultConfig() Config {
    return Config{
        Author:         "aezizhu",
        Endpoint:       "https://generativelanguage.googleapis.com/v1beta",
        Model:          "gemini-1.5-flash",
        Provider:       "gemini",
        DryRun:         true,
        AutoApprove:    false,
        TimeoutSeconds: 30,
        MaxCommands:    10,
        Allowlist: []string{
            `^uci(\s|$)`,
            `^ubus(\s|$)`,
            `^fw4(\s|$)`,
            `^opkg(\s|$)(update|install|remove|list|info)`,
            `^logread(\s|$)`,
            `^dmesg(\s|$)`,
            `^ip(\s|$)`,
            `^ifstatus(\s|$)`,
            `^cat(\s|$)`,
            `^tail(\s|$)`,
            `^grep(\s|$)`,
            `^awk(\s|$)`,
            `^sed(\s|$)`,
        },
        Denylist: []string{
            `^rm\s+-rf\s+/`,
            `^mkfs(\s|$)`,
            `^dd(\s|$)`,
            `^:(){:|:&};:`,
        },
        LogFile: "/tmp/g.log",
        ElevateCommand: "",
    }
}

// Load loads configuration from env, UCI (if available), and optional JSON file.
// Precedence: env > UCI > file > defaults
func Load(path string) (Config, error) {
    cfg := defaultConfig()

    // File
    if path == "" {
        if fileExists("/etc/g/config.json") {
            path = "/etc/g/config.json"
        } else {
            home, _ := os.UserHomeDir()
            p := filepath.Join(home, ".config", "g", "config.json")
            if fileExists(p) {
                path = p
            }
        }
    }
    if path != "" && fileExists(path) {
        b, err := os.ReadFile(path)
        if err != nil {
            return cfg, fmt.Errorf("read config: %w", err)
        }
        if err := json.Unmarshal(b, &cfg); err != nil {
            return cfg, fmt.Errorf("parse config: %w", err)
        }
    }

    // UCI (OpenWrt)
    if key, _ := uciGet("g.@api[0].key"); key != "" {
        cfg.APIKey = key
    }
    if m, _ := uciGet("g.@api[0].model"); m != "" {
        cfg.Model = m
    }
    if ep, _ := uciGet("g.@api[0].endpoint"); ep != "" {
        cfg.Endpoint = ep
    }

    // Env
    if v := strings.TrimSpace(os.Getenv("GEMINI_API_KEY")); v != "" {
        cfg.APIKey = v
    }
    if v := strings.TrimSpace(os.Getenv("GEMINI_ENDPOINT")); v != "" {
        cfg.Endpoint = v
    }
    if v := strings.TrimSpace(os.Getenv("G_MODEL")); v != "" { // optional alias
        cfg.Model = v
    }
    if v := strings.TrimSpace(os.Getenv("G_LOG_FILE")); v != "" {
        cfg.LogFile = v
    }
    if v := strings.TrimSpace(os.Getenv("G_ELEVATE")); v != "" {
        cfg.ElevateCommand = v
    }
    if v := strings.TrimSpace(os.Getenv("G_PROVIDER")); v != "" {
        cfg.Provider = v
    }

    if cfg.APIKey == "" {
        return cfg, errors.New("API key not configured")
    }
    return cfg, nil
}

func fileExists(p string) bool {
    st, err := os.Stat(p)
    return err == nil && !st.IsDir()
}

func uciGet(key string) (string, error) {
    _, err := exec.LookPath("uci")
    if err != nil {
        return "", err
    }
    out, err := exec.Command("uci", "-q", "get", key).Output()
    if err != nil {
        return "", err
    }
    return strings.TrimSpace(string(out)), nil
}


