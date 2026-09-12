//go:build windows

package main
func newBrowserDataExtractor() BrowserDataExtractor {
    return &WindowsBrowserDataExtractor{}
}
