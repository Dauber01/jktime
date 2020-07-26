package ch28

import (
	"reflect"
	"testing"
)

type Employee struct{
	Name string
	Age int
}

type Customer struct{
	Name string
	Age int
}

func TestFillNameAndAge(t *testing.T){
	settings := map[string]interface{}{"Name" : "Mike", "Age" : 40}
	e := Employee{}
	if err := fill
}

func TestDeepEqual(t *testing.T) {
	a := map[int]string{1: "one", 2: "two", 3: "three"}
	b := map[int]string{1: "one", 2: "two", 3: "three"}
	//可以看出该比较会直接报错,表示无法进行比较
	//t.Log(a == b)
	t.Log(reflect.DeepEqual(a, b))

	s1 := []int{1, 2, 3}
	s2 := []int{1, 2, 3}
	s3 := []int{2, 1, 3}
	t.Log("s1 == s2?", reflect.DeepEqual(s1, s2))
	t.Log("s2 == s3?", reflect.DeepEqual(s2, s3))
}
