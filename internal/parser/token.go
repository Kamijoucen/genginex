package parser

type TokenType uint8

const (
	EOF          TokenType = iota // 文件结束
	Identifier                    // 标识符
	Float                         // 浮点数
	Int                           // 整数
	Char                          // 字符
	String                        // 字符串
	Bool                          // 布尔
	Keyword                       // 关键字
	L_PAREN                       // 左括号
	R_PAREN                       // 右括号
	L_BRACE                       // 左大括号
	R_BRACE                       // 右大括号
	L_BRACKET                     // 左中括号
	R_BRACKET                     // 右中括号
	AND                           // &&
	OR                            // ||
	NOT                           // !
	ASSIGN                        // =
	ADD_ASSIGN                    // +=
	SUB_ASSIGN                    // -=
	MUL_ASSIGN                    // *=
	DIV_ASSIGN                    // /=
	MOD_ASSIGN                    // %=
	ADD                           // +
	SUB                           // -
	MUL                           // *
	DIV                           // /
	MOD                           // %
	EQ                            // ==
	NE                            // !=
	GT                            // >
	GE                            // >=
	LT                            // <
	LE                            // <=
	COMMA                         // ,
	SEMICOLON                     // ;
	DOT                           // .
	ARROW                         // ->
	ELLIPSIS                      // ...
	QUESTION                      // ?
	COLON                         // :
	DOUBLE_ADD                    // ++
	DOUBLE_SUB                    // --
	DOUBLE_MUL                    // **
	DOUBLE_QUOTE                  // "
	SINGLE_QUOTE                  // '
	IF                            // if
	ELSE                          // else
	BEGIN                         // begin
	END                           // end
	RULE                          // rule
)

type Token struct {
	Type   TokenType
	Str    string
	Offset int32
}
