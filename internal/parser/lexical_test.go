package parser

import (
	"testing"
)

// TestIfLexer
func TestIfLexer(t *testing.T) {
	rule_test := `
	lisicen = 1
	'"1"' '"2"' "" '' 1.5 5
	`
	lex := NewLexical(rule_test)
	t1 := lex.Next()
	if t1.Type != Identifier {
		t.Error("t1.Type != Identifier")
	}
	t2 := lex.Next()
	if t2.Type != Assign {
		t.Error("t2.Type != Assign")
	}
	t3 := lex.Next()
	if t3.Type != Int {
		t.Error("t3.Type != Int")
	}
	t4 := lex.Next()
	if t4.Type != String {
		t.Error("t4.Type != String")
	}
	t5 := lex.Next()
	if t5.Type != String {
		t.Error("t5.Type != String")
	}
	t6 := lex.Next()
	if t6.Type != String {
		t.Error("t6.Type != String")
	}
	t7 := lex.Next()
	if t7.Type != String {
		t.Error("t7.Type != String")
	}
	t8 := lex.Next()
	if t8.Type != Float {
		t.Error("t7.Type != Float")
	}
	t9 := lex.Next()
	if t9.Type != Int {
		t.Error("t8.Type != Int")
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
