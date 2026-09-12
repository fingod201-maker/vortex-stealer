//go:build !windows && !darwin && !linux

package envtype

func DetectDebugging() (bool, error) { return false, nil }
func Protector() error { return nil }
func DetectSandbox() bool { return false }
func VirtualizationSystem() string { return "" }
