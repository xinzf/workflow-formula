package logic

import (
	"errors"
	"github.com/Knetic/govaluate"
)

type If struct {
	values map[string]interface{}
}

func (this *If) GetName() string {
	return "IF"
}

func (*If) GetCategory() string {
	return "逻辑函数"
}

func (*If) GetDescription() string {
	return `IF函数判断一个条件能否满足；如果满足返回一个值，如果不满足则返回另外一个值
用法：IF(逻辑表达式,为true时返回的值,为false时返回的值)
示例：IF(语文成绩>60,"及格","不及格")，当语文成绩>60时返回及格，否则返回不及格。`
}

func (this *If) SetValues(values map[string]interface{}) {
	this.values = values
}

func (this *If) GetFunc() govaluate.ExpressionFunction {

	return func(arguments ...interface{}) (interface{}, error) {
		if len(arguments) != 3 {
			return nil, errors.New("argument number is not validate")
		}

		judge := arguments[0]
		trueValue := arguments[1]
		falseValue := arguments[2]

		val, flag := judge.(bool)
		if !flag {
			return nil, errors.New("IF: 第一个参数不是有效的 Boolean 数据")
		}

		if val == true {
			return trueValue, nil
		}
		return falseValue, nil
	}
}
