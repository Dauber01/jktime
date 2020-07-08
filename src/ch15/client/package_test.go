package client

import (
	"ch15/series"
	"testing"
)

func TestPackage(t *testing.T) {
	//可以看到在配好对应的环境变量之后,可以访问其他包大写字母开头的函数
	t.Log(series.GetFibonacci(5))
	//可以看到,小写字母开头的fun在其他包无法被访问
	//t.Log(series.sercase(3))
}
