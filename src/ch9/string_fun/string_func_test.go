package string_fun

import (
	"strconv"
	"strings"
	"testing"
)

func TestStringSplit(t *testing.T) {
	s := "A,B,C"
	//可以通过split方法进行切割
	strs := strings.Split(s, ",")
	for _, v := range strs {
		t.Log(v)
	}
	t.Log(strings.Join(strs, "-"))
}

func TestStringConv(t *testing.T) {
	s := strconv.Itoa(10)
	t.Log("it is : " + s)
	//和将int转换为string不同,将string转换为int的时候有两个返回值
	if v, err := strconv.Atoi(s); err == nil {
		t.Log(10 + v)
	}
}
