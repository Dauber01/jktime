package reflect

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func TestTypeAndValue(t *testing.T) {
	var i int64 = 12
	//可以看到两个函数分别可以打出对应的返回值,而且ValueOf()的Type()也可以获取对应的返回值
	t.Log(reflect.TypeOf(i), reflect.ValueOf(i))
	t.Log(reflect.ValueOf(i).Type())
}

func CheckType(v interface{}) {
	t := reflect.TypeOf(v)
	//可以通过kind()来进行判断对应的类型
	switch t.Kind() {
	case reflect.Float64, reflect.Float32:
		fmt.Println("Float")
	case reflect.Int32, reflect.Int64:
		fmt.Println("Int")
	default:
		fmt.Println("can not find the type", t)
	}
}

func TestCheckType(t *testing.T) {
	var f float64 = 3
	//可以看到,当是float的时候输出Float,但是是指针类型的时候,会找不到,并按照日志打印出对应的类型
	CheckType(&f)
}

type Employee struct {
	EmployeeID string
	Name       string `format:"normal"`
	Age        int
}

func fillBySetting(st interface{}, setting map[string]interface{}) error {

	//判断,类型不为指针,指向的不为struct类型,则返回,reflect.TypeOf(st).Elem()表示获取指针指向的结构体的类型
	if reflect.TypeOf(st).Kind() != reflect.Ptr || reflect.TypeOf(st).Elem().Kind() != reflect.Struct {
		return errors.New("the type must be pointor and must point to struct")
	}

	if setting == nil {
		return errors.New("setting is nil")
	}

	var (
		field reflect.StructField
		ok    bool
	)

	for k, v := range setting {
		if field, ok = reflect.ValueOf(st).Type().Elem().FieldByName(k); !ok {
			continue
		}
		if field.Type == reflect.TypeOf(v) {
			reflect.ValueOf(st).Elem().FieldByName(k).Set(reflect.ValueOf(v))
		}
	}
	return nil
}

func TestFillNameAndAge(t *testing.T) {
	//可以看到,都已经如期同步过来了
	setting := map[string]interface{}{"Name": "Mike", "Age": 17}
	e := Employee{}
	if e := fillBySetting(&e, setting); e != nil {
		t.Fatal(e)
	}
	t.Log(e)
	c := new(Customer)
	if ex := fillBySetting(c, setting); ex != nil {
		t.Fatal(ex)
	}
	t.Log(*c)
}

//获取反射方法的方法名必须要首字母大写, 否则报错
func (e *Employee) UpdateAge(newV int) {
	e.Age = newV
}

type Customer struct {
	CookiID string
	Name    string `format:"normal"`
	Age     int
}

func TestInvokeByName(t *testing.T) {
	e := &Employee{"122", "张三", 30}
	//通过value方法进行获取
	t.Logf("Name :value(%[1]v), type(%[1]T)", reflect.ValueOf(*e).FieldByName("Name"))
	//通过type方法进行获取,注意type方法有两个返回值
	if nameFiled, ok := reflect.TypeOf(*e).FieldByName("Name"); !ok {
		t.Error("fail to get the 'name' filed")
	} else {
		t.Log("Tag:format", nameFiled.Tag.Get("format"))
	}
	/* 改了好几遍程序，终于明白为什么要写成这样了， e := &Employee{1, "Mike", 30}
	原因是：
	1. UpdateAge 必须得声明为 e *Employee 类型的成员变量才能将修改的值保存；
	2. Call 调用 UpdateAge 就必须得传入 *Employee 类型的变量才能调用成功。 */
	//所以方法的反射看起来是值传递么？
	//反射通过名字调用函数的实例
	reflect.ValueOf(e).MethodByName("UpdateAge").Call([]reflect.Value{reflect.ValueOf(1)})
	t.Log("update age:", e)
}
