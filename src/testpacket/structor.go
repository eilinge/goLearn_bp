package main

import (
	"fmt"
)

/*
结构Struct(面向对象编程)
	Go中的struct与C中的struct非常相似, 并且Go没有class
	使用type <Name> struct{} 定义, 名称遵循可见性
	支持指向自身的指针类型成员
	支持匿名函数, 可用作成员或定义成员变量
	匿名结构也可以用于Map的值
	可以使用字面量对结构进行初始化
	允许直接通过指针来读写结构成员
	相同类型的成员可以直接拷贝赋值
	支持 == 与 != 比较运算符, 不支持 >/<
	支持匿名字段, 本质上是定义了以某个类型名为名称的字段
	嵌入结构作为匿名字段看起来像继承, 但不是继承
	可以使用匿名字段指针
*/

type person struct {
	Name     string
	Age      int
	Contract struct { // 嵌套结构体, 之前未声明
		Name, City string
	}
}

type person1 struct {
	string
	uint
}

type human struct {
	Sex uint
}

type teacher struct {
	human // 结构体嵌套, 之前已声明
	Name  string
	Age   uint
}

type students struct {
	human
	Name string
	Age  uint
}

func mainStruct() {
	// p := person{"eilinge", 18}
	// p := person{Name: "eilinge", Age: 18}
	// a(p)  // 值传递: p.Age = 18
	// a(&p) // 引用传递: p.Age = 16

	p := &person{Name: "eilinge", Age: 18} // 初始化时, 直接对其进行取&符, 后续修改无需进行*p.Age修改, p.Age修改即可
	// a1(p)
	fmt.Println(p.Age)

	// 匿名结构体初始化
	// a := &struct {
	// 	Name string
	// 	Age  uint
	// }{
	// 	Name: "eilinge",
	// 	Age:  18,
	// }

	// a := person{Name: "lin", Age: 1}
	// a.Contract.Phone = "1231231"
	// a.Contract.City = "China"

	a := person1{"lin", 1}
	// a := person1{1, "1231"}  匿名字段初始化, 需与结构体中顺序相同赋值

	b := a              // 结构体可以进行赋值
	fmt.Println(b == a) // 可以进行比较是否相等

	c := teacher{Name: "eilin", Age: 17, human: human{Sex: 0}}
	// c.human.Sex = 1
	c.Sex = 1
	fmt.Println(c)

	d := person{Name: "duzi", Age: 1}
	d.Contract.Name = "eilinge"
	fmt.Println(d)

	// e := A{B: B{Name: "eilinge"}, C: C{Name: "duzi"}} // {{eilinge} {duzi}}
	// fmt.Println(e)
}

/*
type A struct {
	B
	C
}

type B struct {
	Name string
}

type C struct {
	Name string
}

func a1(per *person) {
	per.Age = 16
	fmt.Println(per.Age)
}
*/
