package jsonconv

import "unicode"

// Utility functions to be used in templates

func firstLetterLowercase(str string) string {
	runes := []rune(str)
	runes[0] = unicode.ToLower(runes[0])
	return string(runes)
}
