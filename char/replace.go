package char

import "github.com/Knetic/govaluate"

type Replace struct {
	values map[string]interface{}
}

func (this *Replace) GetName() string {
	return "REPLACE"
}

func (this *Replace) GetFunc() govaluate.ExpressionFunction {
	return func(arguments ...interface{}) (interface{}, error) {
		return nil, nil
	}
}
func (this *Replace) SetValues(values map[string]interface{}) {
	this.values = values
}

func (this *Replace) GetCategory() string {
	return "文本函数"
}

func (this *Replace) GetDescription() string {
	return `REPLACE函数可以根据指定的字符数，将部分文本替换为不同的文本
用法：REPLACE(文本,开始位置,替换长度,新文本)
示例：REPLACE("简道云应用定制工具",4,6,"企业数据管理平台")返回"简道云企业数据管理平台"`
}
