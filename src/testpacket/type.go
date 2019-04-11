package main

import (
	"fmt"
)

func test() {
	// var x float64 = 3.4
	// var x = 3.4
	// var y string = "abc"
	// var y = "abc"
	// var a uint8 = 10
	// var b uint16 = 10
	// var c int = 100
	// var c = 100

	// fmt.Println("x's type:", reflect.TypeOf(x))
	// fmt.Println("y's type:", reflect.TypeOf(y))
	// fmt.Println("a's type:", reflect.TypeOf(a))
	// fmt.Println("b's type:", reflect.TypeOf(b))
	// fmt.Println("c's type:", reflect.TypeOf(c))

	// 1. var v_name v_type,指定类型,声明若不赋值,则使用默认值0
	var a int
	fmt.Printf("a = %d\n", a)
	a = 10
	fmt.Printf("a = %d\n", a)

	// 2. var v_name [v_type]= value,声明直接赋值,类型可以省略,这样编译器自行判断
	var b int = 101
	var c = 3.14
	fmt.Printf("b = %d, c = %f\n", b, c)

	// 3. v_name := value ,省略var,:= 左侧的变量不能是声明过的,否则编译会错误
	str := "yekai"

	// 引用传递
	// &符号代表取变量的地址，*代表取地址对应的内容。
	// y执行是x的地址单元，地址单元内容被修改，那么x再读取数据也受影响
	x := 99
	fmt.Printf("str = %s, x=%d\n", str, x)

	y := &x  // 传入指针
	*y = 100 // 对引用修改, 其他引用该地址的值咋也会变
	fmt.Printf("str = %s, x=%d\n", str, x)

	// str = yekai, x=99
	// str = yekai, x=100
}

// var ( // 多用于全局变量声明
// 	a1 int
// 	b1 bool
// )

var x1, y1 int
var c1, d1 int = 3, 4
var e1, f1 = "yekai", 3.14

// h, g := 123, "只能在函数内部声明"

func main1() {
	h1, g1 := 123, "只能在函数内部声明"
	fmt.Println(x1, y1, c1, d1, e1, f1, h1, g1)

	_, value := 3, 5 // _代表占位符, 不关心第一个右值, fmt.Println(_) 无法获取_的值
	fmt.Println(value)

}
