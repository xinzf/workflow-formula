package char

import (
	"errors"
	"github.com/Knetic/govaluate"
)

// 使用场景：自动生成指定格式的字符串
type Concat struct {
	values map[string]interface{}
}

func (this *Concat) GetName() string {
	return "CONCAT"
}

func (this *Concat) GetFunc() govaluate.ExpressionFunction {
	return func(arguments ...interface{}) (interface{}, error) {
		if len(arguments) == 0 {
			return nil, errors.New("CONCAT: 参数数量不足")
		}

		str := ""
		for i := 0; i < len(arguments); i++ {
			val, _ := arguments[i].(string)
			str += val
		}
		return str, nil
	}
}
func (this *Concat) SetValues(values map[string]interface{}) {
	this.values = values
}

func (this *Concat) GetCategory() string {
	return "文本函数"
}

func (this *Concat) GetDescription() string {
	return `CONCAT函数可以将多个文本合并成一个文本
用法：CONCAT(文本1,文本2,...)
示例：CONCAT("三年二班","周杰伦")会返回"三年二班周杰伦"`
}
