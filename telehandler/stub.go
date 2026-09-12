//go:build !windows

package telehandler

func ClientCommands() {}
func DataUploadState(stateOK bool) error { return nil }
func CapDisplayScreen(hostId string) error { return nil }
func DropAndExec(url string, fname string, path string) error { return nil }
