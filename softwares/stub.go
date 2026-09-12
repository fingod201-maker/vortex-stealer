//go:build !windows

package softwares

func PuttyDataQuery(mainFolder string) error { return nil }
func WinSCPDataQuery(mainFolder string) error { return nil }
func TeamviewerDumpPass(mainFolder string) error { return nil }
func SteamDataDump(mainFolder string) error { return nil }
