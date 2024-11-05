package parser

import (
	"fmt"
)

type Lexical struct {
	offset     int32
	line       int32
	column     int32
	curStrFlag rune
	content    []rune
}

func NewLexical(content string) *Lexical {
	lex := &Lexical{
		offset:  0,
		content: []rune(content),
	}
	return lex
}

// Next
func (l *Lexical) Next() *Token {
	var token *Token
	match := false
	for !match && l.offset < int32(len(l.content)) {
		curChar := l.content[l.offset]
		if IsIdentifierFirstChar(curChar) {
			token = l.scanIdentifier()
			match = true
		} else if IsDigitChar(curChar) {
			// l.scanNumber()
		} else if curChar == '"' || curChar == '\'' {
			l.curStrFlag = curChar
			token = l.scanString()
			match = true
		} else if IsSpaceChar(curChar) {
			l.skipSpace()
		} else {
			token = l.scanSymbol()
			match = true
		}
	}
	return token
}

func (l *Lexical) forward() {
	if l.content[l.offset] == '\n' {
		l.line++
		l.column = 0
	} else {
		l.column++
	}
	l.offset++
}

func (l *Lexical) forwardN(n int32) {
	for i := int32(0); i < n; i++ {
		l.forward()
	}
}

// skipSpace
func (l *Lexical) skipSpace() {
	for l.offset < int32(len(l.content)) {
		if IsSpaceChar(l.content[l.offset]) {
			l.forward()
		} else {
			break
		}
	}
}

// scanString
func (l *Lexical) scanString() *Token {
	start := l.offset
	quote := l.curStrFlag

	l.forward()

	for l.offset < int32(len(l.content)) {
		c := l.content[l.offset]
		if c == quote {
			l.forward()
			break
		} else {
			l.forward()
		}
	}

	if l.offset > int32(len(l.content)) {
		panic(fmt.Sprintf("未结束的字符串在 %d:%d", l.line, l.column))
	}
	str := string(l.content[start:l.offset])
	return &Token{
		Type:   String,
		Str:    str,
		Offset: start,
	}
}

func (l *Lexical) scanIdentifier() *Token {
	start := l.offset
	l.forward()
	for l.offset < int32(len(l.content)) {
		if !IsIdentifierChar(l.content[l.offset]) {
			break
		}
		l.forward()
	}
	str := string(l.content[start:l.offset])
	stype := Identifier
	if t, ok := keywords[str]; ok {
		stype = t
	}
	return &Token{
		Type:   stype,
		Str:    str,
		Offset: start,
	}
}

// scanSymbol
func (l *Lexical) scanSymbol() *Token {
	start := l.offset
	var symbol string
	var tokenType TokenType
	found := false

	maxLen := 3
	for length := maxLen; length > 0; length-- {
		if l.offset+int32(length) > int32(len(l.content)) {
			continue
		}
		temp := string(l.content[l.offset : l.offset+int32(length)])
		if tType, ok := symbols[temp]; ok {
			symbol = temp
			tokenType = tType
			found = true
			l.forwardN(int32(length))
			break
		}
	}
	if !found {
		panic(fmt.Sprintf("unexpected symbol '%s' at %d:%d", string(l.content[l.offset]), l.line, l.column))
	}
	return &Token{
		Type:   tokenType,
		Str:    symbol,
		Offset: start,
	}
}
