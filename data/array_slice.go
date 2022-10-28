package data

import "fmt"

func PrintArraySlice() {
	nums1 := [3]int{1, 2, 3}
	nums2 := nums1
	nums2[0] = 10
	fmt.Println(nums1[0]) // 数组不变，还是1

	nums3 := []int{1, 2, 3}
	nums4 := nums3
	nums4[0] = 10
	fmt.Println(nums3[0]) // 切片是引用类型，值变 10

	// 同样是用等号进行赋值，对于数组来说进行的是深拷贝（值拷贝），而切片则是浅拷贝（指针拷贝）。
	// 因此对nums2的赋值改变的是nums2地址空间的值，而对nums4的赋值修改的是nums3和nums4共用的内存地址。
}

// 知识点
// 切片
// 容量 cap 是指为该切片准备了cap大小的内存空间，当切片中的数据数量不超过cap时，切片是不需要进行扩容的。
// 长度 len 表示的是切片中元素数量有 len 个，主要应用于切片的初始化。

// 扩容策略
// 数组作为基本数据类型，数组的长度也是类型的一部分，所以数组每次扩容都需要新建一个数组做值拷贝。

// 切片作为包装类型，如果在切片中添加元素后，切片长度未超过cap，只需在slice.array指向的内存区域进行赋值。如果切片长度超过了容量cap，扩容需要修改slice.array指向的内存大小，再进行值拷贝

// 切片的扩容是基于俩个阀值去判断的，一个是当前容量的俩倍，一个是扩容前切片的容量是否大于1024

// 切片拷贝导致的内存泄漏
// 在做切片的拷贝时，如果只是从大切片上获取很小的一部分使用，应该使用copy进行拷贝，而不是用等号进行赋值，避免大切片由于这一块的引用一直得不到释放，从而导致内存泄漏。

// slice是怎么扩容的？
// 如果当前容量小于1024，则判断所需容量是否大于原来容量2倍，如果大于，当前容量加上所需容量；否则当前容量乘2。
// 如果当前容量大于1024，则每次按照1.25倍速度递增容量，也就是每次加上cap/4。
