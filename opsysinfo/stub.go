//go:build !windows

package opsysinfo

type DeviceInfo struct {
    ScreenResolution     string
    Username             string
    CpuName              string
    Hostname             string `json:"hostname"`
    RAM                  uint64 `json:"ram"`
    VirtualizationSystem string `json:"virtualizationSystem"`
    Procs                int32  `json:"procs"`
    CPUArch              string `json:"cpuArch"`
    HostID               string `json:"hostid"`
}

type WindowsInfo struct {
    RegisteredOwner string
    ProductName     string
    DisplayVersion  string
}

type DiskStatus struct {
    DiskPartitions []string
    DiskTotalSpace uint64 `json:"total"`
    DiskFreeSpace  uint64 `json:"free"`
    DiskUsedSpace  uint64 `json:"used"`
}

type NicInfo struct {
    MacAddress   string
    LocalAddress string
}

func HardDriveInfo() DiskStatus { return DiskStatus{} }
func OsVersion() WindowsInfo { return WindowsInfo{} }
func MachineSpecs() DeviceInfo { return DeviceInfo{HostID: "unknown", Username: "unknown", Hostname: "unknown"} }
func LocalNetworkInfo() NicInfo { return NicInfo{} }
func ScreenResolution() (DeviceInfo, int, int) { return DeviceInfo{}, 0, 0 }
