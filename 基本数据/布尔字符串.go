package main

import "fmt"

func main() {
	// && 比较 || 优先级更高  分别表示逻辑乘法  逻辑加法
	//布尔值无法隐式转换成数值（如 0 或 1 ）
	s := "hello, world"
	fmt.Println(len(s))     //输出 s 的长度
	fmt.Println(s[0], s[7]) //  h  w

	fmt.Println(s[:5]) //截取字符串
	fmt.Println(s[0:5])
	fmt.Println(s[7:])
	fmt.Println(s[:])

	//加号 + 连接两个字符串生成一个新字符串
	fmt.Println("good" + s[5:])

	//字符串不可改变，内部数据不可修改
}
func btoi(b bool) int {
	// b 为真，btoi 返回 1 ；
	// b 为假，btoi 返回 0 ；
	if b {
		return 1
	}
	return 0
}

func itob(i int) bool {
	//报告 i 是否为非零值
	return i != 0
}
