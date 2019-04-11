package main

import (
	"fmt"
	"reflect"
)

/*
反射reflection
	反射可大大提高程序的灵活性, 使得interface{} 有更大的发挥余地
	反射使用TypeOf 和 ValueOf函数从接口中获取目标对象信息
	反射会将匿名字段作为独立字段(匿名字段本质)
	想要利用反射修改对象状态, 前提是interface.data是settbale, 即pointer-interface
	通过反射可以"动态"调用方法
*/

type user struct {
	ID   int
	Name string
	Age  int
}

type manager struct {
	user
	title string
}

func (u user) hello(name string) {
	fmt.Println("Hello, ", name, "my name is", u.Name)
}

func mainReflect() {
	u := user{1, "eilinge", 17}
	// info(u)
	m := manager{user: user{1, "ok", 2}, title: "123"}
	t := reflect.TypeOf(m)

	// reflect.StructField{Name:"user", PkgPath:"main", Type:(*reflect.rtype)(0x4b2200),
	// Tag:"", Offset:0x0, Index:[]int{0}, Anonymous:true}  // 是否匿名
	fmt.Printf("%#v\n", t.Field(0))

	// reflect.StructField{Name:"title", PkgPath:"main", Type:(*reflect.rtype)(0x4a01c0),
	// Tag:"", Offset:0x20, Index:[]int{1}, Anonymous:false}  // 是否匿名
	fmt.Printf("%#v\n", t.Field(1))

	// reflect.StructField{Name:"ID", PkgPath:"", Type:(*reflect.rtype)(0x4a0b40), Tag:"",
	// Offset:0x0, Index:[]int{0}, Anonymous:false}
	fmt.Printf("%#v\n", t.FieldByIndex([]int{0, 0}))

	// reflect.StructField{Name:"Name", PkgPath:"", Type:(*reflect.rtype)(0x4a11c0), Tag:"",
	// Offset:0x8, Index:[]int{1}, Anonymous:false}
	fmt.Printf("%#v\n", t.FieldByIndex([]int{0, 1}))

	x := 123
	y := reflect.ValueOf(&x) // 传入指针, 引用传递
	y.Elem().SetInt(999)
	// y.Elem().SetString("true") int无法重新设置成其他类型
	fmt.Println(x) // 999

	u.hello("joe") // Hello,  joe my name is eilinge

	set(&u) // {1 set successfully 17}
	fmt.Println(u)

	// 通过反射, 动态调用方法
	v := reflect.ValueOf(u)
	mv := v.MethodByName("hello")

	args := []reflect.Value{reflect.ValueOf("joe")}
	fmt.Println(args)
	mv.Call(args)
}

// 参数为空接口时, 可以传入任何类型
func info(o interface{}) {
	t := reflect.TypeOf(o)
	fmt.Println("Type:", t.Name())

	// 判断传入的类型是否正确
	if k := t.Kind(); k != reflect.Struct {
		// if k := t.Kind(); k != reflect.Map {
		fmt.Println("XXX")
		return
	}
	v := reflect.ValueOf(o)
	fmt.Println("Fields:")

	// 反射出类型的字段
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		val := v.Field(i).Interface()
		fmt.Printf("%6s: %v = %v\n", f.Name, f.Type, val)
	}

	// 反射出类型的方法
	for i := 0; i < t.NumMethod(); i++ {
		// for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Printf("%6s: %v\n", m.Name, m.Type)
	}
}

func set(o interface{}) {
	v := reflect.ValueOf(o)

	// reflect.Ptr: pointer-interface
	// v.Elem().CanSet(): 是否能被修改
	if v.Kind() == reflect.Ptr && !v.Elem().CanSet() {
		fmt.Println("xxx")
		return
	}
	// 进行赋值, 取得实际的对象
	v = v.Elem()

	f := v.FieldByName("Name") // 取出v的Name的字段, 然后进行修改
	if !f.IsValid() {          // 判断v中是否含有"Name"字段
		fmt.Println("BAD")
		return
	}
	if f.Kind() == reflect.String { // 判断是否为string类型字段
		f.SetString("set successfully")
	}
}
