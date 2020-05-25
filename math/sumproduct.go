package math

import (
	"errors"
	"fmt"
	"github.com/Knetic/govaluate"
)

type Sumproduct struct {
	values map[string]interface{}
}

func (this *Sumproduct) GetName() string {
	return "SUMPRODUCT"
}

func (*Sumproduct) GetCategory() string {
	return "数学函数"
}

func (*Sumproduct) GetDescription() string {
	return `SUMPRODUCT(array1, [array2])
在给定的数组中，将数组间对应的元素相乘，并返回乘积之和。`
}

func (this *Sumproduct) SetValues(values map[string]interface{}) {
	this.values = values
}

func (*Sumproduct) GetFunc() govaluate.ExpressionFunction {
	return func(arguments ...interface{}) (interface{}, error) {
		if len(arguments) != 2 {
			return nil, errors.New("CONCAT: 参数数量不对")
		}
		arr_one, flag_one := arguments[0].([]float64)
		arr_two, flag_two := arguments[1].([]float64)
		if !flag_one {
			return nil, fmt.Errorf("ROUND: 第一个参数不是有效的 []float 类型")
		}
		if !flag_two {
			return nil, fmt.Errorf("ROUND: 第二个参数不是有效的 []float 类型")
		}
		sum := 0.00
		for key, param := range arr_one {
			sum = param*arr_two[key] + sum
		}
		return sum, nil
	}
}
