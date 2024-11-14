package parser

import "github.com/kamijoucen/genginex/internal/base"

// 运算符优先级
var operatorPrecedence = map[base.TokenType]int32{
	base.Or:        10,
	base.And:       20,
	base.Eq:        30,
	base.Ne:        30,
	base.Gt:        30,
	base.Ge:        30,
	base.Lt:        30,
	base.Le:        30,
	base.Add:       40,
	base.Sub:       40,
	base.Mul:       50,
	base.Div:       50,
	base.Mod:       50,
	base.Not:       60,
	base.DoubleAdd: 70,
	base.DoubleMul: 70,
	base.Assign:    80,
	base.AddAssign: 80,
	base.SubAssign: 80,
	base.MulAssign: 80,
	base.DivAssign: 80,
	base.ModAssign: 80,
}
