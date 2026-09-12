//go:build linux

package main

func newBrowserDataExtractor() BrowserDataExtractor {
    return &LinuxBrowserDataExtractor{}
}
