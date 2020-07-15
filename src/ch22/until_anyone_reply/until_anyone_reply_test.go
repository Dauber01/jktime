package until_anyone_reply

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func runTask(id int) string {
	time.Sleep(10 * time.Millisecond)
	return fmt.Sprintf("THe result is from %d", id)
}

func FirstResponse() string {
	numOfRunner := 10
	//可以看到在channel没有buffer的时候,在第一个协程计算的结果得以返回的时候,其余协程的计算结果在存入
	//channel之后,没有人去获取,在程序退出前打出的日志显示剩余协程的数量为11,在大型项目中很容易出现oom
	//ch := make(chan string)
	//在使用了buffer之后,生产者和消费者之间的关系为解藕,所以在其余的协程计算完毕之后,直接将自己的计算结果存入
	//channel之后就直接退出,并没有被阻塞住
	ch := make(chan string, numOfRunner)
	for i := 0; i < numOfRunner; i++ {
		go func(i int) {
			ret := runTask(i)
			ch <- ret
		}(i)
	}
	return <-ch
}

func TestFirstReponse(t *testing.T) {
	//该方法用来获取当前程序在系统中开启的协程数量,用来查看系统是否被阻塞住
	t.Log("brfore", runtime.NumGoroutine())
	t.Log(FirstResponse())
	time.Sleep(time.Second * 1)
	t.Log("after", runtime.NumGoroutine())
}
