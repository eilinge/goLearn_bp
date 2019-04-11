package main

import (
	"fmt"
)

/*
数组/结构体 slice(切片)/映射(map)
每个数组元素都是完全相同的类型——结构体则是由异构的元素组成的。数组和结构体都是有固定内存大小的数据结构。
相比之下，slice和map则是动态的数据结构，它们将根据需要动态增长。
*/

/*
数组Array
	格式: var <varName> [n]<type>, n>0
	数组长度也是类型的一部分, 因此具有不同长度的数组为不同类型
	注意区分指向数组的指针和指针数组
	数组在Go中为值类型
	数组之间可以使用== 或 != 进行比较, 不可以使用<或>
	可以使用new创建数组, 此方法返回一个指向数组的指针
	Go 支持多维数组
*/
// var m [3]int = [3]int{1, 2, 3, 4}
var m [3]int = [3]int{1, 2, 3}
var n = [...]int{1, 2, 3}

// fmt.Println(m[1])  无法放置在函数体外

func arrayTest() {
	fmt.Println(m[1], len(m))
	fmt.Println(n[1], len(n))
	fmt.Println(n == m) // true
}

// 数组作为函数的参数传入, 进行了拷贝, 在函数内部改变数组的值, 是不影响到外面的数组的值
func ArrIsArgs(arr *[4]int) {
	arr[0] = 120
}

func arrayRunTest() {
	x := [...]int{1, 2, 3, 4}
	// ArrIsArgs(x)  // [1 2 3 4]
	ArrIsArgs(&x) // [120 2 3 4]

	// a := [...]int{19: 1}  // [0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 1]
	// a := [...][3]int{{1: 2}, {2: 3}}  // [[0 2 0] [0 0 3]]

	// a := new([...]int)  // 定义错误
	// a := new([2]int)  // &[0, 0]

	// m, n := 1, 2
	// a := [...]*int{&m, &n}  // 指针数组: [0xc00005e058 0xc00005e070]

	var a *[4]int = &x // 指向数组的指针: &[120 2 3 4]
	fmt.Println(x)
	fmt.Println(a)

	y := [...]int{1, 7, 4, 6, 2, 9, 2}
	ll := len(y)
	for i := 0; i < ll-1; i++ { // for循环中, i每一次循环, 指向的地址都不相同, 所以可以 i:=, 而不能用var声明变量
		for j := 0; j < ll-1; j++ {
			if y[j] > y[j+1] {
				temp := y[j]
				y[j] = y[j+1]
				y[j+1] = temp
			}
		}
	}
	fmt.Println(y)
}

/*
slice 切片
	其本身不是数组, 它指向底层的数组
	作为变长数组的替代方案, 可以关联底层数组获取生成
	引用类型, 如果多个slice指向相同底层数组, 其中一个的值改变会影响全部
	可以直接创建或从底层数组获取生成
	使用len()获取元素个数, cap()获取容量
	一般使用make() 创建

	make([]T, len, cap)  // append当len > cap时, 系统会自动成倍扩容, 指向一个新的内存地址的切片, 释放原来的内存
	cap可以省略, 则和len的值相同
	len表示存数的元素个数, cap表示容量, cap >= len
*/

func mainSlice() {
	// a := make([]int, 5, 7)
	// fmt.Println(reflect.TypeOf(a)) // []int
	// fmt.Println(&a) // &[0 0 0 0 0]

	a := []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g'}
	// a := []byte{"a", "b", "c", "d", "e"}  // cannot convert "e" (type untyped string) to type byte
	sl := a[:5]
	fmt.Println(string(sl))

	/*
		Reslice
			Reslice时索引以被slice的切片为准
			索引不可以超过被slice的切片的容量cap()值
			索引越界不会导致底层数组得重新分配而是引发错误
	*/

	sa := a[3:5]
	sas := sa[1:]
	as := sa[1:3]
	fmt.Println(len(sa), cap(sa))                    // 2 4
	fmt.Println(string(sa), string(sas), string(as)) // de e ef
	// asa := sa[1:6]
	// fmt.Println(string(asa))  // slice bounds out of range

	/*
		append
			可以在slice尾部追加元素
			可以将一个slice追加在另一个slice尾部
			如果最终长度未超过追加到slice的容量则返回原始slice
			如果超过追加到的slice的容量则返回原始slice
			如果超过追加到的slice的容量则将重新分配数组并拷贝原始数据
	*/

	s1 := make([]int, 3, 6)
	fmt.Printf("%p\n", s1)
	s1 = append(s1, 1, 2, 3)
	fmt.Printf("%v %p\n", s1, s1) // [0 0 0 1 2 3] 0xc000084030
	s1 = append(s1, 1, 2, 3)
	fmt.Printf("%v %p\n", s1, s1) // [0 0 0 1 2 3 1 2 3] 0xc000050060

	s3 := []int{1, 2, 3, 4, 5}
	s4 := s3[2:5]
	s5 := s3[1:3]
	fmt.Println(s4, s5) // [3 4 5] [2 3]
	s4[0] = 6
	fmt.Println(s4, s5) // 未重新分配地址, 改变相同的引用, 则影响另一个引用的值: [6 4 5] [2 6]
	s4 = append(s4, 1, 3, 4, 5, 9, 6, 8)
	s4[0] = 10
	fmt.Println(s4, s5) // 重新分配地址, 改变相同的引用, 不影响另一个引用的值: [10 4 5 1 3 4 5 9 6 8] [2 6]

	/*
		copy
		a := copy(s1, s2)
		len(s1) > len(s2)
		[10 4 5] [1 3 4 5 9]
		[1 3 4]

		len(s1) < len(s2)
		[10 4 5] [1 3]
		[1 3 5]
	*/
	c1 := s4[:3]
	c2 := s4[3:8]
	fmt.Println(c1, c2)
	copy(c1, c2)

	fmt.Println(c1)

	copy(c2, c1)
	// 1 3 4 5 9
	fmt.Println(c2)

}

type Rectangle struct {
	width  float64
	length float64
}

func structTest() {
	// var r = Rectangle{width: 100, length: 200}
	var r = Rectangle{100, 200}
	fmt.Println(r.width * r.length)
}
