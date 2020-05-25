package logic

import (
	"errors"
	"github.com/Knetic/govaluate"
)

type Not struct {
	values map[string]interface{}
}

func (this *Not) GetName() string {
	return "NOT"
}

func (*Not) GetCategory() string {
	return "逻辑函数"
}

func (*Not) GetDescription() string {
	return `NOT函数返回与指定表达式相反的布尔值。
用法：NOT(逻辑表达式)
示例：NOT(语文成绩>60)，如果语文成绩大于60返回false，否则返回true`
}

func (this *Not) SetValues(values map[string]interface{}) {
	this.values = values
}

func (this *Not) GetFunc() govaluate.ExpressionFunction {
	return func(arguments ...interface{}) (interface{}, error) {
		if len(arguments) != 1 {
			return nil, errors.New("NOT: 参数数量不正确")
		}

		val, flag := arguments[0].(bool)
		if !flag {
			return nil, errors.New("NOT: 参数不是有效的 Boolean 数据")
		}

		if val == true {
			return false, nil
		} else {
			return true, nil
		}
		return nil, nil
	}
}
