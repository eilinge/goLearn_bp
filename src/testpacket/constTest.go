package main

import (
	"fmt"
	"unsafe"
)

// unsafe: 不论字符串的len有多大,sizeof始终返回16,这是为啥,字符串不是也是不可变的吗?
// 实际上字符串类型对应一个结构体,该结构体有两个域,第一个域是指向该字符串的指针,
// 第二个域是字符串的长度,每个域占8个字节,但是并不包含指针指向的字符串的内容,这也就是为什么sizeof始终返回的是16
// sizeof总是在编译期就进行求值，而不是在运行时，这意味着，sizeof的返回值可以赋值给常量

const (
	a = "abc"
	x = "abcdef"
	b = len(a)
	c = unsafe.Sizeof(a)
	d = unsafe.Sizeof(x)
)

// const的枚举中使用iota变量,实现自动增长或者跳档变化。
// 自动继承上一步操作, 并且遇到下一次操作不同时, 重新继承, 自动增长
const (
	login = 0
	logout
	user = iota // 2
	account
	delete = iota * 2 // 4*2
	update            // 5*2
)

const (
	apple, banana = iota + 1, iota + 2
	peach, pear   // 1+1, 1+2
	orange, mongo // 2+1, 2+2
)

func main() {
	// const 可以修饰变量,这样变量的内容不可修改。
	const pi = 3.14
	const r = 10
	const length, width = 10, 5 // 多重复制
	// r = 11
	var area1 = pi * r * r
	fmt.Println(area1)
	fmt.Println(length * width)
	fmt.Println(c, d)

	const ( // 可用于枚举
		Female = 1
		Male   = 2
	)

	fmt.Println(user, delete, update)
	fmt.Println(apple, banana, peach, pear, orange, mongo)

	// 一元运算
	// 算术运算符 +,-,*,/,%,++,--
	// 关系运算符 ==,!=,<=,>=,<,>
	// 位运算符  &,|,^,<<,>>
	// 赋值运算符  =,+=,-= ……

	// 二元运算
	// 逻辑运算符 &&,||,!

	/*
	   6:  0110
	   11: 1011
	   ---------
	   &   0010
	   |	 1111
	   ^   1101
	   &^  0100  // 根据第二个数, 当第二个数为1时, 第一个数1->0
	*/
	// <- (专门用于channel)

	// 使用常量实现计算机存储
	const (
		B float64 = 1 << (iota * 10)
		KB
		MB
		TB
	)

	fmt.Println(B, KB, MB, TB)

	// 指针
	// 在Go当中不支持指针运算以及 "->" 运算符,
	// 而直接采用 "." 选择符来操作指针目标对象的成员
	// "&" 取变量地址, "*"通过指针间接访问目标对象
	// 默认值为nil 而非NULL
	pr := 1
	var rp *int = &pr
	fmt.Println(rp) // 0xc000072168

	// 递增递减语句
	// ++ 与 -- 是作为语句而并不是作为表达式
	// a := a++ 无法实现
}
