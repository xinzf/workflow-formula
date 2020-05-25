package char

import (
	"errors"
	"fmt"
	"github.com/Knetic/govaluate"
	"github.com/cstockton/go-conv"
)

type Left struct {
	values map[string]interface{}
}

func (this *Left) GetName() string {
	return "LEFT"
}

func (this *Left) GetFunc() govaluate.ExpressionFunction {
	return func(arguments ...interface{}) (interface{}, error) {
		if len(arguments) != 2 {
			return nil, errors.New("LEFT: 参数数量不正确")
		}

		if _, flag := arguments[0].(string); !flag {
			return nil, fmt.Errorf("LEFT: 第 1 个参数不是有效 String 类型")
		}

		pos, _ := conv.Int(arguments[1])
		if pos == 0 {
			return nil, fmt.Errorf("LEFT: 第 2 个参数不合法")
		}

		rs := []rune(arguments[0].(string))
		return string(rs[:pos]), nil
	}
}
func (this *Left) SetValues(values map[string]interface{}) {
	this.values = values
}

func (this *Left) GetCategory() string {
	return "文本函数"
}

func (this *Left) GetDescription() string {
	return `LEFT函数可以从一个文本的第一个字符开始返回指定个数的字符
用法：LEFT(文本,文本长度)
示例：LEFT("三年二班周杰伦",2)返回"三年"，也就是"三年二班周杰伦"的从左往右的前2个字符`
}
