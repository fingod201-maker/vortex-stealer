//go:build darwin

package main

func newBrowserDataExtractor() BrowserDataExtractor {
    return &DarwinBrowserDataExtractor{}
}
