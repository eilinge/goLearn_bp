package main

import (
	"fmt"
)

type a1 struct {
	Name string
}

type b1 struct {
	Name string
}

func (a *a1) print() { // 指针传递: 指向同一地址, 值会随着改变发生相应变化
	fmt.Println("ni hao methods: a")
}

func (b b1) print() { // 值拷贝: 简单拷贝值, 值不会随着改变发生相应变化
	b.Name = "eilinge" // 同一文件中, 方法可以直接访问公/私属性
	fmt.Println("ni hao methods: b", b.Name)
}

type tZ int

func (tz *tZ) print() {
	fmt.Println("type change TZ")
}

func (tz *tZ) increate(num int) {
	*tz += tZ(num) // 进行类型转换
}

func mainMethod() {
	a1 := a1{}
	a1.Name = "eilinge"
	a1.print()
	fmt.Println(a1.Name)

	b1 := b1{}
	b1.Name = "eilin"
	b1.print()
	fmt.Println(b1.Name)

	var c tZ
	c.print()

	// (*tZ).print(&c) // 另一种调用的方式

	c.increate(100)
	fmt.Println(c)
}

/*
	struct  = class
	(str struct) methods() = class.methods
	嵌套(优先级) = 继承
	多态
	可封装
*/
