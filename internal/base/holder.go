package base

import (
	"reflect"

	"github.com/kamijoucen/genginex/context"
)

type ExpressionNode interface {
	Evaluate(dc *context.DataContext, Vars map[string]reflect.Value) (reflect.Value, error)
}

type ExpressionHolder interface {
	AcceptExpression(expression *Expression) error
}

type ExpressionAtomHolder interface {
	AcceptExpressionAtom(atom *ExpressionAtom) error
}

type MathExpressionHolder interface {
	AcceptMathExpression(mh *MathExpression) error
}

type MethodCallHolder interface {
	AcceptMethodCall(methodCall *MethodCall) error
}

type StatementsHolder interface {
	AcceptStatements(statement *Statements) error
}

type StringHolder interface {
	AcceptString(str string) error
}

type ThreeLevelCallHolder interface {
	AcceptThreeLevelCall(threeLevelCall *ThreeLevelCall) error
}

type VariableHolder interface {
	AcceptVariable(name string) error
}

type IntegerHolder interface {
	AcceptInteger(val int64) error
}

type MapVarHolder interface {
	AcceptMapVar(mv *MapVar) error
}

type FunctionCallHolder interface {
	AcceptFunctionCall(funcCall *FunctionCall) error
}

type ConstantHolder interface {
	AcceptConstant(cons *Constant) error
}

type AtSalienceHolder interface {
	AcceptSalience(val int64) error
}

type AtNameHolder interface {
	AcceptName(val string) error
}

type AtIdHolder interface {
	AcceptId(val int64) error
}

type AtDescHolder interface {
	AcceptDesc(val string) error
}

type AssignmentHolder interface {
	AcceptAssignment(a *Assignment) error
}

type ArgsHolder interface {
	AcceptArgs(funcArg *Args) error
}
