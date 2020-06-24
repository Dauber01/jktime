package lcop

import "testing"

func TestWhileLoop(t *testing.T) {
	n := 0
	/* 相当于其他语言中的while(n<5) */
	for n < 5 {
		t.Log(n)
		n++
	}
}
