package impl_test

import (
	"fmt"
	"testing"
)

type Code string

type Programmer interface {
	WriteHelloWorld() Code
}

type GoProgrammer struct {
}

func (p *GoProgrammer) WriteHelloWorld() Code {
	return "fmt.Println(\"Hello Word!\")"
}

type JavaProgrammer struct {
}

func (p *JavaProgrammer) WriteHelloWorld() Code {
	return "System.out.Println(\"Hello Word!\")"
}

//%T经常使用为输出对应的类型
func writeFirstProgrammer(p Programmer) {
	fmt.Printf("%T %v\n", p, p.WriteHelloWorld())
}

func TestPolymorphism(t *testing.T) {
	//用下面这种方式声明的时候,返回的是类型,而在该函数中,需要输入的为指针,所以需要用&
	//因为Programmer是一个interface类型,只能是一个指针类型
	goProg := &GoProgrammer{}
	javaProg := new(JavaProgrammer)
	writeFirstProgrammer(goProg)
	writeFirstProgrammer(javaProg)
}
