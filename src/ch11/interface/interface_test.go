package interface_test

import "testing"

type Programmer interface {
	WriteHelloWorld() string
}

//GoProgrammer实现了Programmer,但是并没有显示继承声明,采用了duck type,只要看方法签名一样就可以
//也就是说可以先写很多方法,后方发现签名是重复的,就可以声明一个interface,用来整合
type GoProgrammer struct {
}

func (gp *GoProgrammer) WriteHelloWorld() string {
	return "fmt.Println(\"Hello World\")"
}

func TestClient(t *testing.T) {
	var gp Programmer
	gp = new(GoProgrammer)
	t.Log(gp.WriteHelloWorld())
}
