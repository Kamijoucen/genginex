package parser

import (
	"fmt"
	"testing"
)

const rule_else_if_test = `
'"lisicen"' '' test
`

// TestIfLexer
func TestIfLexer(t *testing.T) {
	lex := NewLexical(rule_else_if_test)

	for {
		token := lex.Next()
		if token == nil {
			break
		}
		fmt.Println(token)
		t.Log(token)
	}

}

// TestIsIdentifierFirstChar
func TestIsIdentifierFirstChar(t *testing.T) {
	// 测试字母
	if !IsIdentifierFirstChar('a') {
		t.Error("IsIdentifierFirstChar('a') failed")
	}
	if !IsIdentifierFirstChar('A') {
		t.Error("IsIdentifierFirstChar('A') failed")
	}
	// 测试数字
	if IsIdentifierFirstChar('1') {
		t.Error("IsIdentifierFirstChar('1') should fail")
	}
	// 测试符号
	if IsIdentifierFirstChar('!') {
		t.Error("IsIdentifierFirstChar('!') should fail")
	}
	// 测试中文
	if !IsIdentifierFirstChar('中') {
		t.Error("IsIdentifierFirstChar('中') failed")
	}
	// 测试下划线
	if !IsIdentifierFirstChar('_') {
		t.Error("IsIdentifierFirstChar('_') failed")
	}
	// 测试美元符号
	if !IsIdentifierFirstChar('$') {
		t.Error("IsIdentifierFirstChar('$') failed")
	}
	// 测试空格
	if IsIdentifierFirstChar(' ') {
		t.Error("IsIdentifierFirstChar(' ') should fail")
	}
	// 测试其他语言 Unicode 字符
	if !IsIdentifierFirstChar('ñ') {
		t.Error("IsIdentifierFirstChar('ñ') failed")
	}
}
