package test

import (
	"fmt"
	"testing"
	"time"

	"github.com/kamijoucen/genginex/builder"
	"github.com/kamijoucen/genginex/context"
	"github.com/kamijoucen/genginex/engine"
)

const rule_else_if_test = `
rule "elseif_test" "test"
begin

a = 8
if a < 1 {
	println("a < 1")
} else if a >= 1 && a <6 {
	println("a >= 1 && a <6")
} else if a >= 6 && a < 7 {
	println("a >= 6 && a < 7")
} else if a >= 7 && a < 10 {
	println("a >=7 && a < 10")
} else        {
	println("a > 10")
}

end
`

func Test_elseif(t *testing.T) {

	// 运行上下文
	dataContext := context.NewDataContext()

	// 外部变量

	dataContext.Add("println", fmt.Println)

	//init rule engine
	ruleBuilder := builder.NewRuleBuilder(dataContext)

	//resolve rules from string
	start1 := time.Now().UnixNano()

	// 编译规则，缓存到builder中
	// 使用antlr4解析规则
	err := ruleBuilder.BuildRuleFromString(rule_else_if_test)
	end1 := time.Now().UnixNano()

	println(fmt.Sprintf("rules num:%d, load rules cost time:%d ns", len(ruleBuilder.Kc.RuleEntities), end1-start1))

	if err != nil {
		panic(err)
	}

	// engine实体，按道理，编译到ast就可以直接run
	eng := engine.NewGengine()

	start := time.Now().UnixNano()
	// true: means when there are many rules， if one rule execute error，continue to execute rules after the occur error rule

	// 遍历ast
	err = eng.Execute(ruleBuilder, true)
	end := time.Now().UnixNano()
	if err != nil {
		panic(err)
	}
	println(fmt.Sprintf("execute rule cost %d ns", end-start))

}
