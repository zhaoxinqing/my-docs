package data

import "fmt"

func ChannleRead(ch1 chan int, ch2 chan bool) {
	for {
		v, ok := <-ch1
		if ok {
			fmt.Printf("read a int is %d\n", v)
		} else {
			ch2 <- true
		}
	}

}

func ChannelWrite(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}

func ChannelCloseTest() {
	ch1 := make(chan int)
	ch2 := make(chan bool)

	go ChannelWrite(ch1)
	go ChannleRead(ch1, ch2)

	<-ch2
}
