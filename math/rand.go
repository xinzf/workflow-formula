package math

import (
	"errors"
	"github.com/Knetic/govaluate"
	"math/rand"
)

type Rand struct {
	values map[string]interface{}
}

func (this *Rand) GetName() string {
	return "RAND"
}

func (*Rand) GetCategory() string {
	return "数学函数"
}

func (*Rand) GetDescription() string {
	return `RAND()
返回大于等于 0 且小于 1 的均匀分布随机实数。每一次触发计算都会变化。`
}

func (this *Rand) SetValues(values map[string]interface{}) {
	this.values = values
}

func (*Rand) GetFunc() govaluate.ExpressionFunction {
	return func(arguments ...interface{}) (interface{}, error) {
		if len(arguments) > 0 {
			return nil, errors.New("RAND: 参数数量不对")
		}
		return rand.Float64(), nil
	}
}
