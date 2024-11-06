package parser

import (
	"fmt"
	"strconv"

	"github.com/kamijoucen/genginex/internal/base"
)

type Parser struct {
	tokens []*Token
	offset int32
}

func NewParser(tks []*Token) *Parser {
	return &Parser{
		tokens: tks,
	}
}

// ParseStatements
func (p *Parser) ParseStatements() (*base.Statements, error) {

	

	return nil, nil
}

// ParseRuleEntity
func (p *Parser) ParseRuleEntity() (*base.RuleEntity, error) {
	if _, err := p.checkTokenAndForward(Rule); err != nil {
		return nil, err
	}
	nameToken, err := p.checkTokenAndForward(String)
	if err != nil {
		return nil, err
	}
	if nameToken.Str == "" {
		return nil, fmt.Errorf("rule name is empty")
	}
	ruleDesc := ""
	if p.currentToken().Type == String {
		ruleDesc = p.currentToken().Str
		_ = p.nextToken()
	}
	salience := int64(0)
	if p.currentToken().Type == Int {
		temp := p.currentToken().Str
		salience, _ = strconv.ParseInt(temp, 10, 64)
		_ = p.nextToken()
	}
	if _, err := p.checkTokenAndForward(Begin); err != nil {
		return nil, err
	}

	// TODO parse rule content

	if _, err := p.checkTokenAndForward(End); err != nil {
		return nil, err
	}
	t := &base.RuleEntity{
		RuleName:        nameToken.Str,
		Salience:        salience,
		RuleDescription: ruleDesc,
		RuleContent:     nil,
	}
	return t, nil
}

// nextToken
func (p *Parser) nextToken() *Token {
	p.offset++
	if p.offset >= int32(len(p.tokens)) {
		return nil
	}
	return p.tokens[p.offset]
}

// currentToken
func (p *Parser) currentToken() *Token {
	if p.offset >= int32(len(p.tokens)) {
		return nil
	}
	return p.tokens[p.offset]
}

// checkTokenAndForward
func (p *Parser) checkTokenAndForward(t TokenType) (*Token, error) {
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
