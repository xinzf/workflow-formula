package logic

import (
	"errors"
	"github.com/Knetic/govaluate"
)

type Or struct {
	values map[string]interface{}
}

func (this *Or) GetName() string {
	return "OR"
}

func (*Or) GetCategory() string {
	return "逻辑函数"
}

func (*Or) GetDescription() string {
	return `如果任意参数为真，OR 函数返回布尔值true；如果所有参数为假，返回布尔值false。
用法：OR(逻辑表达式1,逻辑表达式2,...)
示例：OR(语文成绩>90,数学成绩>90,英语成绩>90)，任何一门课成绩> 90，返回true，否则返回false`
}

func (this *Or) SetValues(values map[string]interface{}) {
	this.values = values
}

func (this *Or) GetFunc() govaluate.ExpressionFunction {
	return func(arguments ...interface{}) (interface{}, error) {
		if len(arguments) == 0 {
			return nil, errors.New("OR: 参数数量不足")
		}
		data := false
		for _, a := range arguments {
			val, flag := a.(bool)
			if !flag {
				return nil, errors.New("OR: 不是有效 Boolean 数据")
			}
			if val == true {
				data = true
				break
			}
		}

		return data, nil
	}
}
