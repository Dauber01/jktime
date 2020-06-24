package condition

import (
	"testing"
)

func TestIfMultiSec(t *testing.T) {
	if a := 1 == 1; a {
		t.Log("1==1")
	}
	//双返回值主要是为了进行如下的处理
	/*if v,err := someFun(); err==nil {
		t.Log("1==1")
	}else {
		t.Log("1==1")
	}*/
}

func TestSwichMultiCase(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch i {
		//可以看到switch后面支持多个case的选项
		case 0, 2:
			t.Log("Even")
			//可以看到,在不使用break的时候,仍然可以自动跳出
		case 1, 3:
			t.Log("odd")
		default:
			t.Log("it is not 0-3")
		}
	}
}

func TestSwitchCaseCondition(t *testing.T) {
	for i := 0; i < 5; i++ {
		//这个switch的函数整体上看起来更像if else的规则
		switch {
		case i%2 == 0:
			t.Log("Even")
		case i%2 == 1:
			t.Log("Odd")
		default:
			t.Log("Unkonw")
		}
	}
}
