package base

type TokenType uint8

const (
	UnKnow      TokenType = iota // 文件结束
	Identifier                   // 标识符
	Float                        // 浮点数
	Int                          // 整数
	String                       // 字符串
	Bool                         // 布尔
	Keyword                      // 关键字
	LParen                       // (
	RParen                       // )
	LBrace                       // {
	RBrace                       // }
	LBracket                     // [
	RBracket                     // ]
	And                          // &&
	Or                           // ||
	Not                          // !
	Assign                       // =
	AddAssign                    // +=
	SubAssign                    // -=
	MulAssign                    // *=
	DivAssign                    // /=
	ModAssign                    // %=
	Add                          // +
	Sub                          // -
	Mul                          // *
	Div                          // /
	Mod                          // %
	Eq                           // ==
	Ne                           // !=
	Gt                           // >
	Ge                           // >=
	Lt                           // <
	Le                           // <=
	Comma                        // ,
	Semicolon                    // ;
	Dot                          // .
	Arrow                        // ->
	Ellipsis                     // ...
	Question                     // ?
	Colon                        // :
	DoubleAdd                    // ++
	DoubleMul                    // **
	DoubleQuote                  // "
	SingleQuote                  // '
	If                           // if
	Else                         // else
	Begin                        // begin
	End                          // end
	Rule                         // rule
)

var keywords = map[string]TokenType{
	"if":    If,
	"else":  Else,
	"begin": Begin,
	"end":   End,
	"rule":  Rule,
}

// GetKeywordType
func GetKeywordType(str string) TokenType {
	if t, ok := keywords[str]; ok {
		return t
	}
	return UnKnow
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

// GetSymbolType
func GetSymbolType(str string) TokenType {
	if t, ok := symbols[str]; ok {
		return t
	}
	return UnKnow
}

type Token struct {
	Type   TokenType
	Str    string
	Offset int32
}
