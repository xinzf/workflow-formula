package char

import (
	"errors"
	"github.com/Knetic/govaluate"
	"unicode/utf8"
)

type Len struct {
	values map[string]interface{}
}

func (this *Len) GetName() string {
	return "LEN"
}

func (this *Len) GetFunc() govaluate.ExpressionFunction {
	return func(arguments ...interface{}) (interface{}, error) {
		if len(arguments) != 1 {
			return nil, errors.New("LEN: 参数数量不对")
		}
		val, _ := arguments[0].(string)
		str := utf8.RuneCountInString(val)
		return str, nil
	}
}
func (this *Len) SetValues(values map[string]interface{}) {
	this.values = values
}

func (this *Len) GetCategory() string {
	return "文本函数"
}

func (this *Len) GetDescription() string {
	return `LEN函数可以获取文本中的字符个数
用法：LEN(文本)
示例：LEN("朝辞白帝彩云间")返回7，因为这句诗中有7个字符`
}
