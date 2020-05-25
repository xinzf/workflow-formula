package math

import (
	"fmt"
	"github.com/Knetic/govaluate"
	"math"
	//"sort"
)

type Round struct {
	values map[string]interface{}
}

func (this *Round) GetName() string {
	return "ROUND"
}

func (*Round) GetCategory() string {
	return "数学函数"
}

func (*Round) GetDescription() string {
	return `ROUND(number, num_digits)
将数字四舍五入到指定的位数。
number: 必需。 要四舍五入的数字。
num_digits: 必需。 要进行四舍五入运算的位数。`
}

func (this *Round) SetValues(values map[string]interface{}) {
	this.values = values
}

func (*Round) GetFunc() govaluate.ExpressionFunction {
	return func(arguments ...interface{}) (interface{}, error) {
		number, n_flag := arguments[0].(float64)
		num_digit, d_flag := arguments[1].(float64)
		if !n_flag {
			return nil, fmt.Errorf("ROUND: 第一个参数不是有效的 float 类型")
		}
		if !d_flag {
			return nil, fmt.Errorf("ROUND: 第二个参数不是有效的 float 类型")
		}

		num_digits := int(num_digit)
		rs := math.Pow10(num_digits)
		tmp := math.Floor(number*rs + 0.5)
		res := tmp / rs
		return res, nil
	}
}
