//go:build !windows

package datatransfer

type DataDumpInfo struct {
    DumpFile    string
    DumpSize    int64
    UploadSite  string
    DownloadURL string
}

func UploadDataDump(mainFolder string, finalDataDump string) (DataDumpInfo, error) {
    return DataDumpInfo{}, nil
}
