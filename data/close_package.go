package data

import (
	"fmt"
	"runtime"
	"time"
)

// 闭包
func ClosePackage() {
	//
	fmt.Println("----------------引用传递----------------")
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println(i)
		}()
	}
	time.Sleep(10 * time.Millisecond)

	//
	fmt.Println("----------------值传递----------------")
	for i := 0; i < 10; i++ {
		go func(x int) {
			fmt.Println(x)
		}(i)
	}
	time.Sleep(10 * time.Millisecond)
}

// 闭包
func ClosureA10() {
	//设置最大的可同时使用的 CPU 核数 为1
	runtime.GOMAXPROCS(1)

	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("A:", i)
		}()
	}

	//主线程等待其他线程执行完毕
	time.Sleep(3 * time.Second)
}

//
// A: 10
// A: 10
// A: 10
// A: 10
// A: 10
// A: 10
// A: 10
// A: 10
// A: 10
// A: 10

//
func Closure() {
	//设置最大的可同时使用的 CPU 核数 为1
	runtime.GOMAXPROCS(1)

	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("B:", i)
		}(i)
	}

	//主线程等待其他线程执行完毕
	time.Sleep(3 * time.Second)
}

//
// B: 9
// B: 0
// B: 1
// B: 2
// B: 3
// B: 4
// B: 5
// B: 6
// B: 7
// B: 8

// 总结
// 调用go就会启动一个goroutine，这需要一定时间
// 启动一个goroutine的速度远小于循环执行的速度，所以即使是第一个goroutine刚起启动时，外层的循环也执行到了最后一步了。
// 由于所有的goroutine共享i，而且这个i会在最后一个使用它的goroutine结束后被销毁，所以最后的输出结果都是最后一步的A:10。
