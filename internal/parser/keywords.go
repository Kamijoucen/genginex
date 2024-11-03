package parser

var keywords = map[string]TokenType{
	"if":    IF,
	"else":  ELSE,
	"begin": BEGIN,
	"end":   END,
	"rule":  RULE,
}
