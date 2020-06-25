package _map

import (
	"testing"
)

func TestMapInit(t *testing.T) {

	m := map[int]int{1: 1, 2: 2, 3: 3}
	t.Log(m[1])
	t.Logf("m len %d", len(m))
	m1 := map[int]int{}
	m1[4] = 16
	t.Logf("m1 len %d", len(m1))
	m2 := make(map[int]int, 10)
	t.Logf("m2 len %d", len(m2))
	//可以看出cap()函数不能用于map结构
	//t.Log(cap(m2))
}

func TestMapExit(t *testing.T) {
	m := map[int]int{}
	m[2] = 0
	//因为返回值都是初始值0,所以无法区分哪个元素存在
	t.Log(m[1], m[2])
	//可以通过第二个返回值来判断是否存在,存在true,不存在false;
	if v, ok := m[2]; ok {
		t.Log("2 exiting v = ", v)
	} else {
		t.Log("2 is not exit")
	}
}

func TestMapRange(t *testing.T) {
	m := map[string]int{"one": 1, "two": 2, "three": 3}
	for k, v := range m {
		t.Log(k, v)
	}
}
