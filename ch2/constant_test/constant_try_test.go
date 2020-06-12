package constant_test

import "testing"

//在go语言中对于连续的常量一般会选择用 iota + 1 进行赋值
const (
	Monday = iota + 1
	Tuesday
	Wednesday
)

//在go语言中也可以像其他语言那样进行逐一的赋值
/*const (
	Monday = 1
	Tuesday = 2
	Wednesday = 3
)*/

func TestConstantTry(t *testing.T) {
	t.Log(Monday, Tuesday)
}
