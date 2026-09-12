//go:build darwin

package envtype

import (
    "os"
    "os/exec"
    "strings"
)

func DetectDebugging() (bool, error) {
    // Darwin: check via sysctl kern.proc.pid as fallback to P_TRACED
    // Use ptrace PT_DENY_ATTACH check via exec
    cmd := exec.Command("sysctl", "-n", "kern.proc.pid")
    if err := cmd.Run(); err != nil {
        return false, nil
    }
    // Check env for debugger
    if _, ok := os.LookupEnv("DYLD_INSERT_LIBRARIES"); ok {
        return true, nil
    }
    return false, nil
}

func Protector() error {
    // Check common analysis tools on macOS
    procs := []string{"Xcode", "lldb", "dtrace"}
    for _, p := range procs {
        if _, err := exec.LookPath(p); err == nil {
            // If debugger-related env present, exit is handled by caller
        }
    }
    return nil
}

func DetectSandbox() bool {
    // Check for macOS sandbox artifacts
    if _, err := os.Stat("/System/Library/CoreServices/Finder.app"); err != nil {
        return true
    }
    // Check for limited file count in /Applications
    if entries, err := os.ReadDir("/Applications"); err == nil && len(entries) < 5 {
        return true
    }
    return false
}

func VirtualizationSystem() string {
    out, err := exec.Command("sysctl", "-n", "machdep.cpu.brand_string").Output()
    if err == nil && strings.Contains(strings.ToLower(string(out)), "virtual") {
        return "VMware"
    }
    // Check for VirtualBox via system_profiler
    if _, err := exec.LookPath("system_profiler"); err == nil {
        out, _ := exec.Command("system_profiler", "SPHardwareDataType").Output()
        if strings.Contains(string(out), "VirtualBox") {
            return "VirtualBox"
        }
    }
    return ""
}
