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
	DoubleSub                    // --
	DoubleMul                    // **
	DoubleQuote                  // "
	SingleQuote                  // '
	If                           // if
	Else                         // else
	Begin                        // begin
	End                          // end
	Rule                         // rule
)

type Token struct {
	Type   TokenType
	Str    string
	Offset int32
}
