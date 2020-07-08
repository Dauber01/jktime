package share_mem

import (
	"sync"
	"testing"
	"time"
)

func TestCounter(t *testing.T) {
	counter := 0
	//可以看到输出的计数小于5000
	for i := 0; i < 5000; i++ {
		go func() {
			counter++
		}()
	}
	time.Sleep(1 * time.Second)
	t.Logf("counter = %d", counter)
}

func TestCounterThreadSafe(t *testing.T) {
	var mut sync.Mutex
	var wg sync.WaitGroup
	counter := 0

	for i := 0; i < 5000; i++ {
		//要等待的任务加1
		wg.Add(1)
		go func() {
			//异常处理,避免死锁的情况
			defer func() {
				mut.Unlock()
			}()
			//可以看到,在加锁之后为线程安全的,计数为5000
			mut.Lock()
			counter++
			//要等待的任务结束一个
			wg.Done()
		}()
	}
	//可以等到上面的结束
	wg.Wait()
	//一直进行线程睡眠是防止外面的协程执行速度小于里面协程的执行速度,造成有一些协程没有执行完的问题
	//time.Sleep(1 * time.Second)
	t.Logf("counter = %d", counter)
}
