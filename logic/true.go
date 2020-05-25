package logic

import "github.com/Knetic/govaluate"

type True struct {
}

func (this *True) GetName() string {
	return "TRUE"
}

func (*True) GetCategory() string {
	return "逻辑函数"
}

func (*True) GetDescription() string {
	return `TRUE函数返回布尔值true
用法：TRUE()
示例：略`
}

func (*True) SetValues(values map[string]interface{}) {

}

func (*True) GetFunc() govaluate.ExpressionFunction {
	return func(arguments ...interface{}) (interface{}, error) {
		return true, nil
	}
}
