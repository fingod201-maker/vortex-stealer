//go:build !windows

package messenger

func DiscordDataDump(mainFolder string) error { return nil }
func GetTelegramProcessPath(mainFolder string) error { return nil }
func GetTelegramDefaultPath(mainFolder string) error { return nil }
