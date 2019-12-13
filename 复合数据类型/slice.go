package main

import (
	"fmt"
)

func main() {
	// slice 表示一个拥有相同类型元素的可变长度的序列  通常写成 []T 其中元素类型都是 T 型

	// slice 是一种轻量级的数据结构，可以用来访问数组的部分或者全部的元素，这个数组称为 slice 的底层数组

	// slice 有三种属性 （指针  长度  容量）

	//月份为例的数组声明
	months := [...]string{1: "January", 2: "February", 3: "March", 4: "April", 5: "May", 6: "June", 7: "July", 8: "August", 9: "September", 10: "October", 11: "November", 12: "December"}
	Q2 := months[4:7]
	summer := months[6:9]
	fmt.Println(Q2)
	fmt.Println(summer)

	//输出 Q2 和 summer 的共同元素
	for _, s := range summer {
		for _, q := range Q2 {
			if s == q {
				fmt.Println(s)
			}
		}
	}

	// cap(s) slice 的引用超过了被引用对象的容量
	// len(s) slice 的引用超过了被引用对象的长度
	//都会导致最终的 slice 比原来的 slice 长
	//fmt.Println(summer[:20])   //超过被引用对象的边界
	endlessSummer := summer[:5] //在 slice 容量范围内扩展了 slice
	fmt.Println(endlessSummer)

	a := [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println(a)
	reverse(a[:])
	fmt.Println(a)

	//将一个 slice 左(右)移 n 个元素
	reverse(a[:2])
	fmt.Println(a)
	reverse(a[2:])
	fmt.Println(a)

	//引用类型 ，例如（指针 通道）操作符 == 检查的是 引用相等性 ，即它们是否指向相同的元素

	//slice 唯一允许的比较操作是和 nil 作比较

	//slice 类型的零值是 nil ，值为 nil 的 slice 没有对应的底层数组 有些非 nil 的 slice 长度和容量是零
	//检查一个 slice 是否为空，使用 len（s）== 0，而不是 s == nil
	var s []int    // len(s)==0,s==nil
	s = nil        // len(s)==0,s==nil
	s = []int(nil) // len(s)==0,s==nil
	s = []int{}    // len(s)==0,s!=nil
	fmt.Println(len(s))

	//内置函数 make 可以创建一个具有指定 元素类型 长度 容量 的 slice ，其中容量参数可以省略，在这种情况下， slice的长度和容量相等
	//make([]T,len)
	//make([]T,len,cap)  //和 make([]T,cap)[:len] 功能相同

}

//反转整形 slice 中的元素，适用任意长度的整形 slice
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// slice  无法做比较，标准库中提供了高度优化的函数 bytes.Equal 来比较两个字节 slice （[] byte）
//对于其他类型的 slice ，我们必须自己写函数来比较
func equal(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}
