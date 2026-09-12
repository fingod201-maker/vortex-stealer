//go:build js

package main

func newBrowserDataExtractor() BrowserDataExtractor {
    return &MobileBrowserDataExtractor{}
}
