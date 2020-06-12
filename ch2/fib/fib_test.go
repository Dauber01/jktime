package fib

import (
	"fmt"
	"testing"
)

//var的声明一般用于全局的/静态的变量声明的时候
var a int

func TestFib(t *testing.T) {

	//在赋值变量的时候可以指定初始类型
	/*var a int = 1
	var b int = 1*/
	//一般声明的时候允许用声明块的方式一起声明多个变量
	var (
		//a int = 1
		//在使用默认的推断类型的时候,可以不指定类型
		b = 1
	)
	//可以不用声明的类型,而利用:=对变量进行推断
	//a:=1
	a = 1
	//一般在测试的单元中没有必要使用fmt,而使用t.log()函数就好
	t.Log(a)
	for i := 0; i < 5; i++ {
		fmt.Print(" ", b)
		tmp := a
		a = b
		b = tmp + a
	}
	fmt.Println()

}

func TestExchange(t *testing.T) {
	a := 1
	b := 2
	/*tmp := a
	a = b
	b = tmp*/
	//交换赋值很简洁,可以通过一个语句进行交换赋值
	a, b = b, a
	t.Log(a, b)

}
