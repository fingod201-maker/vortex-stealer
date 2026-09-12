//go:build !windows && !darwin && !linux && !js

package main

func newBrowserDataExtractor() BrowserDataExtractor {
    return nil
}
