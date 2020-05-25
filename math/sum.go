package math

import (
	"fmt"
	"github.com/Knetic/govaluate"
)

type Sum struct {
	values map[string]interface{}
}

func (this *Sum) GetName() string {
	return "SUM"
}

func (*Sum) GetCategory() string {
	return "数学函数"
}

func (*Sum) GetDescription() string {
	return `SUM(number1, [number2], …)
函数使所有以参数形式给出的数字相加并返回和。`
}

func (this *Sum) SetValues(values map[string]interface{}) {
	this.values = values
}

func (*Sum) GetFunc() govaluate.ExpressionFunction {
	return func(arguments ...interface{}) (interface{}, error) {
		res := 0.00
		for key, param := range arguments {
			val, flag := param.(float64)
			if !flag {
				return nil, fmt.Errorf("CONCAT: 第 %d 个参数不是有效的 float64 类型", key+1)
			}
			res = val + res
		}
		return res, nil
	}
}
