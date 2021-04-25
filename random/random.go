package random

import (
	"fmt"
        "crypto/rand"
)

const (
    letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890_@#!&|()/<>[]+-%*~"
    alphanumericLetters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
)

// GetRandomString is get random string
func GetRandomString(n int) (string, error) {
	if n == 0 {
		return "", fmt.Errorf("invalid argument")
	}
        buf := make([]byte, n)
        readLen, err := rand.Read(buf)
        if err != nil {
		return "", fmt.Errorf("can not get random data: %w", err)
        }
        letterLen := len(letters)
        for i := 0; i < readLen; i++ {
                buf[i] = letters[int(buf[i]) % letterLen]
        }
        return string(buf), nil
}

// GetAlphanumericRandomString is get alphanumeric random string
func GetAlphanumericRandomString(n int) (string, error) {
	if n == 0 {
		return "", fmt.Errorf("invalid argument")
	}
        buf := make([]byte, n)
        readLen, err := rand.Read(buf)
        if err != nil {
		return "", fmt.Errorf("can not get random data: %w", err)
        }
        letterLen := len(alphanumericLetters)
        for i := 0; i < readLen; i++ {
                buf[i] = alphanumericLetters[int(buf[i]) % letterLen]
        }
        return string(buf), nil
}
