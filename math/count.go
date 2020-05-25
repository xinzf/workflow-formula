package math

import (
	"github.com/Knetic/govaluate"
)

type Count struct {
	values map[string]interface{}
}

func (this *Count) GetName() string {
	return "COUNT"
}

func (*Count) GetCategory() string {
	return "数学函数"
}

func (*Count) GetDescription() string {
	return `COUNT(value1, [value2], …)
统计参数个数。`
}

func (this *Count) SetValues(values map[string]interface{}) {
	this.values = values
}

func (this *Count) GetFunc() govaluate.ExpressionFunction {
	return func(arguments ...interface{}) (interface{}, error) {
		//if len(arguments) == 0 {
		//	return nil, errors.New("COUNT: 参数数量不足")
		//}

		//var total int
		//for _,arg:=range arguments{
		//val,flag:=this.values[arg]
		//if !flag {
		//	logrus.Errorf("COUNT: 没有找到 %s 指定的数据",arg)
		//	continue
		//}
		//}
		return len(arguments), nil
	}
}
