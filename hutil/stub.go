//go:build !windows

package hutil

func CompressDataFiles(mainFolder string, finalDataDump string, pass string) error { return nil }
func GetProcessIdByName(procName string) (int, error) { return 0, nil }
func GetProcessPath(pName string) (string, error) { return "", nil }
func TerminateProcess(procName string) error { return nil }
