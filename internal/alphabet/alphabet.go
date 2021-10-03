package alphabet

import (
	"secondlab/internal/xor"
	"unicode"
)

func CipherAlphabetically(input, key string) string {
	text := toAlphabetIndex([]rune(input))
	indexedKey := toAlphabetIndex([]rune(key))
	text = modSequence(text, indexedKey)
	return string(toAlphabetCode(text))
}

func CipherBinary(input, key string) string {
	text := toAlphabetIndex([]rune(input))
	indexedKey := toAlphabetIndex([]rune(key))
	text = xor.CipherInt(text, indexedKey)
	return string(toAlphabetCode(text))
}

func IsRussian(input string) bool {
	for _, symbol := range input {
		if !unicode.Is(unicode.Cyrillic, symbol) {
			return false
		}
	}
	return true
}

func toAlphabetIndex(text []rune) []int {
	dif := 'а' - '0'
	output := make([]int, len(text))
	for i, sym := range text {
		output[i] = int(sym - dif - '0')
	}
	return output
}

func toAlphabetCode(text []int) []rune {
	output := make([]rune, len(text))
	for i := range text {
		output[i] = 'а' + rune(text[i])
	}
	return output
}

func modSequence(text, key []int) []int {
	for i := range text {
		text[i] = mod(text[i], key[i])
	}
	return text
}

func mod(txtSym, keySym int) int {
	return (txtSym + keySym) % 32
}
