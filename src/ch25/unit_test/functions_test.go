package testing

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSquareAssert(t *testing.T) {
	inputs := [...]int{1, 2, 3}
	expectds := [...]int{1, 4, 9}
	for i := 0; i < len(inputs); i++ {
		ret := square(i + 1)
		//在引用了断言包之后可以直接使用断言,可以看到效果其实是一样的
		assert.Equal(t, expectds[i], ret)
		/* if ret != expectds[i] {
			t.Errorf("input is %d, the expect is %d, the actual is %d", inputs[i], expectds[i], ret)
		} */
	}
}

//可以看到,end照常输出,说明不会终端程序
func TestErrorInCode(t *testing.T) {
	fmt.Println("start")
	t.Error("error")
	fmt.Println("end")
}

//可以看到,输出中不见后面的end,被中断
func TestFatalInCode(t *testing.T) {
	fmt.Println("start")
	t.Fatal("fatal")
	fmt.Println("end")
}
