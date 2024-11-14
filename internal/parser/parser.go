package parser

import (
	"fmt"
	"strconv"

	"github.com/kamijoucen/genginex/internal/base"
)

type Parser struct {
	tokens []*base.Token
	offset int32
}

func NewParser(tks []*base.Token) *Parser {
	return &Parser{
		tokens: tks,
	}
}

// ParseStatements
func (p *Parser) ParseStatements() (*base.Statements, error) {
	return nil, nil
}

// ParseExpression
func (p *Parser) ParseExpression() (base.ExpressionNode, error) {
	lhs, err := p.ParsePrimary()
	if err != nil {
		return nil, err
	}
	return p.parseBinaryOpRhs(lhs, int32(0))
}

// PricePrimary
func (p *Parser) ParsePrimary() (base.ExpressionNode, error) {
	return nil, nil
}

// parseBinaryOpRhs
func (p *Parser) parseBinaryOpRhs(expr base.ExpressionNode, prec int32) (base.ExpressionNode, error) {
	// 此时 expr 是 1
	for {
		// curTok 是 +
		curTok := p.currentToken()
		if curTok == nil {
			return nil, nil
		}
		// + 的优先级是 40
		tokPrec := operatorPrecedence[curTok.Type]
		// 40 < 0 == false
		// 如果上一个运算符的优先级, 大于等于当前运算符的优先级, 则返回
		if tokPrec < prec {
			return expr, nil
		}
		// eat +
		_ = p.nextToken()
		// rhs 是 2
		rhs, err := p.ParsePrimary()
		if err != nil {
			return nil, err
		}
		// nextTok 是 *
		nextTok := p.currentToken()
		// nextPrec 是 50
		nextPrec := operatorPrecedence[nextTok.Type]
		// 50 > 40  * 的优先级高于 +
		// 这里就是处理下一个运算符优先级更高的情况
		if tokPrec < nextPrec {
			// 实参的rhs是 2, tokPrec 是 40
			rhs, err = p.parseBinaryOpRhs(rhs, tokPrec+1)
			if err != nil {
				return nil, err
			}
		}
		expr = &base.BinaryExpression{
			Operator: "",
			Lhs:      expr,
			Rhs:      rhs,
		}
	}
}

// ParseRuleEntity
func (p *Parser) ParseRuleEntity() (*base.RuleEntity, error) {
	if _, err := p.checkTokenAndForward(base.Rule); err != nil {
		return nil, err
	}
	nameToken, err := p.checkTokenAndForward(base.String)
	if err != nil {
		return nil, err
	}
	if nameToken.Str == "" {
		return nil, fmt.Errorf("rule name is empty")
	}
	ruleDesc := ""
	if p.currentToken().Type == base.String {
		ruleDesc = p.currentToken().Str
		_ = p.nextToken()
	}
	salience := int64(0)
	if p.currentToken().Type == base.Int {
		temp := p.currentToken().Str
		salience, _ = strconv.ParseInt(temp, 10, 64)
		_ = p.nextToken()
	}
	if _, err := p.checkTokenAndForward(base.Begin); err != nil {
		return nil, err
	}

	sts, err := p.ParseStatements()
	if err != nil {
		return nil, err
	}

	if _, err := p.checkTokenAndForward(base.End); err != nil {
		return nil, err
	}
	t := &base.RuleEntity{
		RuleName:        nameToken.Str,
		Salience:        salience,
		RuleDescription: ruleDesc,
		RuleContent: &base.RuleContent{
			Statements: sts,
		},
	}
	return t, nil
}

// nextToken
func (p *Parser) nextToken() *base.Token {
	p.offset++
	if p.offset >= int32(len(p.tokens)) {
		return nil
	}
	return p.tokens[p.offset]
}

// currentToken
func (p *Parser) currentToken() *base.Token {
	if p.offset >= int32(len(p.tokens)) {
		return nil
	}
	return p.tokens[p.offset]
}

// checkTokenAndForward
func (p *Parser) checkTokenAndForward(t base.TokenType) (*base.Token, error) {
	if p.offset >= int32(len(p.tokens)) {
		return nil, fmt.Errorf("unexpected end of rule")
	}
	curTok := p.tokens[p.offset]
	if curTok.Type != t {
		return nil, fmt.Errorf("unexpected token %d", p.tokens[p.offset].Type)
	}
	p.offset++
	return curTok, nil
}
