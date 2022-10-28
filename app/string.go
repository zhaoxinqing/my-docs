package app

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// string 转 byte[]

// []byte 转 string

// byte 是 uint8 的别名，其字面量是 8 位整数值，它可以更改每个字节或字符
// rune 是 int32 的别名,其字面量是 32 位整数

func ByteString() {
	s := "hello,zhaoxinqing"
	b := []byte(s)

	// for _, v := range s {
	// 	fmt.Println(string(v))
	// }

	// for _, v := range b {
	// 	fmt.Println(string(v))
	// }

	// 20 > 123456789 ,

	fmt.Print(string(b[1:6]))
}

// 如何高效地拼接字符串?
// Go 语言中，字符串是只读的，也就意味着每次修改操作都会创建一个新的字符串。如果需要拼接多次，应使用 strings.Builder，最小化内存拷贝次数。
func MergeString() {
	var str strings.Builder
	for i := 0; i < 100; i++ {
		str.WriteString("a")
	}
	fmt.Println(str.String())
}

// 统计字符串长度
func CountStrLens() {
	// string底层是通过byte数组实现的。中文字符在unicode下占2个字节，在utf-8编码下占3个字节，而golang默认编码是utf-8
	var str = "hello 你好"
	fmt.Println(len(str)) // 12

	// 获取字符串的长度，而不是字符串底层占得字节长度
	fmt.Println("RuneCountInString:", utf8.RuneCountInString(str)) // RuneCountInString: 8

	//截取中文字符串时，首先需要将字符串转换成 rune 数组
	str0 := "嗨客网(Hello HaiCoder!)"
	str1 := str[0:2]
	strC := []rune(str0)
	str2 := strC[0:2]
	fmt.Println("str1 =", string(str1), "str2 =", string(str2)) // str1 = he str2 = 嗨客
}

func GetChineseCharacter() {

}
