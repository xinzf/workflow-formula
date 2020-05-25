package char

import (
	"errors"
	"fmt"
	"github.com/Knetic/govaluate"
)

// 使用场景：比较两个字段的值是否相等，配合IF使用
type Exact struct {
	values map[string]interface{}
}

func (this *Exact) GetName() string {
	return "EXACT"
}

func (this *Exact) GetFunc() govaluate.ExpressionFunction {
	return func(arguments ...interface{}) (interface{}, error) {
		if len(arguments) != 2 {
			return nil, errors.New("EXACT: 参数数量不正确")
		}

		for i := 0; i < len(arguments); i++ {
			if _, flag := arguments[i].(string); !flag {
				return nil, fmt.Errorf("第 %d 个参数不是有效的 String 类型", i)
			}
		}

		if arguments[0].(string) == arguments[1].(string) {
			return true, nil
		}
		return false, nil
	}
}
func (this *Exact) SetValues(values map[string]interface{}) {
	this.values = values
}

func (this *Exact) GetCategory() string {
	return "文本函数"
}

func (this *Exact) GetDescription() string {
	return `EXACT函数可以比较两个文本是否完全相同，完全相同则返回true，否则返回false
用法：EXACT(文本1, 文本2)
示例：EXACT(手机号,中奖手机号)，如果两者相同，返回true，如果不相同，返回false`
}
