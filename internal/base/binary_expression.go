package base

import (
	"reflect"

	"github.com/kamijoucen/genginex/context"
)

type BinaryExpression struct {
	SourceCode
	Lhs      ExpressionNode
	Rhs      ExpressionNode
	Operator string
}

// Evaluate the binary expression
func (b *BinaryExpression) Evaluate(dc *context.DataContext, Vars map[string]reflect.Value) (reflect.Value, error) {
	return reflect.ValueOf(nil), nil
}
