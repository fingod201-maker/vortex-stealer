//go:build !windows && !darwin && !linux

package wifi

func DumpWifiPasswords(mainFolder string) error { return nil }
