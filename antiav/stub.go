//go:build !windows && !darwin && !linux

package antiav

func DisableUAC() {}
func DisableWDInitiate() {}
func AvProcs() {}
