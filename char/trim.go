package char

import "github.com/Knetic/govaluate"

type Trim struct {
	values map[string]interface{}
}

func (this *Trim) GetName() string {
	return "TRIM"
}

func (this *Trim) GetFunc() govaluate.ExpressionFunction {
	return func(arguments ...interface{}) (interface{}, error) {
		return nil, nil
	}
}
func (this *Trim) SetValues(values map[string]interface{}) {
	this.values = values
}

func (this *Trim) GetCategory() string {
	return "文本函数"
}

func (this *Trim) GetDescription() string {
	return `TRIM函数可以删除文本首尾的空格
用法：TRIM(文本)
示例：TRIM(" 简道云 ")返回"简道云"`
}
