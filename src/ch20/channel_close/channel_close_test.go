package channel_close

import (
	"fmt"
	"sync"
	"testing"
)

//缓存数据的生产者
func dataProducer(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		//关闭ch,同时关闭所有消费者
		close(ch)
		//可以看到,再继续向里存会出现panic
		ch <- 100
		wg.Done()
	}()
}

//缓存数据的接受者
func dataReceiver(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for {
			//管道有两个返回值,其中当管道被关闭的时候,第二个返回值为false,否则为true
			if data, ok := <-ch; ok {
				fmt.Println(data)
			} else {
				break
			}

		}
		wg.Done()
	}()
}

func TestCloseChannel(t *testing.T) {
	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(1)
	dataProducer(ch, &wg)
	wg.Add(1)
	dataReceiver(ch, &wg)
	wg.Wait()
}
