package _map

import "testing"

func TestMapWithFunValue(t *testing.T) {
	m := map[int]func(op int) int{}
	//可以看出,对于函数值为map的结构
	m[1] = func(op int) int {
		return op
	}
	m[2] = func(op int) int {
		return op * op
	}
	m[3] = func(op int) int {
		return op * op * op
	}
	t.Log(m[1](2), m[2](2), m[3](2))
}

func TestMapForSet(t *testing.T) {
	//可以看出,用value类型为boolean的map来进行存储,可以实现set的所有操作
	mySet := map[int]bool{}
	mySet[1] = true
	n := 1
	if mySet[n] {
		t.Logf("%d is exiting", n)
	} else {
		t.Logf("%d is not exit", n)
	}
	mySet[2] = true
	t.Log(len(mySet))
	delete(mySet, 1)
	if mySet[n] {
		t.Logf("%d is exiting", n)
	} else {
		t.Logf("%d is not exit", n)
	}
}
