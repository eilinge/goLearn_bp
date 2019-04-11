package main

import (
	"fmt"
)

/*
接口interface
	接口是一个或多个方法签名的集合
	只要某个类型拥有了该接口的所有方法签名, 即算实现该接口, 无需显示声明实现了哪个接口, 这称为Structural Typing
	接口只有方法声明, 没有实现, 没有数据字段
	接口可以匿名嵌入其它接口, 或嵌入到结构中
	将对象赋值给接口时, 会发生拷贝, 而接口内部存储的是指向这个复制品的指针, 即无法修改复制品的状态,
	也无法获取指针
	只有当接口存储的类型和对象都是nil时, 接口才等于Nil
	接口调用不会做receiver的自动转换
	接口同样支持匿名字段方法
	接口也可实现类似OOP中的多态(无继承)
	空接口可以作为任何类型数据的容器
*/
type USB interface {
	Name() string
	Connector // 嵌入式
}

type Connector interface {
	Connect()
}
type phoneConnect struct {
	name string
}

// 只要某个类型拥有了该接口的所有方法(Name, Connect)签名, 即算实现该接口, 无需显示声明实现了哪个接口, 这称为Structural Typing
func (pc phoneConnect) Name() string { // 使结构体能够调用接口方法/var a USB=phoneConnect{"connectPhone"}编译成USB
	return pc.name
}
func (pc phoneConnect) Connect() { // 使结构体能够调用接口方法
	fmt.Println("Connect:", pc.name)
}

func mainInterface() {
	var a USB
	a = phoneConnect{"connectPhone"}
	a.Connect()           // Connect: connectPhone
	Disconnect(a)         // Disconnected: connectPhone
	fmt.Println(a.Name()) // connectPhone

	pc := phoneConnect{"phoneConnect"} // USB接口类型
	var b Connector
	b = Connector(pc) // USB嵌入了Connector: USB转换成Connector类型, 反之则不行
	b.Connect()       // Connect: phoneConnect

	// 将对象赋值给接口时, 会发生拷贝, 而接口内部存储的是指向这个复制品的指针, 即无法修改复制品的状态,
	pc.name = "pc"
	b.Connect() // Connect: phoneConnect

	// 只有当接口存储的类型和对象都是nil时, 接口才等于Nil
	var x interface{}
	fmt.Println(x == nil) // true

	var p *int = nil
	x = p
	fmt.Println(x == nil) // false
}

// 传入空接口
func Disconnect(usb interface{}) {
	// 判断phoneConnect在usb中是什么结构
	// if pc, ok := usb.(phoneConnect); ok {
	// 	fmt.Println("Disconnected:", pc.name)
	// 	return
	// }

	// 程序自动判断usb中的结构
	switch v := usb.(type) {
	case phoneConnect:
		fmt.Println("Disconnected:", v.name)
	default:
		fmt.Println("Unknown decive...")
	}
}
