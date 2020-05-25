package math

import (
	"errors"
	"github.com/Knetic/govaluate"
	"github.com/cstockton/go-conv"
)

type Int struct {
	values map[string]interface{}
}

func (this *Int) GetName() string {
	return "INT"
}

func (*Int) GetCategory() string {
	return "数学函数"
}

func (*Int) GetDescription() string {
	return `INT(number)
将数据转换为整形。`
}

func (this *Int) SetValues(values map[string]interface{}) {
	this.values = values
}

func (*Int) GetFunc() govaluate.ExpressionFunction {
	return func(arguments ...interface{}) (interface{}, error) {
		if len(arguments) == 0 {
			return nil, errors.New("INT: 参数不足")
		}

		return conv.Int(arguments[0])
	}
}
