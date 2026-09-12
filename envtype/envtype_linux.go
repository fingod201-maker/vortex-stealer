//go:build linux

package envtype

import (
    "os"
    "os/exec"
    "strings"
)

func DetectDebugging() (bool, error) {
    // Linux: check TracerPid in /proc/self/status
    data, err := os.ReadFile("/proc/self/status")
    if err != nil {
        return false, nil
    }
    for _, line := range strings.Split(string(data), "\n") {
        if strings.HasPrefix(line, "TracerPid:") {
            fields := strings.Fields(line)
            if len(fields) == 2 && fields[1] != "0" {
                return true, nil
            }
        }
    }
    return false, nil
}

func Protector() error {
    // Check for common analysis tools via process list
    out, err := exec.Command("ps", "-e").Output()
    if err != nil {
        return nil
    }
    suspicious := []string{"wireshark", "tcpdump", "strace", "ltrace", "gdb"}
    lower := strings.ToLower(string(out))
    for _, s := range suspicious {
        if strings.Contains(lower, s) {
            // Caller will os.Exit, just return
            break
        }
    }
    return nil
}

func DetectSandbox() bool {
    // Check for sandbox artifacts: limited cgroup, /proc/self/cgroup
    if data, err := os.ReadFile("/proc/self/cgroup"); err == nil {
        if strings.Contains(string(data), "docker") || strings.Contains(string(data), "kubepods") {
            return true
        }
    }
    // Check for low RAM / CPU count typical of sandbox
    if entries, err := os.ReadDir("/proc/self"); err == nil && len(entries) < 50 {
        return true
    }
    return false
}

func VirtualizationSystem() string {
    // Check systemd-detect-virt
    if out, err := exec.Command("systemd-detect-virt").Output(); err == nil {
        v := strings.TrimSpace(string(out))
        if v != "none" && v != "" {
            return v
        }
    }
    // Fallback via lscpu
    if out, err := exec.Command("lscpu").Output(); err == nil {
        if strings.Contains(strings.ToLower(string(out)), "hypervisor") {
            return "Hypervisor"
        }
    }
    return ""
}
