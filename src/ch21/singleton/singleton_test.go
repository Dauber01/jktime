package singleton

import (
	"fmt"
	"sync"
	"testing"
	"unsafe"
)

type Singleton struct {
}

var singleInstance *Singleton
var once sync.Once

func getSingletonObj() *Singleton {
	//表示不管在多个协程中,都只执行一次,利用他可以实现单例模式,同时在其中声明的变量
	//为其内部的局部变量,所以要在外部时候的话,要在函数前面进行声明
	once.Do(func() {
		fmt.Println("obj create")
		singleInstance = new(Singleton)
	})
	return singleInstance
}

func TestOnceInstance(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		go func() {
			//为了防止主协程在各个协程执行完毕之前就已经退出,采用waitgroup机制
			wg.Add(1)
			singlePor := getSingletonObj()
			//可以看到获取对象的地址也均为一个地址
			fmt.Printf("%x\n", unsafe.Pointer(singlePor))
			wg.Done()
		}()
	}
	wg.Wait()
}
