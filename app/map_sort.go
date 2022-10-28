package app

import (
	"fmt"
	"sort"
)

// Golang中的map默认是无序的，每次遍历，得到的输出结果可能不一样。

// Golang中的map排序：
// 将map的key放到切片中；
// 对切片排序；
// 遍历切片，然后来按key来输出map的值。

// 根据Key排序
func sortMapByKey(mp map[string]int) {
	// 1、将map的key放到切片中
	var newMap = make([]string, 0)
	for k, _ := range mp {
		newMap = append(newMap, k)
	}
	// 2、对切片排序
	sort.Strings(newMap)
	// 3、遍历切片，然后按key来输出map的值
	for _, v := range newMap {
		fmt.Printf("根据key排序后的新集合为:[%v]=%v \n", v, mp[v])
	}
}

// 根据value排序
func sortMapByValue(mp map[string]int) {
	var newMp = make([]int, 0)
	var newMpKey = make([]string, 0)
	for oldk, v := range mp {
		newMp = append(newMp, v)
		newMpKey = append(newMpKey, oldk)
	}
	sort.Ints(newMp)
	for k, v := range newMp {
		fmt.Printf("根据value排序后的新集合为:[%v]=%v \n", newMpKey[k], v)
	}
}

// 默认增序
func SortMap() {

	map1 := make(map[string]int)
	map1["3"] = 11
	map1["5"] = 5
	map1["1"] = 7
	map1["4"] = 23
	map1["2"] = 42

	sortMapByKey(map1)
	fmt.Println()
	sortMapByValue(map1)
}
