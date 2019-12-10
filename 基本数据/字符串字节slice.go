package main

import (
	"bytes"
	"fmt"
)

// 4个标准包对字符串操作：bytes  strings  strconv  unicode
// strings  搜索 替换 比较 修整 切分 连接字符串
// bytes    操作字节 slice （[]byte类型 其某些属性和字符串相同）
// strconv  主要用于 布尔值 整数 浮点数 与 字符串形式的转换
// unicode  判别文字符号值特性的函数（IsDigit IsLetter IsUpper IsLower）每个函数以单个文字符号值作为参数，返回布尔值

func main() {
	q := comma("12345")
	fmt.Println(q)

	for i, r := range "Hello,世界" {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}
	n := 0
	var s string = "456789"
	for _, _ = range s {
		n++

	}
	fmt.Println(s)
	//字节 slice 的元素允许随意修改
	//字符串和字节 slice 相互转换
	a := "abc"
	b := []byte(s)
	s2 := string(b)
	fmt.Println(a, b, s2)

	// bytes 包 提供的 Buffer 类型能高效处理字节 slice
	// bytes.Buffer变量无需初始化，零值本来就有效
	fmt.Println(intsToString([]int{1, 2, 3})) //[1,2,3]
}
func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

func intsToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}
