//go:build !windows && !darwin && !linux

package install

func CheckDeviceFingerprint() bool { return false }
func FingerprintDevice() error { return nil }
func InstallPersistence() error { return nil }
func UninstallPersistence() {}
func DataDumpLocation() string { return "/tmp" }
func ProcessLock() bool { return false }
func ProcessUnlock() {}
