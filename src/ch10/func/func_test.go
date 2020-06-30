package _func

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func returnMultiValues() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}

//该函数用来嵌套函数,增加计时的部分
func timeSpent(inner func(op int) int) func(op int) int {
	return func(op int) int {
		start := time.Now()
		ret := inner(op)
		fmt.Print("time spent:", time.Since(start).Seconds())
		return ret
	}
}

func slowFun(op int) int {
	time.Sleep(time.Second * 1)
	return op
}

func TestFn(t *testing.T)  {
	//可以看到,函数可以有多个返回值
	/*a, b := returnMultiValues();
	t.Log(a , b)*/
	//如果想忽略其中的一个返回值的时候,可以用_进行代替该返回值
	a, _ := returnMultiValues();
	t.Log(a)

	//这就是函数式编程
	fun := timeSpent(slowFun)
	t.Log(fun(10))
}

func multipilyFun(ops ...int) int  {
	count := 0
	for _, v := range ops {
		count += v
	}
	return count
}

func TestFun(t *testing.T)  {
	c1 := multipilyFun(1, 2, 3, 4)
	c2 := multipilyFun(1, 2, 3, 4, 5)
	t.Log(c1, c2)
}

func clear()  {
	fmt.Println("clear resources.")
}

func TestDeferFun(t *testing.T)  {
	defer clear()
	t.Log("start")
	//panic函数表示会出现错误
	panic("err")
	//可以看到在异常之后的语句不会执行,但是defer的可以执行
	t.Log("end")
}