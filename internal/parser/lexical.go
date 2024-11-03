package parser

import "unicode"

type Lexical struct {
	offset  int32
	line    int32
	column  int32
	content []rune
	tokens  []*Token
}

func NewLexical(content string) *Lexical {
	lex := &Lexical{
		offset:  0,
		content: []rune(content),
		tokens:  make([]*Token, 0),
	}
	lex.scan()
	return lex
}

// Next
func (l *Lexical) Next() *Token {
	return nil
}

// Peek
func (l *Lexical) Peek() *Token {
	return nil
}

// scan
func (l *Lexical) scan() {
	for l.offset < int32(len(l.content)) {
		if unicode.IsSpace(l.content[l.offset]) {
			l.offset++
		} else {
			l.scanToken()
		}
	}
}

// nextChar
func (l *Lexical) nextChar() {
	l.offset++
}

// scanToken
func (l *Lexical) scanToken() {
	curChar := l.content[l.offset]
	if IsIdentifierFirstChar(curChar) {
		l.scanIdentifier()
	} else if IsDigitChar(curChar) {
		// l.scanNumber()
	} else {
		// l.scanSymbol()
	}
}

// scanIdentifier
func (l *Lexical) scanIdentifier() {
	start := l.offset
	l.offset++
	for l.offset < int32(len(l.content)) {
		if !IsIdentifierChar(l.content[l.offset]) {
			break
		}
		l.offset++
	}
	str := string(l.content[start:l.offset])
	stype := Identifier
	if t, ok := keywords[str]; ok {
		stype = t
	}
	l.tokens = append(l.tokens, &Token{
		Type:   stype,
		Str:    str,
		Offset: start,
	})
}
g