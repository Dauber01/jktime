package type_test

import "testing"

//自定义类型的方式
type MyInt int64

func TestImplicit(t *testing.T) {
	//var a int = 1
	var a int32 = 1
	//之后将变量a定义为int32,这时候进行隐式转换仍然不能成功
	var b int64
	//因为当前我的mac os系统是64位架构,所以a其实默认为64位,但仍无法完成类型的转换
	b = int64(a)
	var c MyInt
	//这时候我们可以看到,虽然自定义类型MyInt其实就是int64,但是也不能互相转换
	c = MyInt(b)

	t.Log(a, b, c)
}

func TestPoint(t *testing.T) {
	a := 1
	//获取指针的时候和c是一样的直接用&获取指针
	aPtr := &a
	//可以看到指针不支持运算,即不能进行对应的数组用下标直接取数据
	//aPtr = aPtr + 1
	t.Log(a, aPtr)
	//可以通过%T打印变量的类型
	t.Logf("%T, %T", a, aPtr)
}

func TestString(t *testing.T) {
	//因为string在go中是属于值类型,所以它未初始化的时候的默认值是"",而不是nil
	var s string
	t.Log("*" + s + "*")
	t.Log(len(s))
}
