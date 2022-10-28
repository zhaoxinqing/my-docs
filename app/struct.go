package app

import "fmt"

// 空 struct{} 的用途
// 1、使用空结构体 struct{} 可以节省内存，一般作为占位符使用，表明这里并不需要一个值。

// fmt.Println(unsafe.Sizeof(struct{}{})) // 0
// 比如使用 map 表示集合时，只关注 key，value 可以使用 struct{} 作为占位符。如果使用其他类型作为占位符，例如 int，bool，不仅浪费了内存，而且容易引起歧义。

type Set map[string]struct{}

func StructTest() {
	set := make(Set)

	for _, item := range []string{"A", "A", "B", "C"} {
		set[item] = struct{}{}
	}
	fmt.Println(len(set)) // 3
	if _, ok := set["A"]; ok {
		fmt.Println("A exists") // A exists
	}
}

// 2、再比如，使用信道(channel)控制并发时，我们只是需要一个信号，但并不需要传递值，这个时候，也可以使用 struct{} 代替。

func StructChannel() {
	ch := make(chan struct{}, 1)
	go func() {
		<-ch
		// do something
	}()
	ch <- struct{}{}
	// ...
}

// 3、再比如，声明只包含方法的结构体。

type Lamp struct{}

func (l Lamp) On() {
	println("On")

}
func (l Lamp) Off() {
	println("Off")
}
