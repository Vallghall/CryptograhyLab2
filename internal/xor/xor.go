package xor

import (
	"fmt"
	"strconv"
)

func CipherString(text, key string) (string, error) {
	input, err := strconv.ParseInt(text, 2, 32)
	if err != nil {
		return "", fmt.Errorf("invalid input, original text must be of int type")
	}
	binKey, err := strconv.ParseInt(key, 2, 32)
	if err != nil {
		return "", fmt.Errorf("invalid input, key must be of int type")
	}

	return fmt.Sprintf("%08b", input^binKey), nil
}

func CipherInt(text, key []int) []int {
	for i, sym := range text {
		text[i] = sym ^ key[i]
	}
	return text
}
