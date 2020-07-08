package panic_recover

import (
	"errors"
	"fmt"
	"testing"
)

func TestPanicVxExit(t *testing.T) {
	//可以看到通过defer和recover可以抓捕异常并进行处理,组织程序的终止
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recover from", err)
		}
	}()

	//可以看出虽然有异常,但是还是会打出defer中的内容
	/* defer func() {
		fmt.Println("Finally!")
	}() */
	fmt.Println("start")
	//当用panic退出的时候,会有调用栈和其他的信息
	panic(errors.New("something wrong!"))
	//可以看出利用exit函数退出的没有调用栈信息和其他信息
	//os.Exit(-1)
}
