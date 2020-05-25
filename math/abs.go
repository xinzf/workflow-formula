package math

import (
	"fmt"
	"github.com/Knetic/govaluate"
	"math"
	//"sort"
)

type Abs struct {
	values map[string]interface{}
}

func (this *Abs) GetName() string {
	return "ABS"
}

func (*Abs) GetCategory() string {
	return "数学函数"
}

func (*Abs) GetDescription() string {
	return `ABS(number)
返回数字的绝对值。`
}

func (this *Abs) SetValues(values map[string]interface{}) {
	this.values = values
}

func (*Abs) GetFunc() govaluate.ExpressionFunction {
	return func(arguments ...interface{}) (interface{}, error) {
		val, flag := arguments[0].(float64)
		if !flag {
			return nil, fmt.Errorf("ABS: 参数不是有效的 float 类型")
		}
		return math.Abs(val), nil
	}
}
