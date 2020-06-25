package slice_test

import "testing"

func TestSliceLen(t *testing.T) {
	//通过var的形式进行初始化
	var s0 []int
	//len()表示长度,cap()函数返回容量
	t.Log(len(s0), cap(s0))
	//在初始化的时候直接进行赋值
	s1 := []int{1, 2, 3, 4}
	t.Log(len(s1), cap(s1))
	//append函数对切片进行添加,添加在下一位
	s1 = append(s1, 5)
	t.Log(len(s1), cap(s1))
	//可以通过make函数进行初始化
	s2 := make([]int, 5, 8)
	t.Log(len(s2), cap(s2))
	t.Log(s2[0], s2[1], s2[2], s2[3], s2[4])
	s2 = append(s2, 1)
	t.Log(len(s2), cap(s2), s2[5])
}

func TestSliceCap(t *testing.T) {
	var s []int
	for i := 0; i < 10; i++ {
		//可以看出随着元素的增加,每当原来的容量存不下的时候,需要对其进行原来的2倍扩容,
		//所以append函数需要进行重新赋值，因为有可能存储地址变化了,所以,在调优的时候要注意,
		//因为变化的时候会发生大量的数据复制
		s = append(s, i)
		t.Log(len(s), cap(s))
	}
}

func TestSliceShareMemory(t *testing.T) {
	year := [12]string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
	Q2 := year[3:6]
	//可以看出,虽然切片的时候看起来只切了3-6之前的3位,但由于后面的存储空间是连续的,所以容量为开始到连续空间的结束
	t.Log(Q2, len(Q2), cap(Q2))
	summer := year[5:8]
	t.Log(summer, len(summer), cap(summer))
	//将切片2的对应位改变之后会发现切片1和数组都受到了影响,说明是共享内存
	summer[0] = "Unknow"
	t.Log(Q2)
	t.Log(year)
}

func TestSliceCompare(t *testing.T) {
	//分别定义切片a和b,并将他们进行比较,可以看到校验通不过,不能这么进行比较
	/*a := []int{1, 2, 3, 4}
	b := []int{1, 2, 3, 4}
	if a == b {
		t.Log("a == b")
	}*/
}
