package logic

import (
	"errors"
	"github.com/Knetic/govaluate"
	"github.com/cstockton/go-conv"
	"reflect"
)

type Match struct {
	values map[string]interface{}
}

func (this *Match) GetName() string {
	return "MATCH"
}

func (*Match) GetCategory() string {
	return "逻辑函数"
}

func (*Match) GetDescription() string {
	return `用于匹配参数是数组的时候，是否符合被匹配数据相等
用法：MATCH(value,"要匹配的值","...")
示例：MATCH(李晨,"李**","陈**")`
}

func (this *Match) SetValues(values map[string]interface{}) {
	this.values = values
}

func (this *Match) GetFunc() govaluate.ExpressionFunction {
	return func(arguments ...interface{}) (interface{}, error) {
		if len(arguments) < 2 {
			return nil, errors.New("MATCH: 参数数量不足")
		}

		targets := make([]interface{}, 0)
		re := reflect.ValueOf(arguments[0])
		switch re.Kind() {
		case reflect.Slice:
			for i := 0; i < re.Len(); i++ {
				targets = append(targets, re.Index(i))
			}
		default:
			targets = append(targets, arguments[0])
		}

		finder := func(val interface{}) bool {
			vv, _ := conv.String(val)
			for _, v := range targets {
				vvv, _ := conv.String(v)
				if reflect.DeepEqual(vvv, vv) {
					return true
				}
			}
			return false
		}

		data := false
		for i := 1; i < len(arguments); i++ {
			if finder(arguments[i]) {
				data = true
				break
			}
		}
		return data, nil
	}
}
