package bdd_test

import (
	"testing"

	//在引用的前面有一个.表示引用到当前包里面,当我们在使用他们的时候,不需要写很长的方法调用前缀
	. "github.com/smartystreets/goconvey/convey"
)

func TestSpec(t *testing.T) {
	Convey("Given 2 even numbers", t, func() {
		a := 2
		b := 4

		Convey("when add the two numbers", func() {
			c := a + b

			Convey("Then the result is still even", func() {
				//该包自带断言,我们在该方法的左右两个参数各输入需要比较的两端,中间传入相应的关系和规则
				So(c%2, ShouldEqual, 0)
			})
		})
	})
}
