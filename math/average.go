package math

import (
	"fmt"
	"github.com/Knetic/govaluate"
)

type Average struct {
	values map[string]interface{}
}

func (this *Average) GetName() string {
	return "AVERAGE"
}

func (*Average) GetCategory() string {
	return "数学函数"
}

func (*Average) GetDescription() string {
	return `AVERAGE函数可以获取一组数值的算术平均值
用法：AVERAGE(数字1,数字2,...)
示例：AVERAGE(语文成绩,数学成绩, 英语成绩)返回三门课程的平均分`
}

func (this *Average) SetValues(values map[string]interface{}) {
	this.values = values
}

func (*Average) GetFunc() govaluate.ExpressionFunction {
	return func(arguments ...interface{}) (interface{}, error) {
		var total float64
		for i := 0; i < len(arguments); i++ {
			if val, flag := arguments[i].(float64); !flag {
				return nil, fmt.Errorf("AVERAGE: 第 %d 个参数不是有效的数字格式", i)
			} else {
				total += val
			}
		}

		return total / float64(len(arguments)), nil
	}
}
