package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

/*
并发concurrency
	goroutine只是由官方实现的超级"线程池"
	每个实例4-5kb的栈内存占用和由于实现机制而大幅减少的创建和销毁开销,
	是制造Go号称的高并发的根本原因
	goroutine简单易用, 也在语言层面上给予了开发者巨大的便利

并发主要又切换时间片来实现"同时"运行, 在并行是直接利用多核实现多线程的运行,
但Go可以设置使用核数, 以发挥多核计算机的能力

Goroutine奉行通过通信来共享内存, 而不是共享内存来通信

Channel
	Channel是gotroutine沟通的桥梁, 大都是阻塞同步的
	通过make创建, close关闭
	Channel是引用类型
	可以使用for rang来迭代不断操作channel
	可以设置单向或双向通道
	可以设置缓存大小, 在未被填满前不会发生阻塞(异步)

Select
	可处理一个或多个channel的发送与接收
	同时有多个可用的channel时按随机顺序处理
	可用空的select来阻塞main函数
	可设置超时
*/

func mainSingleChan() {
	// go goTest() // goTest, main同时都在运行, 当main函数暂停时, goTest才执行fmt
	// time.Sleep(2 * time.Second)

	// c := make(chan bool)
	// // c := make(chan bool, 2) //设置2个缓存, 实现异步
	// go func() {
	// 	fmt.Println("GO GO...")
	// 	c <- false // 存入通道
	// 	c <- true
	// 	close(c) // fatal error: all goroutines are asleep - deadlock!
	// }()
	// // <-c // 关闭通道
	// // 不断的进行迭代, 直到close(), 才退出函数
	// for v := range c {
	// 	fmt.Println(v)
	// }

	// runtime.GOMAXPROCS(runtime.NumCPU())
	// c := make(chan bool)
	// for i := 0; i < 10; i++ {
	// 	go goTest(c, i)
	// }
	// <-c // 当<-c能够取到值时, 程序便退出了

	// 由于有多个核在并行执行, 顺序是无序的,
	/*
		PS C:\Users\wuchan4x\Desktop\eilinge\go_study\base> go run .\goConcurrency.go
		9 4999950001
		1 4999950001
		PS C:\Users\wuchan4x\Desktop\eilinge\go_study\base> go run .\goConcurrency.go
		0 4999950001
		8 4999950001
		9 4999950001
	*/

	runtime.GOMAXPROCS(runtime.NumCPU())
	// c := make(chan bool, 10) // 未防止有些线程未执行结束就退出, 则可以设置与线程相同的缓存
	// for i := 0; i < 10; i++ {
	// 	go goTest(c, i)
	// }

	// for i := 0; i < 10; i++ {
	// 	<-c
	// }
	wg := sync.WaitGroup{} // 通过同步包, 等待所有线程结束退出
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go goTest(&wg, i)
	}
	wg.Wait()

}

func goTest(wg *sync.WaitGroup, index int) {
	// fmt.Println("Go")
	a := 1
	for i := 0; i < 100000; i++ {
		a += i
	}
	fmt.Println(index, a)

	// c <- true

	// for index == 9 { // 当执行到9时, 通知main函数可以退出
	// 	c <- true
	// }

	wg.Done()
}

func mainDoubleChan() {
	c1, c2 := make(chan int), make(chan string)

	o := make(chan bool, 2)
	go func() {
		for {
			select {
			case v, ok := <-c1:
				// 当关闭chan时, o通道中存入true, 退出匿名函数
				if !ok {
					o <- true
					break
				}
				fmt.Println("c1", v)

			case v, ok := <-c2:
				if !ok {
					o <- true
					break
				}
				fmt.Println("c2", v)
			}
		}
	}()

	c1 <- 1
	c2 <- "hi"
	c1 <- 3
	c2 <- "hello"

	close(c1)
	// close(c2)

	for i := 0; i < 2; i++ {
		<-o
	}
}

func setTimeOut() {
	// 可设置超时
	c := make(chan bool)
	select {
	case v := <-c:
		fmt.Println(v)
	case <-time.After(3 * time.Second):
		fmt.Println("Time out")
	}
}

var c chan string

func pingPong() {
	i := 0
	for {
		fmt.Println(<-c)                              // 等待接收
		c <- fmt.Sprintf("From pingPong: HI, #%d", i) // string存入chan
		i++
	}
}
func main() {
	c = make(chan string)
	go pingPong()
	for i := 0; i < 10; i++ {
		c <- fmt.Sprintf("From main: Hello, #%d", i) // string存入chan
		fmt.Println(<-c)                             //等待接收
	}
}
