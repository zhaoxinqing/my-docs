package app

import "fmt"

// defer 的执行顺序
// 1、多个 defer 语句，遵从后进先出(Last In First Out，LIFO)的原则，最后声明的 defer 语句，最先得到执行。
// 2、defer 在 return 语句之后执行，但在函数退出之前，defer 可以修改返回值。

func DeferTest01() int {
	i := 0
	defer func() {
		fmt.Println("defer1")
	}()
	defer func() {
		i += 1
		fmt.Println("defer2")
	}()
	return i
}

// defer2
// defer1
// return 0
// 返回值并没有被修改，这是由于 Go 的返回机制决定的，执行 return 语句后，Go 会创建一个临时变量保存返回值，因此，defer 语句修改了局部变量 i，并没有修改返回值

func DeferTest02() (i int) {
	i = 0
	defer func() {
		i += 1
		fmt.Println("defer2")
	}()
	return i
}

// defer2
// return 1
// 对于有名返回值的函数，执行 return 语句时，并不会再创建临时变量保存，因此，defer 语句修改了 i，即对返回值产生了影响

func DeferTest() {
	fmt.Println("return", DeferTest01())
	fmt.Println("return", DeferTest02())
}

func f(n int) {
	defer fmt.Println(n)
	n += 100
}

func TestF() {
	f(1) //打印 1 而不是 101。
	// defer 语句执行时，会将需要延迟调用的函数和参数保存起来，
	// 也就是说，执行到 defer 时，参数 n(此时等于1) 已经被保存了。因此后面对 n 的改动并不会影响延迟函数调用的结果。
}

// defer 的作用域是函数，而不是代码块，因此 if 语句退出时，defer 不会执行，而是等 101 打印后，整个函数返回时，才会执行。
func TestMain() {
	n := 1
	if n == 1 {
		defer fmt.Println(n)
		n += 100
	}
	fmt.Println(n)
}

// 101
// 1
