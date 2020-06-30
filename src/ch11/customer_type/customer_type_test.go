package customer_type

import (
	"fmt"
	"time"
)

//可以通过抽象类型大大的简化代码的复杂程度
type IntConv func(op int) int

func timeSpent(inner IntConv) IntConv {
	return func(op int) int {
		start := time.Now()
		ret := inner(op)
		fmt.Print("time spent:", time.Since(start).Seconds())
		return ret
	}
}
