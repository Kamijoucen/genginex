package parser

import (
	"fmt"

	"github.com/kamijoucen/genginex/internal/base"
)

var EofToken = &base.Token{
	Type: base.Eof,
}

type Lexical struct {
	offset  int32
	line    int32
	column  int32
	content []rune
}

func NewLexical(content string) *Lexical {
	lex := &Lexical{
		content: []rune(content),
	}
	return lex
}

// Next
func (l *Lexical) Next() (tk *base.Token, err error) {
	if l.offset >= int32(len(l.content)) {
		return EofToken, nil
	}
	match := false
	for !match && l.offset < int32(len(l.content)) {
		curChar := l.content[l.offset]
		if IsIdentifierFirstChar(curChar) {
			tk, err = l.scanIdentifier()
			match = true
		} else if IsDigitChar(curChar) {
			tk, err = l.scanNumber()
			match = true
		} else if curChar == '"' || curChar == '\'' {
			tk, err = l.scanString()
			match = true
		} else if IsSpaceChar(curChar) {
			l.skipSpace()
		} else {
			tk, err = l.scanSymbol()
			match = true
		}
	}
	return
}

// scanNumber
func (l *Lexical) scanNumber() (*base.Token, error) {
	start := l.offset
	isFloat := false

	for l.offset < int32(len(l.content)) {
		curChar := l.content[l.offset]
		if IsDigitChar(curChar) {
			l.forward()
		} else if curChar == '.' && !isFloat {
			isFloat = true
			l.forward()
		} else {
			break
		}
	}
	str := string(l.content[start:l.offset])
	tokenType := base.Int
	if isFloat {
		tokenType = base.Float
	}
	return &base.Token{
		Type:   tokenType,
		Str:    str,
		Offset: start,
	}, nil
}

// scanString
func (l *Lexical) scanString() (*base.Token, error) {
	start := l.offset
	quote := l.content[l.offset]

	l.forward()

	for l.offset < int32(len(l.content)) {
		c := l.content[l.offset]
		l.forward()
		if c == quote {
			break
		}
	}

	if l.offset > int32(len(l.content)) {
		return nil, fmt.Errorf("未结束的字符串在 %d:%d", l.line, l.column)
	}
	str := string(l.content[start:l.offset])
	return &base.Token{
		Type:   base.String,
		Str:    str,
		Offset: start,
	}, nil
}

func (l *Lexical) scanIdentifier() (*base.Token, error) {
	start := l.offset
	l.forward()
	for l.offset < int32(len(l.content)) {
		if !IsIdentifierChar(l.content[l.offset]) {
			break
		}
		l.forward()
	}
	str := string(l.content[start:l.offset])
	stype := base.Identifier
	if t := base.GetKeywordType(str); t != base.UnKnow {
		stype = t
	}
	return &base.Token{
		Type:   stype,
		Str:    str,
		Offset: start,
	}, nil
}

// scanSymbol
func (l *Lexical) scanSymbol() (*base.Token, error) {
	start := l.offset
	var symbol string
	var tokenType base.TokenType
	found := false

	maxLen := 3
	for length := maxLen; length > 0; length-- {
		if l.offset+int32(length) > int32(len(l.content)) {
			continue
		}
		temp := string(l.content[l.offset : l.offset+int32(length)])
		if t := base.GetSymbolType(temp); t != base.UnKnow {
			symbol = temp
			tokenType = t
			found = true
			l.forwardN(int32(length))
			break
		}
	}
	if !found {
		return nil, fmt.Errorf("unexpected symbol '%s' at %d:%d", string(l.content[l.offset]), l.line, l.column)
	}
	return &base.Token{
		Type:   tokenType,
		Str:    symbol,
		Offset: start,
	}, nil
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
