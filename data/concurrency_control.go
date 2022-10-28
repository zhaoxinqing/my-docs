package data

import (
	"fmt"
	"time"
)

type Demo struct {
	input         chan string
	output        chan string
	goroutine_cnt chan int
}

func NewDemo() *Demo {
	d := new(Demo)
	d.input = make(chan string, 8192)
	d.output = make(chan string, 8192)
	d.goroutine_cnt = make(chan int, 10)
	return d
}
func (demo *Demo) Goroutine() {
	demo.input <- time.Now().Format("2006-01-02 15:04:05")
	time.Sleep(time.Millisecond * 500)
	<-demo.goroutine_cnt
}
func (demo *Demo) Handle() {
	for t := range demo.input {
		fmt.Println("datatime is :", t, "goroutine count is :", len(demo.goroutine_cnt))
		demo.output <- t + "handle"
	}
}
func ConcurrencyControlTest() {
	demo := NewDemo()
	go demo.Handle()
	for i := 0; i < 10000; i++ {
		demo.goroutine_cnt <- 1
		go demo.Goroutine()
	}
	close(demo.input)
}

// 如上边示例，Goroutine()函数，每隔500毫秒写入一个时间戳到管道中，不考虑管道的读取时间，
// 也就是说，每个Goroutine会存在大概500毫秒时间，如果不做控制的话，一瞬间可以开启上万个甚至更多的goroutine出来，这样系统就会奔溃。

// 在上述代码中，我们引入了带10个buffer的chan int字段，
// 每创建一个goroutine时，就会向这个chan中写入一个1，每完成一个goroutine时，就会从chan中弹出一个1。
// 当chan中装满10个1时，就会自动阻塞，等待goroutine执行完，弹出chan中的值时，才能继续开启goroutine。
// 通过chan阻塞特点，实现了goroutine的最大并发量控制
