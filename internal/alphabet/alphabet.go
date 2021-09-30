package alphabet

import "unicode"

func CipherAlphabetically(input, key string) string {
	text := toAlphabetIndex([]rune(input))
	indexedKey := toAlphabetIndex([]rune(key))
	text = modSequence(text, indexedKey)
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

func toAlphabetIndex(text []rune) []rune {
	for i := range text {
		text[i] -= 'а'
	}
	return text
}

func toAlphabetCode(text []rune) []rune {
	for i := range text {
		text[i] += 'а'
	}
	return text
}

func modSequence(text, key []rune) []rune {
	for i := range text {
		text[i] = mod(text[i], key[i])
	}
	return text
}

func mod(txtSym, keySym rune) rune {
	return (txtSym + keySym) % 32
}
