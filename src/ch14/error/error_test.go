package error

import (
	"errors"
	"testing"
)

//可以在开始的时候声明异常的类型
var NotInRangeError error = errors.New("n should be in [2, 1000]")
var NotBeFiveError error = errors.New("n should not be 5")

//可以看到,可以在返回值的时候增加(),存入多返回值,来对异常进行返回
func GetFibonacci(n int) ([]int, error) {
	if n < 2 || n > 100 {
		//创建异常的方法如下所示
		return nil, NotInRangeError
	}
	if n == 5 {
		return nil, NotBeFiveError
	}
	fibList := []int{1, 1}
	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}
	return fibList, nil
}

func TestGetFibonacci(t *testing.T) {
	//调用的时候,先进行判读返回值是否有异常,之后进行操作
	if v, error := GetFibonacci(5); error != nil {
		//通过判断异常的类型进行不同的操作
		if NotInRangeError == error {
			t.Log("NotInRangeError")
		} else if NotBeFiveError == error {
			t.Log("NotBeFiveError")
		}
	} else {
		t.Log(v)
	}
}
