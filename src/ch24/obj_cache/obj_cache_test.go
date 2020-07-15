package obj_cache

import (
	"fmt"
	"sync"
	"testing"
)

func TestSyncPool(t *testing.T) {
	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Create a new object.")
			return 100
		},
	}
	v := pool.Get().(int)
	fmt.Println(v)
	pool.Put(3)
	//可以在代码中立刻发起一次gc,一般不会有人使用,可以看到gc之后,又重新创建了对象,证明了生命周期
	//runtime.GC()
	v1, _ := pool.Get().(int)
	fmt.Println(v1)
	//可以看到存入的对象在被拿出来但是并没有重新put的情况下,pool中就不存在对象了
	v2, _ := pool.Get().(int)
	fmt.Println(v2)
}

func TestSyncPoolEvey(t *testing.T) {
	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Create a new object.")
			return 10
		},
	}
	pool.Put(100)
	pool.Put(100)
	pool.Put(100)
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		go func() {
			wg.Add(1)
			//可以看到存入的3个100被取出来之后就都走默认的新创建流程
			v, _ := pool.Get().(int)
			fmt.Println(v)
			wg.Done()
		}()
	}
	wg.Wait()

}
