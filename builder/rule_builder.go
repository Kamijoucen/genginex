package builder

import (
	"errors"
	"fmt"
	"sort"
	"strings"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr"

	"github.com/kamijoucen/genginex/context"
	"github.com/kamijoucen/genginex/internal/base"
	parser2 "github.com/kamijoucen/genginex/internal/iantlr/alr"
	"github.com/kamijoucen/genginex/internal/iparser"
	"github.com/kamijoucen/genginex/internal/parser"
	"github.com/kamijoucen/genginex/internal/tool"
)

type RuleBuilder struct {
	Kc *base.KnowledgeContext
	Dc *context.DataContext

	buildLock sync.Mutex
}

func NewRuleBuilder(dc *context.DataContext) *RuleBuilder {
	kc := base.NewKnowledgeContext()
	return &RuleBuilder{
		Kc: kc,
		Dc: dc,
	}
}

// chinese comment :全量更新
// if update success, all old rules will be delete and you inject new rules will be in the gengine
func (builder *RuleBuilder) BuildRuleFromString(ruleString string) error {
	builder.buildLock.Lock()
	defer builder.buildLock.Unlock()

	if strings.TrimSpace(ruleString) == "" {
		// nil ruleString check
		return fmt.Errorf("inject ruleString is %s", ruleString)
	}

	// AST存储
	kc := base.NewKnowledgeContext()

	lx := parser.NewLexical(ruleString)
	tks := make([]*base.Token, 0)

	for {
		tk, err := lx.Next()
		if err != nil {
			return err
		}
		if tk.Type == base.Eof {
			break
		}
		tks = append(tks, tk)
	}

	p := parser.NewParser(tks)

	stmts, err := p.ParseStatements()


	in := antlr.NewInputStream(ruleString)
	lexer := parser2.NewgengineLexer(in)
	lexerErrListener := iparser.NewGengineErrorListener()
	lexer.AddErrorListener(lexerErrListener)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// TODO listener内添加ast
	listener := iparser.NewGengineParserListener(kc)

	psr := parser2.NewgengineParser(stream)
	psr.BuildParseTrees = true
	// grammar listener
	grammarErrListener := iparser.NewGengineErrorListener()
	psr.AddErrorListener(grammarErrListener)
	antlr.ParseTreeWalkerDefault.Walk(listener, psr.Primary())

	if len(lexerErrListener.GrammarErrors) > 0 {
		return fmt.Errorf("%+v", lexerErrListener.GrammarErrors)
	}

	if len(grammarErrListener.GrammarErrors) > 0 {
		return errors.New(fmt.Sprintf("%+v", grammarErrListener.GrammarErrors))
	}

	if len(listener.ParseErrors) > 0 {
		return errors.New(fmt.Sprintf("%+v", listener.ParseErrors))
	}

	// sort
	for _, v := range kc.RuleEntities {
		kc.SortRules = append(kc.SortRules, v)
	}
	if len(kc.SortRules) > 1 {
		sort.SliceStable(kc.SortRules, func(i, j int) bool {
			return kc.SortRules[i].Salience > kc.SortRules[j].Salience
		})
	}

	for k, v := range kc.SortRules {
		kc.SortRulesIndexMap[v.RuleName] = k
	}

	builder.Kc = kc
	return nil
}

