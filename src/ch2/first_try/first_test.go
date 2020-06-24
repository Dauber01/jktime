package first_try

import "testing"

//测试的源文件的名字必须以_test结尾,里面的函数必须以Test开头,函数名字的首字母大写表示可以跨包访问
func TestFirstTry(t *testing.T) {
	t.Log("hello word!")
}
