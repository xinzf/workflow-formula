package char

import (
	"errors"
	"github.com/Knetic/govaluate"
	"strings"
)

type Lower struct {
	values map[string]interface{}
}

func (this *Lower) GetName() string {
	return "LOWER"
}

func (this *Lower) GetFunc() govaluate.ExpressionFunction {
	return func(arguments ...interface{}) (interface{}, error) {
		if len(arguments) == 0 {
			return nil, errors.New("LOWER: 参数不足")
		}

		val, flag := arguments[0].(string)
		if !flag {
			return nil, errors.New("LOWER: 参数不是有效字符类型")
		}

		return strings.ToLower(val), nil
	}
}

func (this *Lower) SetValues(values map[string]interface{}) {
	this.values = values
}

func (this *Lower) GetCategory() string {
	return "文本函数"
}

func (this *Lower) GetDescription() string {
	return `LOWER函数可以将一个文本中的所有大写字母转换为小写字母
用法：LOWER(文本)
示例：LOWER("JAYZ")返回"jayz"`
}
