package data

import "fmt"

// recorver作用：采用 recover() 函数来捕获异常，使得即使报 panic,程序也能继续运行下去。

// 通常写法
// defer func() {
// 	if err := recover(); err !=nil {
// 		fmt.Println(err)
// 	}
// }()

func PanicRecorver() {
	test1() //输出：this is test 1
	test2() //输出：this is test 2  panic: test 2 is panic   直接挂掉
	test3()
}

func test1() {
	fmt.Println("this is test 1")
}

func test2() {
	// 写在这里才能捕获 panic ,才能执行test3
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	fmt.Println("this is test 2")
	panic("test 2 is panic")
}

func test3() {
	fmt.Println("this is test 3")
}
