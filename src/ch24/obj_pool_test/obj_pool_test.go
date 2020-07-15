package obj_pool_test

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

//因为仅仅是用来做演示所以选用了空结构体,也可以使用其他的对象进行填充
type ReusableObj struct {
}

type ObjPool struct {
	bufChan chan *ReusableObj
}

//和其他语言的对象池一样,需要用一个初始化的方法,并且该示例选择了在开始的时候初始化全部的内部对象
//而有的语言则是在最开始使用内部对象的时候才开始进行创建的
func NewObjPool(numOfObj int) *ObjPool {
	objPool := ObjPool{}
	objPool.bufChan = make(chan *ReusableObj, numOfObj)
	for i := 0; i < numOfObj; i++ {
		objPool.bufChan <- &ReusableObj{}
	}
	return &objPool
}

//获取对象
func (p *ObjPool) GetObj(timeout time.Duration) (*ReusableObj, error) {
	select {
	case ret := <-p.bufChan:
		return ret, nil
	case <-time.After(timeout):
		return nil, errors.New("time out")
	}
}

//重新存入的函数,注意当其中的channel已经满了的时候,利用select立刻返回失败
func (p *ObjPool) ReleaseObj(obj *ReusableObj) error {
	select {
	case p.bufChan <- obj:
		return nil
	default:
		return errors.New("over flow")
	}
}

func TestObjPool(t *testing.T) {
	pool := NewObjPool(10)
	//可以看到当pool已经满了的情况下在往其中存会显示我们设置的错误
	if err := pool.ReleaseObj(&ReusableObj{}); err != nil {
		t.Error(err)
	}
	for i := 0; i < 11; i++ {
		if v, error := pool.GetObj(time.Second * 1); error != nil {
			t.Error(error)
		} else {
			fmt.Printf("%T\n", v)
			//可以看到在只往外取不归还的时候,pool为空的时候会返回一个超时错误
			/* if err := pool.ReleaseObj(v); err != nil {
				t.Error(err)
			} */
		}
	}
	fmt.Println("Done")
}

//如果想要这个线程池可以存各种类型的对象,则将所有的对象类型改为interface{}即可,但是在每次取出的时候都需要判断
//对象的类型,所以不建议存不同类型的对象
