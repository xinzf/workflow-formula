package logic

import "github.com/Knetic/govaluate"

type False struct {
}

func (this *False) GetName() string {
	return "FALSE"
}

func (*False) GetCategory() string {
	return "逻辑函数"
}

func (*False) GetDescription() string {
	return `FALSE函数返回布尔值false
用法：FALSE()
示例：略`
}

func (*False) SetValues(values map[string]interface{}) {

}

func (*False) GetFunc() govaluate.ExpressionFunction {
	return func(arguments ...interface{}) (interface{}, error) {
		return false, nil
	}
}