// chinese comment:增量更新
// if a rule already exists, this method will use the new rule to replace the old one
// if a rule doesn't exist, this method will add the new rule to the existed rules list
// in detail: copy from old -> update the copy -> use the updated copy to replace old
func (builder *RuleBuilder) BuildRuleWithIncremental(ruleString string) error {
	// make sure incremental update is thread safety!
	builder.buildLock.Lock()
	defer builder.buildLock.Unlock()

	if strings.TrimSpace(ruleString) == "" {
		// nil ruleString check
		return errors.New(fmt.Sprintf("incremental inject ruleString is %s", ruleString))
	}

	in := antlr.NewInputStream(ruleString)
	lexer := parser2.NewgengineLexer(in)
	lexerErrListener := iparser.NewGengineErrorListener()
	lexer.AddErrorListener(lexerErrListener)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	kc := base.NewKnowledgeContext()
	listener := iparser.NewGengineParserListener(kc)

	psr := parser2.NewgengineParser(stream)
	psr.BuildParseTrees = true

	grammarErrListener := iparser.NewGengineErrorListener()
	psr.AddErrorListener(grammarErrListener)
	antlr.ParseTreeWalkerDefault.Walk(listener, psr.Primary())

	if len(lexerErrListener.GrammarErrors) > 0 {
		return errors.New(fmt.Sprintf("%+v", lexerErrListener.GrammarErrors))
	}

	if len(grammarErrListener.GrammarErrors) > 0 {
		return errors.New(fmt.Sprintf("%+v", grammarErrListener.GrammarErrors))
	}

	if len(listener.ParseErrors) > 0 {
		return errors.New(fmt.Sprintf("%+v", listener.ParseErrors))
	}

	if len(kc.RuleEntities) == 0 {
		return errors.New(fmt.Sprintf("no rules need to update or add."))
	}

	// copy
	newRuleEntities := make(map[string]*base.RuleEntity, len(builder.Kc.RuleEntities))
	for mk, mv := range builder.Kc.RuleEntities {
		newRuleEntities[mk] = mv
	}

	// copy
	newSortRules := make([]*base.RuleEntity, len(builder.Kc.SortRules))
	for sk, sv := range builder.Kc.SortRules {
		newSortRules[sk] = sv
	}

	// kc store the new rules
	for k, v := range kc.RuleEntities {

		if vm, ok := newRuleEntities[k]; ok {
			// repalce update
			// search
			index := builder.Kc.SortRulesIndexMap[v.RuleName]
			if v.Salience == vm.Salience {
				// replace
				newSortRules[index] = v
			} else {
				newSortRules := append(newSortRules[:index], newSortRules[index+1:]...)
				// search location to insert
				low, mid := tool.BinarySearch(newSortRules, v.Salience)

				ire := []*base.RuleEntity{v}
				if mid == 0 {
					newRe := append(ire, newSortRules[low:]...)
					newSortRules = append(newSortRules[:low], newRe...)
				} else {
					newRe := append(ire, newSortRules[mid:]...)
					newSortRules = append(newSortRules[:mid], newRe...)
				}

				// update the sort index
				indexMap := make(map[string]int)
				for k, v := range newSortRules {
					indexMap[v.RuleName] = k
				}
				builder.Kc.SortRulesIndexMap = indexMap
			}

			newRuleEntities[k] = v
		} else {
			// add update
			low, mid := tool.BinarySearch(newSortRules, v.Salience)

			ire := []*base.RuleEntity{v}
			if mid == 0 {
				newRe := append(ire, newSortRules[low:]...)
				newSortRules = append(newSortRules[:low], newRe...)
			} else {
				newRe := append(ire, newSortRules[mid:]...)
				newSortRules = append(newSortRules[:mid], newRe...)
			}

			// update the sort index
			indexMap := make(map[string]int)
			for k, v := range newSortRules {
				indexMap[v.RuleName] = k
			}
			builder.Kc.SortRulesIndexMap = indexMap

			newRuleEntities[k] = v
		}
	}

	builder.Kc.RuleEntities = newRuleEntities
	builder.Kc.SortRules = newSortRules

	return nil
}

func (builder *RuleBuilder) RemoveRules(ruleNames []string) error {
	builder.buildLock.Lock()
	defer builder.buildLock.Unlock()

	if len(ruleNames) == 0 {
		return errors.New(fmt.Sprintf("no rules need to be remove! "))
	}

	newRuleEntities := make(map[string]*base.RuleEntity)
	for name, entity := range builder.Kc.RuleEntities {
		flag := true
		for _, delName := range ruleNames {
			if delName == name {
				flag = false
				break
			}
		}
		if flag {
			newRuleEntities[name] = entity
		}
	}

	newSortRuleEntities := make([]*base.RuleEntity, 0)
	for _, entity := range newRuleEntities {
		newSortRuleEntities = append(newSortRuleEntities, entity)
	}

	if len(newSortRuleEntities) > 1 {
		sort.SliceStable(newSortRuleEntities, func(i, j int) bool {
			return newSortRuleEntities[i].Salience > newSortRuleEntities[j].Salience
		})
	}

	newSortRulesIndexMap := make(map[string]int)
	for k, v := range newSortRuleEntities {
		newSortRulesIndexMap[v.RuleName] = k
	}

	kc := base.NewKnowledgeContext()

	kc.RuleEntities = newRuleEntities
	kc.SortRules = newSortRuleEntities
	kc.SortRulesIndexMap = newSortRulesIndexMap

	builder.Kc = kc
	return nil
}

func (builder *RuleBuilder) IsExist(ruleNames []string) []bool {
	builder.buildLock.Lock()
	defer builder.buildLock.Unlock()

	if len(ruleNames) == 0 {
		return make([]bool, 0)
	}

	exist := make([]bool, 0)
	for _, name := range ruleNames {
		_, ok := builder.Kc.RuleEntities[name]
		exist = append(exist, ok)
	}
	return exist
}
