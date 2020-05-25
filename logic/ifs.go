package logic

import (
	"errors"
	"fmt"
	"github.com/Knetic/govaluate"
)

type Ifs struct {
	values map[string]interface{}
}

func (this *Ifs) GetName() string {
	return "IFS"
}

func (*Ifs) GetCategory() string {
	return "逻辑函数"
}

func (*Ifs) GetDescription() string {
	return `IFS函数检查是否满足一个或多个条件，且返回符合第一个TRUE条件的值，IFS可以取代多个嵌套IF语句。
用法：IFS(逻辑表达式1,逻辑表达式1为true返回该值,逻辑表达式2,逻辑表达式2为true返回该值,...)
示例：IFS(语文成绩>90,"优秀",语文成绩>80,"良好",语文成绩>=60,"及格",语文成绩<60,"不及格")，根据成绩返回对应的评价。`
}

func (this *Ifs) SetValues(values map[string]interface{}) {
	this.values = values
}

func (this *Ifs) GetFunc() govaluate.ExpressionFunction {
	return func(arguments ...interface{}) (interface{}, error) {
		if len(arguments)%2 != 0 {
			return nil, errors.New("IFS: 参数必须是偶数集合")
		}

		for i := 0; i < len(arguments); i = i + 2 {
			val, flag := arguments[i].(bool)
			if !flag {
				return nil, fmt.Errorf("IFS: 第 %d 个参数不是有效的 Boolean 数据", i)
			}
			if val == true {
				return arguments[i+1], nil
			}
		}
		return nil, nil
	}
}
