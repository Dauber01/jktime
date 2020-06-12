package main

import (
	"fmt"
	"os"
)

func main() {

	//可以通过os.Args来获取参数的值
	if len(os.Args) > 1 {
		fmt.Println(os.Args[0])
		fmt.Println(os.Args[1])
	}

	//main函数是项目的启动函数,对应的package必须是main,但是上层目录的文件夹不一定命名为main
	fmt.Println("Hello world!")

	//main函数没有返回值一说，必须通过os函数进行返回，当os code不为0的时候,会显示程序异常结束,并附带异常码
	os.Exit(-1)

}
