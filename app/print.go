package app

import "fmt"

type Stu struct {
	Name string
}

//
func PrintStruct() {
	// %v 和 %+v 都可以用来打印 struct 的值，区别在于 %v 仅打印各个字段的值，%+v 还会打印各个字段的名称。
	fmt.Printf("%v\n", Stu{"Tom"})  // {Tom}
	fmt.Printf("%+v\n", Stu{"Tom"}) // {Name:Tom}
}
