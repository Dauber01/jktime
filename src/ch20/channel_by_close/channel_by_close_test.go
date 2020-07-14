package channel_by_close

import (
	"fmt"
	"testing"
	"time"
)

func isChannelled(cancelChan chan struct{}) bool {
	select {
	case <-cancelChan:
		return true
	default:
		return false
	}
}

//在该func中可以看到向cancel中返回了一个空的结构体
func cancel_1(cancelChan chan struct{}) {
	cancelChan <- struct{}{}
}

func cancel_2(cancelChan chan struct{}) {
	close(cancelChan)
}

func TestChannel(t *testing.T) {
	cancelChan := make(chan struct{}, 0)
	for i := 0; i < 5; i++ {
		go func(i int, cannelCh chan struct{}) {
			for {
				if isChannelled(cannelCh) {
					break
				}
			}
			time.Sleep(time.Millisecond * 5)
			fmt.Println(i, "Done")
		}(i, cancelChan)

	}
	//可以看到,当使用关闭channel的方法的时候,由于是广播式的,所以协程中的对应的方法都会收到通知,从而关闭
	//所有的协程
	cancel_2(cancelChan)
	//当调用cancel的时候发现向channel中存入数据,只有一个协程会收到信号,但是发送对应协程数量的信号,又会使项目
	//丧失弹性拓展的能力
	//cancel_1(cancelChan)
	//睡眠的时间一定要大于协程睡眠的时间,否则将会有协程没执行完毕
	time.Sleep(time.Second * 1)
}
