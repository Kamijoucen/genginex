package parser

var keywords = map[string]TokenType{
	"if":    IF,
	"else":  ELSE,
	"begin": BEGIN,
	"end":   END,
	"rule":  RULE,
}

var symbols = map[string]TokenType{
	// 括号
	"(": L_PAREN,
	")": R_PAREN,
	"{": L_BRACE,
	"}": R_BRACE,
	"[": L_BRACKET,
	"]": R_BRACKET,

	// 算术运算符
	"+":  ADD,
	"-":  SUB,
	"*":  MUL,
	"/":  DIV,
	"%":  MOD,
	"++": DOUBLE_ADD,
	"--": DOUBLE_SUB,
	"**": DOUBLE_MUL,

	// 算术赋值运算符
	"+=": ADD_ASSIGN,
	"-=": SUB_ASSIGN,
	"*=": MUL_ASSIGN,
	"/=": DIV_ASSIGN,
	"%=": MOD_ASSIGN,

	// 逻辑运算符
	"!":  NOT,
	"&&": AND,
	"||": OR,

	// 比较运算符
	">":  GT,
	">=": GE,
	"<":  LT,
	"<=": LE,
	"==": EQ,
	"!=": NE,

	// 赋值运算符
	"=": ASSIGN,

	// 其他符号
	",":   COMMA,
	";":   SEMICOLON,
	".":   DOT,
	"?":   QUESTION,
	":":   COLON,
	"->":  ARROW,
	"...": ELLIPSIS,
}
