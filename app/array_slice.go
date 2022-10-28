package app

import "fmt"

//
func PrintSlice() {
	a := []int{1, 2, 3}
	for k, v := range a {
		if k == 0 {
			a[0], a[1] = 100, 200
			fmt.Print(a)
		}
		a[k] = 100 + v
	}
	fmt.Print(a)
}

// [100 200 3][101 300 103]

//
func PrintArray() {
	a := [3]int{1, 2, 3}
	for k, v := range a {
		if k == 0 {
			a[0], a[1] = 100, 200
			fmt.Print(a)
		}
		a[k] = 100 + v
	}
	fmt.Print(a)
}

// [100 200 3][101 102 103]

// 总结：
// slice是一个指向数组的指针，那么修改slice的时候会改变数组的值；
// 数组传的是值，不会改变元素的值；
// 虽然slice也是值语义，但是其本身是指针类型的，所以会改变值，而不意味这slice传引用。

// slice支持append操作， 数组不支持
// 数组也可以进行切片，返回值是slice，改变slice时会同步修改数组内容，相当于取得这个数组的指针
// 数组赋值时会拷贝内容，slice会赋值引用和指针一样
// 数组定长，定义的时候就需要确定。切片长度不定，append时会自动扩容
