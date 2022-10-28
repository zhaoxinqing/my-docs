package app

import (
	"fmt"
	"time"
)

// 无缓冲的 channel 和 有缓冲的 channel 的区别？
// 对于无缓冲的 channel，发送方将阻塞该信道，直到接收方从该信道接收到数据为止，而接收方也将阻塞该信道，直到发送方将数据发送到该信道中为止。
// 对于有缓存的 channel，发送方在没有空插槽（缓冲区使用完）的情况下阻塞，而接收方在信道为空的情况下阻塞。

func NoBufferChannel() {
	st := time.Now()
	ch := make(chan bool)
	go func() {
		time.Sleep(time.Second * 2)
		<-ch
	}()
	ch <- true // 无缓冲，发送方阻塞直到接收方接收到数据。
	fmt.Printf("cost %.1f s\n", time.Now().Sub(st).Seconds())
	time.Sleep(time.Second * 5)
}

func BufferChannel() {
	st := time.Now()
	ch := make(chan bool, 2)
	go func() {
		time.Sleep(time.Second * 2)
		<-ch
	}()
	ch <- true
	ch <- true                                                // 缓冲区为 2，发送方不阻塞，继续往下执行
	fmt.Printf("cost %.1f s\n", time.Now().Sub(st).Seconds()) // cost 0.0 s
	ch <- true                                                // 缓冲区使用完，发送方阻塞，2s 后接收方接收到数据，释放一个插槽，继续往下执行
	fmt.Printf("cost %.1f s\n", time.Now().Sub(st).Seconds()) // cost 2.0 s
	time.Sleep(time.Second * 5)
}
