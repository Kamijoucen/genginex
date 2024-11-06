package parser

type TokenType uint8

const (
	Eof         TokenType = iota // 文件结束
	Identifier                   // 标识符
	Float                        // 浮点数
	Int                          // 整数
	String                       // 字符串
	Bool                         // 布尔
	Keyword                      // 关键字
	LParen                       // 左括号
	RParen                       // 右括号
	LBrace                       // 左大括号
	RBrace                       // 右大括号
	LBracket                     // 左中括号
	RBracket                     // 右中括号
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

// 运算符优先级
var operatorPrecedence = map[TokenType]int32{
	Or:        10,
	And:       20,
	Eq:        30,
	Ne:        30,
	Gt:        30,
	Ge:        30,
	Lt:        30,
	Le:        30,
	Add:       40,
	Sub:       40,
	Mul:       50,
	Div:       50,
	Mod:       50,
	Not:       60,
	DoubleAdd: 70,
	DoubleMul: 70,
	Assign:    80,
	AddAssign: 80,
	SubAssign: 80,
	MulAssign: 80,
	DivAssign: 80,
	ModAssign: 80,
}

type Token struct {
	Type   TokenType
	Str    string
	Offset int32
}
