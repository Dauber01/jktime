package groutine

import (
	"fmt"
	"testing"
	"time"
)

func TestGroutine(t *testing.T) {
	//可以看到,是通过不同的协程进行处理的
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)
	}
	//当这么写的时候,会发现输出的都是10,这是因为i是共享的,而将其传入函数的话,go的传递方式为
	//值传递,所以对应每个协程中的函数的参数地址的指针均不一致,所以不会发生该问题
	/* for i := 0; i < 10; i++ {
		go func() {
			fmt.Println(i)
		}()
	} */
	time.Sleep(time.Millisecond * 50)
}
