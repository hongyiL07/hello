package main

import (
	"fmt"
)

func main() {
	//内置函数 append 用来将元素追加到 slice 的后面

	//循环演示了如何使用 append 来为一个 rune 类型的 slice 添加元素
	var runes []rune
	for _, r := range "hello world" {
		runes = append(runes, r)
	}
	fmt.Println(runes)
	//最简单的方法是 []rune("hello world")
	fmt.Println([]rune("hello"))

	// slice 并不是纯引用类型，而是像下面这种聚合类型
	//type IntSlice struct{}
	//ptr *int
	//len ,cap int

	//appendInt只能给 slice 添加一个元素，但是内置的 append 函数可以同时给 slice 添加多个元素，甚至添加另一个 slice 里的所有元素
	var x []int
	x = append(x, 1)
	x = append(x, 2, 3)
	x = append(x, 4, 5, 6)
	x = append(x, x...) //追加 x 中的所有元素
	//省略号将 slice 转换为参数列表
	fmt.Println(x)

}

//为一个[]int数组 slice 定义的方法 appendInt
func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		// slice 仍有增长空间，扩展 slice 内容
		z = x[:zlen]
	} else {
		// slice 已无空间，为它分配一个新的底层数组
		//为了达到分摊线性复杂性， 容量扩展一倍
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x) //内置 copy 函数
	}
	z[len(x)] = y
	return z
	//输入参数 slice x 和函数返回值 slice z 拥有相同的底层数组
	//长度超出 容量翻倍
}

//简单修改 appendInt 函数来匹配 append 的功能
func appendInt1(x []int, y ...int) []int {
	var z []int
	zlen := len(x) + len(y) //扩展 slice z 的长度至少到 zlen
	copy(z[len(x):], y)
	fmt.Println(zlen)
	return z
}
