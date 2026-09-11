package main

type BrowserDataExtractor interface {
    ExtractCookies(path string) ([]byte, error)
    ExtractLoginData(path string) ([]byte, error)
    ExtractCreditCards(path string) ([]byte, error)
    ExtractHistory(path string) ([]byte, error)
    ExtractExtensions(path string) ([]byte, error)
}
