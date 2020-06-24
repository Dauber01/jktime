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

const (
	Readable = 1 << iota
	Writable
	Executable
)

func TestConstantTry(t *testing.T) {
	t.Log(Monday, Tuesday)
}

func TestConstantTry1(t *testing.T) {
	a := 7 //0111
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)
}
