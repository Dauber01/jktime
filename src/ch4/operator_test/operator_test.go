package operator_test

import "testing"

func TestCompareArray(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 2, 4, 3}
	//c := [...]int{1, 2, 3, 4, 5}
	d := [...]int{1, 2, 3, 4}
	t.Log(a == b)
	//可以看到,因为a和c的数组长度不同,所以不允许进行比较
	//t.Log(a == c)
	//相等要求位置和值都相等
	t.Log(a == d)
}
