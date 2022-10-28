package app

// * 运算符，也称为解引用运算符，用于访问地址中的值。
// ＆ 运算符，也称为地址运算符，用于返回变量的地址。
func foo() *int {
	v := 11
	return &v
}

func main() {
	m := foo()
	println(*m) // 11
}
