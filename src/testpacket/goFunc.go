package main

import "fmt"

/*
函数function
	GO函数不支持嵌套, 重载和默认参数
	但支持以下特性:
		无需声明原型
		不定长度变参
		多返回值
		命名返回值参数
		匿名函数
		闭包
	关键字:func, 且左大括号不能另起一行
	函数也可以作为一种类型使用
*/

func mainTest() {
	// x, y := 1, 2  // 值传递, 对传递后的值修改, 不会影响原值: 1 2
	// atest(x, y)

	// s := []int{1, 2, 3, 4} // slice: 引用传递, 对传递后的值修改, 影响原值: [3 4 3 4]
	// a := 2
	// atest(a)  // a int: 值传递
	// atest(&a) // a *int: 引用传递
	// fmt.Println(a)

	a := func() {
		fmt.Println("lambal func")
	}
	x := 3
	b := atest // 函数类型的使用
	a()
	b(&x)

	m := closure(10)
	fmt.Println(m(1)) // 11
	fmt.Println(m(2)) // 12

}

// func atest(a []int) { // 不定长变参, 必须放在最后一个参数
// 	a[0] = 3
// 	a[1] = 4
// 	fmt.Println(a)
// }

// func atest(a ...int) {}  // 不定长变参, 必须放在最后一个参数
func atest(a *int) {
	*a = 3
	fmt.Println(*a)
}

func closure(x int) func(int) int { // 闭包中进行值拷贝
	fmt.Printf("%p\n", &x) // 0xc00005e0a0
	return func(y int) int {
		fmt.Printf("%p\n", &x) // 0xc00005e0a0
		return x + y
	}
}

/*
Defer
	执行方式类似其他语言中的析构函数, 在函数体执行结束后, 按照调用顺序的相反顺序逐个执行
	即使函数发生严重错误也会执行
	支持匿名函数的调用
	常用于资源清理, 文件关闭, 解锁, 记录时间等操作
	通过与匿名函数配合可在return之后修改函数计算结果
	如果函数体内某个变量作为defer时匿名函数的参数, 则在定义defer时即已获得了拷贝, 否则则是引用某个变量的地址

	Go没有异常机制, 但有panic/recover 模式来处理错误
	Panic 可以在任何地方引发, 终止程序运行; 但recover只有在defer调用的函数中有效
*/

func mainFunc() {
	// for i := 0; i < 3; i++ {
	// 	// defer fmt.Println(i) // 2 1 0
	// 	defer func() {
	// 		fmt.Println(i) // 3 3 3: 当for循环体退出时: i=3, defer调用的是i=3的赋值
	// 	}() // (): 进行调用
	// }

	// a()
	// b()
	// c()

	var fs = [4]func(){}

	for i := 0; i < 4; i++ {
		defer fmt.Println("defer i = ", i) // 未引用闭包中的i, 值拷贝: 0, 1, 2, 3
		defer func() {
			fmt.Println("defer_closure i = ", i) // func()未传入参数, i为闭包中的i, 引用传递:4
		}()
		fs[i] = func() {
			fmt.Println("closure i = ", i) // func()未传入参数, i为闭包中的i, 引用传递:4
		}
	}

	for _, f := range fs { // 未使用defer声明的函数, 最先执行
		f()
	}
	/*
		closure i =  4
		closure i =  4
		closure i =  4
		closure i =  4
		defer_closure i =  4
		defer i =  3
		defer_closure i =  4
		defer i =  2
		defer_closure i =  4
		defer i =  1
		defer_closure i =  4
		defer i =  0
	*/
}

/*
func a() {
	fmt.Println("Func A")
}

func b() {
	defer func() { // 提前定义, 在panic之后, 进行恢复调用: Func A, Recover in B, Func C
		if err := recover(); err != nil {
			fmt.Println("Recover in B")
		}
	}()
	panic("Func B") // Func A, panic: Func B, err
}

func c() {
	fmt.Println("Func C")
}
*/
