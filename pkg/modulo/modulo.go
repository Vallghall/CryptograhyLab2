package modulo

import (
	"fmt"
	"secondlab/internal/alphabet"
	"unicode/utf8"
)

func CipherMod32(input, key string) (string, error) {
	if !alphabet.IsRussian(input) {
		return "", fmt.Errorf("your input text must consist of russian characters only")
	}
	if utf8.RuneCountInString(key) < utf8.RuneCountInString(input) {
		return "", fmt.Errorf("your input key must NOT be shorter than the input text")
	}
	return alphabet.CipherAlphabetically(input, key), nil
}

func CipherXOR(input, key string) (string, error) {
	if !alphabet.IsRussian(input) {
		return "", fmt.Errorf("your input text must consist of russian characters only")
	}
	if utf8.RuneCountInString(key) < utf8.RuneCountInString(input) {
		return "", fmt.Errorf("your input key must NOT be shorter than the input text")
	}
	return alphabet.CipherBinary(input, key), nil
}
