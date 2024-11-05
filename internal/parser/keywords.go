package parser

var keywords = map[string]TokenType{
	"if":    If,
	"else":  Else,
	"begin": Begin,
	"end":   End,
	"rule":  Rule,
}

var symbols = map[string]TokenType{
	// 括号
	"(": LParen,
	")": RParen,
	"{": LBrace,
	"}": RBrace,
	"[": LBracket,
	"]": RBracket,

	// 算术运算符
	"+":  Add,
	"-":  Sub,
	"*":  Mul,
	"/":  Div,
	"%":  Mod,
	"++": DoubleAdd,
	"--": DoubleSub,
	"**": DoubleMul,

	// 算术赋值运算符
	"+=": AddAssign,
	"-=": SubAssign,
	"*=": MulAssign,
	"/=": DivAssign,
	"%=": ModAssign,

	// 逻辑运算符
	"!":  Not,
	"&&": And,
	"||": Or,

	// 比较运算符
	">":  Gt,
	">=": Ge,
	"<":  Lt,
	"<=": Le,
	"==": Eq,
	"!=": Ne,

	// 赋值运算符
	"=": Assign,

	// 其他符号
	",":   Comma,
	";":   Semicolon,
	".":   Dot,
	"?":   Question,
	":":   Colon,
	"->":  Arrow,
	"...": Ellipsis,
}
