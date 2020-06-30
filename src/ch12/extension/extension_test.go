package extension

import (
	"fmt"
	"testing"
)

type Pet struct {
}

func (p *Pet) Speak() {
	fmt.Print("...")
}

func (p *Pet) SpeakTo(host string) {
	p.Speak()
	fmt.Println(" ", host)
}

//当直接填入Pet的时候,go有默认的设置,dog可直接调用pst的方法和属性
type Dog struct {
	//p *Pet
	Pet
}

func (d *Dog) Speak() {
	fmt.Print("wang ")
}

//func (d *Dog) SpeakTo(host string) {
//	d.Speak()
//	fmt.Println(" ", host)
//}

func TestDog(t *testing.T) {
	//dog := new(Dog)
	//dog.SpeakTo("lilei")
	dog := new(Dog)
	//可以看到始终无法实现类型转换,因为go不支持继承,所以不支持类型转换
	//var p = (*Pet)(dog)
	//想查看方法适口可以重栽
	dog.SpeakTo("chao")
}
