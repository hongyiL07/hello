package main

import (
	"fmt"
)

func main() {
	//这种情况下，底层数组的元素只是部分被修改
	data := []string{"one", "", "three"}
	fmt.Println(nonempty(data))
	fmt.Println(data)

	// slice 可以用来实现栈
	//给定一个空的 slice 元素 stack，可以用 append 向 slice 尾部追加值
	//stack = append(stack,v)   push v
	//栈的顶部是最后一个元素
	//top := stack[len(stack)-1]    //顶部
	//通过弹出最后一个元素来缩减栈
	//stack = stack[:len(stack)-1]   // pop

	s := []int{5, 6, 7, 8, 9}
	fmt.Println(remove(s, 2))

	f := []int{5, 6, 7, 8, 9}
	fmt.Println(remove1(f, 2))

}

// Nonempty 演示了 slice  的就地修改算法
//nonempty 返回一个新的 slice ， slice 中的元素都是非空字符串
//在函数的调用过程中，底层数组的元素发生了改变
func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
	//输入的 slice 和输出的 slice 拥有相同的底层数组，避免函数内部重新分配一个数组
}

//函数 nonempty 还可以利用 append 函数来写
func nonempty2(strings []string) []string {
	out := strings[:0]
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}

//为了从 slice 的中间移除一个元素，并保留剩下的元素顺序，可以使用函数 copy 来将高位索引的元素向前面移动来覆盖被移除元素的所在位置
func remove(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

//如果不需要维持 slice 中剩余元素的顺序，可以简单的将 slice 的最后一个元素赋值给被移除元素所在的索引位置
func remove1(slice []int, i int) []int {
	slice[i] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}
