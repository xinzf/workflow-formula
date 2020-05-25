package formula

import (
	"errors"
	"github.com/Knetic/govaluate"
	"github.com/xinzf/workflow-formula/char"
	"github.com/xinzf/workflow-formula/logic"
	"github.com/xinzf/workflow-formula/math"
)

type Function interface {
	GetName() string
	GetFunc() govaluate.ExpressionFunction
	SetValues(values map[string]interface{})
	GetCategory() string
	GetDescription() string
}

type Out struct {
	Name string              `json:"name"`
	Funs []map[string]string `json:"funs"`
}

func NewCalculator(values map[string]interface{}) *Calculator {
	c := &Calculator{
		values: values,
		functions: []Function{
			// 文本函数
			new(char.Concat), new(char.Exact), new(char.Left), new(char.Len), new(char.Len),
			new(char.Lower), new(char.Replace), new(char.Rept), new(char.Right),
			new(char.Search), new(char.Split), new(char.Trim), new(char.Upper),
			// 逻辑函数
			new(logic.And), new(logic.False), new(logic.If), new(logic.Ifs), new(logic.Not),
			new(logic.Or), new(logic.True), new(logic.Match),
			// 数学函数
			new(math.Average), new(math.Count), new(math.Max), new(math.Min), new(math.Abs),
			new(math.Round), new(math.Int), new(math.Float), new(math.Mod), new(math.Product), new(math.Sum),
			new(math.Sumproduct), new(math.Rand), new(math.Power),
		},
	}
	return c
}

type Calculator struct {
	values    map[string]interface{}
	functions []Function
}

func (this *Calculator) build() map[string]govaluate.ExpressionFunction {
	funcs := make(map[string]govaluate.ExpressionFunction)
	for _, f := range this.functions {
		f.SetValues(this.values)
		funcs[f.GetName()] = f.GetFunc()
	}
	return funcs
}

func (this *Calculator) Functions() []Out {
	outs := make([]Out, 0)
	indexs := make(map[string]int, 0)
	for _, f := range this.functions {
		index, flag := indexs[f.GetCategory()]
		if !flag {
			outs = append(outs, Out{
				Name: f.GetCategory(),
				Funs: []map[string]string{},
			})
			indexs[f.GetCategory()] = len(outs) - 1
			index = len(outs) - 1
		}
		outs[index].Funs = append(outs[index].Funs, map[string]string{
			"name":        f.GetName(),
			"description": f.GetDescription(),
		})
	}

	return outs
}

func (this *Calculator) EvalBool(expression string) (bool, error) {
	val, err := this.Eval(expression)
	if err != nil {
		return false, err
	}

	v, flag := val.(bool)
	if !flag {
		return false, errors.New("返回值不是有效 Boolean 数据")
	}
	return v, nil
}

func (this *Calculator) Eval(expression string) (val interface{}, err error) {
	defer func() {
		if er := recover(); er != nil {
			switch er.(type) {
			case error:
				err = er.(error)
			}
		}
	}()

	funcs := this.build()

	var exp *govaluate.EvaluableExpression
	exp, err = govaluate.NewEvaluableExpressionWithFunctions(expression, funcs)
	if err != nil {
		return nil, err
	}

	var data interface{}
	data, err = exp.Evaluate(this.values)
	if err != nil {
		return nil, err
	}

	return data, nil
}
