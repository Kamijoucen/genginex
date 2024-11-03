package parser

import "unicode"

// IsSpaceChar 空格字符
func IsSpaceChar(r rune) bool {
	return unicode.IsSpace(r)
}

// IsIdentifierFirstChar 字母, _, $, 或者 unicode 字母
func IsIdentifierFirstChar(r rune) bool {
	return unicode.IsLetter(r) || r == '_' || r == '$'
}

// IsIdentifierChar 字母, 数字, _, $, 或者 unicode 字母
func IsIdentifierChar(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r) || r == '_' || r == '$'
}

// IsDigitChar 数字
func IsDigitChar(r rune) bool {
	return unicode.IsDigit(r)
}
