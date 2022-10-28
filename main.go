package main

import (
	"fmt"
	"math/rand"
	"time"
)

// func main() {
// 	app.BinaryTreeTest()
// }

func InsertSort(st []int) []int {
	if len(st) <= 1 {
		return st
	}
	for i := 1; i < len(st); i++ {
		back := st[i]
		j := i - 1
		for j >= 0 && back < st[j] {
			st[j+1] = st[j]
			j--
		}
		st[j+1] = back
	}
	return st
}

//
func main() {

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var list []int
	for i := 0; i < 100; i++ {
		list = append(list, r.Intn(100))
	}
	fmt.Println(InsertSort(list))
}
