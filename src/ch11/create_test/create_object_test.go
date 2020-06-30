package create_test

import (
	"fmt"
	"testing"
	"unsafe"
)

type Employee struct {
	Id   string
	Name string
	Age  int
}

func (e Employee) String() string {
	fmt.Printf("Address is %x\n", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("ID:%s/Name:%s/Age:%d", e.Id, e.Name, e.Age)
}

//当函数定义为指针的时候,仍然可以通过结构体进行访问
//func (e *Employee) String() string {
//通过这个方法获取name字段的地址
//	fmt.Printf("Address is %x\n", unsafe.Pointer(&e.Name))
//	return fmt.Sprintf("ID:%s/Name:%s/Age:%d", e.Id, e.Name, e.Age)
//}

func TestCreateEmployeeObj(t *testing.T) {
	e := Employee{"0", "Bob", 20}
	e1 := Employee{Name: "Mike", Age: 30}
	e2 := new(Employee) //返回指针
	e2.Id = "2"
	e2.Age = 22
	e2.Name = "Rose"
	t.Log(e)
	t.Log(e1)
	//可以看到,当id没有被初始化的时候,打出的值为string类型的初始值
	t.Log(e1.Id)
	t.Log(e2)
	t.Logf("e is %T", e)
	//打出的类型为*create_test.Employee,即Employee的指针类型,与&e的类型相同
	t.Logf("e2 is %T", e2)
}

func TestMethodInStruct(t *testing.T) {
	//我们尝试下可以发现,在如参为结构体的函数中,如果将入参传入指针,仍然可以生效
	e := &Employee{"9", "lilei", 18}
	//在调用函数中同样打出地址,尝试之后发现,如果直接传结构体的话,地址是不一样的,说明发生了复制的过程
	//有额外的开销,所以我们在写函数的时候尽量写后一种方式
	fmt.Printf("Address is %x\n", unsafe.Pointer(&e.Name))
	//说明在函数中,如果是指针的话,访问并不需要通过->来进行,因为其自己会进行判断，从而确定访问的方法
	t.Log(e.String())
}
