package series

import (
	"fmt"
)

//可以看到,可以在返回值的时候增加(),存入多返回值,来对异常进行返回
func GetFibonacci(n int) []int {
	fibList := []int{1, 1}
	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}
	return fibList
}

//用来查看小写字母开头的fun是否能被别的包访问
func sercase(i int) int {
	return i * i
}

//可以看到,如果该包的函数被调用两次或者调用两个函数,init函数仍然只会被执行一次
func init() {
	fmt.Println("init1")
}

//可以看到,在一个包/源文件中可以有多个init函数,且可以同时被执行
func init() {
	fmt.Println("init2")
}
