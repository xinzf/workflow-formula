package formula

import (
	"testing"
)

func init() {
}

func TestLogic(t *testing.T) {
	//str := `OR(field1>90,AND(field2>100,field3=='xxxx'))`
	formula := NewCalculator(map[string]interface{}{
		"field1": []string{"bbbc", "asaa", "111"},
		"field2": 801,
		"field3": "xxxx",
	})

	str := `MATCH(field1,"aaa","bbb","fsdfd",111)`
	ret, err := formula.Eval(str)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%+v\n", ret)
}
