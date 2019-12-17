package main

import (
	"fmt"
	"strings"
)

//函数再 Go 语言中是头等重要的值：函数变量也有类型，而且它们可以赋给变量或者传递或者从其他函数中返回。
//函数变量可以像其他函数一样调用

//函数类型的零值是 nil（空值），调用一个空的函数变量报错
func main() {
	f := square
	fmt.Println(f(3))

	//函数变量可以和空值相比较
	if f != nil {
		fmt.Println(f(4))
	}

	//它们本身不可比较，所以不可互相进行比较或者作为键值出现在 map 中

	//函数变量使得函数不仅将数据进行参数化，将函数的行为当作参数进行传递
	fmt.Println(strings.Map(add1, "HAL-9000"))
	fmt.Println(strings.Map(add1, "VMS"))
	fmt.Println(strings.Map(add1, "Admix"))
}
func square(n int) int {
	return n * n
}
func negative(n int) int {
	return -n
}
func product(m, n int) int {
	return m * n
}

//函数变量使得函数不仅将数据进行参数化，将函数的行为当作参数进行传递
func add1(r rune) rune {
	return r + 1
}

//使用函数变量，可以使得我们将每个节点的操作逻辑从遍历树形结构的逻辑中分开
//func forEachNode(n *html.Node,pre,post func(n *html.Node))  {
//	if pre != nil{
//		pre(n)
//	}
//	for  c:= n.FirstChild; c != nil;c=c.NextSibling{
//		forEachNode(c,pre,post)
//	}
//	if post!= nil{
//		post(n)
//	}
//}         forEachNode函数接收两个函数作为参数，一个在本节点所有子节点都被访问前调用
