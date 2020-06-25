package string

import "testing"

func TestString(t *testing.T) {
	var s string
	t.Log(s)
	s = "hello"
	//尝试改变字符串切片的第二位报错,说明string是一个不可变更的切片
	//s[1] = '3'
	t.Log(len(s))
	//从其编码上可以看出是三个段组成的
	//s = "\xE4\xB8\xA5"
	//该编码为随意打上的,应该不属于任何一个编码集,为了证明string结构可以存储任意的二进制数据
	s = "\xE4\xBA\xBB\xFF"
	t.Log(s)
	t.Log(len(s))
	s = "中"
	t.Log(len(s))
	//rune用来提取出unicode的切片
	c := []rune(s)
	t.Log(len(c))
	t.Logf("中 unicode %x", c[0])
	t.Logf("中 UTF8 %x", s)
}

func TestStringToRune(t *testing.T) {
	s := "中华人民共和国"
	for _, c := range s {
		//因为打印出的是rune而不是原来的字节
		t.Logf("%[1]c %[1]x", c)
	}
}
