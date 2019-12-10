package main

import (
	"fmt"
)

func main() {
	p := new(int)   //*int 类型的 p ，指向未命名的 int 变量 ， new（int） 直接在表达式中使用
	fmt.Println(*p) //输出 0
	*p = 2          //把未命名的 int 设置成 2
	fmt.Println(*p) //输出 2

	//每次调用 new 返回一个具有唯一地址的不同变量
	q := new(int)
	w := new(int)
	fmt.Println(q == w) //false

	z := gcd(8, 6)
	fmt.Println(z)

	c := fib(8)
	fmt.Println(c)

}
func gcd(x, y int) int { //计算两个数的最大公约数
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

func fib(n int) int { //计算斐波那契的第 8 数
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}
