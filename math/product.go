package math

import (
	"github.com/Knetic/govaluate"
	"github.com/cstockton/go-conv"
)

type Product struct {
	values map[string]interface{}
}

func (this *Product) GetName() string {
	return "PRODUCT"
}

func (*Product) GetCategory() string {
	return "数学函数"
}

func (*Product) GetDescription() string {
	return `PRODUCT(number1, [number2], …)
函数使所有以参数形式给出的数字相乘并返回乘积。`
}

func (this *Product) SetValues(values map[string]interface{}) {
	this.values = values
}

func (*Product) GetFunc() govaluate.ExpressionFunction {
	return func(arguments ...interface{}) (interface{}, error) {
		res := 1.0
		for _, param := range arguments {
			val, _ := conv.Float64(param)
			//val, flag := param.(float64)
			//if !flag {
			//	return nil, fmt.Errorf("CONCAT: 第 %d 个参数不是有效的 float64 类型", key)
			//}
			res = val * res
		}
		return res, nil
	}
}
