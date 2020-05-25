package char

import "github.com/Knetic/govaluate"

type Split struct {
	values map[string]interface{}
}

func (this *Split) GetName() string {
	return "SPLIT"
}

func (this *Split) GetFunc() govaluate.ExpressionFunction {
	return func(arguments ...interface{}) (interface{}, error) {
		return nil, nil
	}
}
func (this *Split) SetValues(values map[string]interface{}) {
	this.values = values
}

func (this *Split) GetCategory() string {
	return "文本函数"
}

func (this *Split) GetDescription() string {
	return `SPLIT函数可以将文本按指定分割符分割成数组
用法：SPLIT(文本,分隔符_文本)
示例：SPLIT("简道云-应用定制工具","-")返回"简道云，应用定制工具"`
}
