package unsafe_programming

import (
	"time"
	"sync"
	"runtime/internal/atomic"
	"testing"
	"unsafe"
)

func TestUnsafe(t *testing.T) {
	i := 10
	pot := *(*float64)(unsafe.Pointer(&i))
	//可以看到输出值的类型是完全不一样的,通常我们不能用这种方式进行类型转换
	t.Log(unsafe.Pointer(&i))
	t.Log(pot)
}

type MyInt int

func TestAliseUnsafe(t *testing.T) {
	i := []int{1, 2, 3}
	//可以看到,该种不安全的转换可以用作别名的指针类型转换
	pot := *(*[]MyInt)(unsafe.Pointer(&i))
	t.Log(pot)
}

func TestAtomic(t *testing.T) {
	var shareBUfferPtr unsafe.Pointer
	writeDataFn := func() {
		for i := 0; i < 100; i++ {
			data = append(data, i)
		}
		atomic.StorePointer(&shareBUfferPtr, unsafe.Pointer(&data))
	}
	readDataFn := func() {
		data := atomic.LoadPointer(&shareBUfferPtr)
		fmt.Println(data, *(*int[])data)
	}
	var wg sync.WaitGroup
	writeDataFn()
	for i := 0; i < 10;i++{
		wg.Add(1)
		go func(){
			for i := 0; i < 10; i++{
				writeDataFn()
				time.Sleep(time.Millisecond * 100)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
