package char

import "github.com/Knetic/govaluate"

type Upper struct {
	values map[string]interface{}
}

func (this *Upper) GetName() string {
	return "UPPER"
}

func (this *Upper) GetFunc() govaluate.ExpressionFunction {
	return func(arguments ...interface{}) (interface{}, error) {
		return nil, nil
	}
}
func (this *Upper) SetValues(values map[string]interface{}) {
	this.values = values
}

func (this *Upper) GetCategory() string {
	return "文本函数"
}

func (this *Upper) GetDescription() string {
	return `UPPER函数可以将一个文本中的所有小写字母转换为大写字母
用法：UPPER(文本)
示例：UPPER("jayz")返回"JAYZ"`
}
