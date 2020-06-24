package array_test

import (
	"testing"
)

func TestArrayInit(t *testing.T) {
	var arr [3]int
	arr1 := [3]int{1, 2, 3}
	//当你不想去查有多少位的时候,用[...]即可
	arr2 := [...]int{3, 2, 1}
	t.Log(arr[0], arr[1])
	t.Log(arr1, arr2)
}

func TestArrayTravel(t *testing.T) {
	array := [...]int{1, 2, 3, 4, 5}
	array[0] = 10
	//利用传统的方式,用数组下标进行便利
	for i := 0; i < len(array); i++ {
		t.Log(array[i])
	}

	//用for的方式来进行便利,利用range函数
	for inx, v := range array {
		t.Log(inx, v)
	}

	//如果有变量被声明但没有使用,在编译的时候会有问题,所以需要用_来表示忽略该变量
	for _, v := range array {
		t.Log(v)
	}
}

func TestArraySection(t *testing.T) {

}
