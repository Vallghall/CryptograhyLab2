package xor

import (
	"fmt"
	"strconv"
)

func Cipher(text, key string) (string, error) {
	input, err := strconv.Atoi(text)
	if err != nil {
		return "", fmt.Errorf("invalid input, original text must be of int type")
	}
	intKey, err := strconv.Atoi(key)
	if err != nil {
		return "", fmt.Errorf("invalid input, key must be of int type")
	}

	return fmt.Sprintf("%08b", input^intKey), nil
}
