package main

/*
循环语句For
	Go只有for一个循环语句关键字, 但支持3种形式
	初始化和布进表达式可以是多个值
	条件语句每次循环都会被重新检查, 因此不建议在条件语句中使用函数,
	尽量提前计算好以变量或常量代替
for init; condition; post { } 和c语言的for一样
for condition { } 和c语言的while一样
for { } 和c语言的for(;;)一样
	init： 一般为赋值表达式，给控制变量赋初值
	condition： 关系表达式或逻辑表达式，循环控制条件
	post： 一般为赋值表达式，给控制变量增量或减量

for 循环的 range 格式可以对 slice(切片)、map、数组、字符串等进行迭代循环
for key, value := range oldMap {
	newMap[key] = value
}

for condition {} / for {}: 迭代变量需提前定义初始值
*/

import (
	"fmt"
	// "reflect"
	"strconv"
)

func forTest() {
	var x = 0
	for {
		x++
		if x > 3 {
			fmt.Println(x)
			break
		}
		fmt.Println("Over")
	}
	var sum = 0
	// for var i= 1; i <= 100; i++ { // syntax error: var declaration not allowed in for initializer
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)

	var newsum = 0
	var j = 1
	for j <= 100 {
		newsum += j
		j++
	}
	fmt.Println(newsum)

	nums := [6]int{10, 2, 3, 5}

	for x, y := range nums {
		fmt.Printf("nums[%d] is %d\n", x, y)
	}

	af := string(65)
	fmt.Println(af) // A

	var fa int = 65
	bf := strconv.Itoa(fa)
	fa, _ = strconv.Atoi(bf)
	fmt.Println(fa)
}

/*
选择语句switch
	可以使用任何类型或表达式作为条件语句
	不需要写break, 一旦条件符合自动终止
	如希望继续执行下一个case, 需使用fallthrough语句
	支持一个初始化表达式(可以是并行方式), 右侧需跟分号
	左括号必须和条件语句在同一行
*/
func switchTest() {
	a := 1
	// switch a {  // 全局变量
	// case 0:
	// 	fmt.Println("a = 0")
	// case 1:
	// 	fmt.Println("a = 1")
	// default:
	// 	fmt.Println("None")
	// }

	switch {
	// switch a { 当判断条件有 >/<时, 则会报错
	case a >= 0:
		fmt.Println("a >= 0")
		fallthrough
	case a >= 1:
		fmt.Println("a >= 1")
	default:
		fmt.Println("None")
	}

	switch x := 1; { // 作用域是局部的, 跳出循环之后, 未定义
	case x >= 0:
		fmt.Println("x >= 0")
		fallthrough // 当满足条件之后, 继续执行程序
	case x >= 1:
		fmt.Println("x >= 1")
	default:
		fmt.Println("None")
	}
}

/*
 跳转语句:goto, break, continue
	3个语句都可以配合标签使用
	标签名区分大小写, 若不适用会造成编译错误
	Break与continue配合标签可用于多层循环的跳出
	Goto是调整执行位置, 与其它2个语句配合标签的结果并不相同
*/
func main2() {
	var a int = 10
LOOP:
	// a的值为: 10
	// a的值为: 11
	// a的值为: 12
	// a的值为: 13
	// a的值为: 14
	// a的值为: 16
	// a的值为: 17
	// a的值为: 18
	// a的值为: 19
	for a < 20 {
		if a == 15 {
			// 跳过迭代
			a++
			goto LOOP // 调整执行位置
		}
		// fmt.Printf("a的值为: %d\n", a)
		a++
	}
	// LOOP:
	// a的值为: 10
	// a的值为: 11
	// a的值为: 12
	// a的值为: 13
	// a的值为: 14

LABEL1:
	for {
		for i := 0; i < 10; i++ {
			if i > 3 {
				break LABEL1 // 跳出和LABEL1同一级别的循环, 程序结束
			}
		}
	}
	fmt.Println("OK")

	for {
		for i := 0; i < 10; i++ {
			if i > 3 {
				goto LABEL2 // 调整执行位置, 重新执行循环体
			}
		}
	}
LABEL2: // 为避免死循环, 需将LABEL放在for执行之后

	fmt.Println("OK")

LABEL3: // for 与 LABEL3属于同一级别循环体
	for i := 0; i < 10; i++ { // 为避免死循环, 第一级循环体为有限循环
		for {
			// for i := 0; i < 10; i++ {
			// if i > 3 {
			fmt.Println(i)
			continue LABEL3 // 跳出当前循环体, 继续执行和LABEL3同一级别的循环
			// 0
			// ...
			// 10
			// goto LABEL3 // 重新执行循环体
			// 0
			// 0
			// ...
			// }
		}
	}
	fmt.Println("OK")

	for i := 0; i < 10; i++ { // 为避免死循环, 第一级循环体为有限循环
		for {
			fmt.Println(i)
			goto LABEL4 // 调整执行位置, 退出循环
		}
	}
LABEL4:
	// 0
	fmt.Println("OK")
}
