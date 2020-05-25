package char

import "github.com/Knetic/govaluate"

type Rept struct {
	values map[string]interface{}
}

func (this *Rept) GetName() string {
	return "REPT"
}

func (this *Rept) GetFunc() govaluate.ExpressionFunction {
	return func(arguments ...interface{}) (interface{}, error) {
		return nil, nil
	}
}
func (this *Rept) SetValues(values map[string]interface{}) {
	this.values = values
}

func (this *Rept) GetCategory() string {
	return "文本函数"
}

func (this *Rept) GetDescription() string {
	return `REPT函数可以将文本重复一定次数
用法：REPT(文本,重复次数)
示例：REPT("居润",3)返回"居润居润居润"`
}
