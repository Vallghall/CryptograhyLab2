package modulo

import (
	"fmt"
	"secondlab/internal/alphabet"
)

func Cipher(input, key string) (string, error) {
	if !alphabet.IsRussian(input) {
		return "", fmt.Errorf("your input text must consist of russian characters only")
	}
	if len(key) < len(input) {
		return "", fmt.Errorf("your input key must NOT be shorter than the input text")
	}
	return alphabet.CipherAlphabetically(input, key), nil
}
