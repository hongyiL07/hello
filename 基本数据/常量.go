package main

import "fmt"

func main() {
	const pi = 3.1415926 //近似数；math.Pi 是更精确的近似

	//定义多个常量
	const (
		e = 2566
		r = 856665
	)

	//len cap real imag complex unsafe.Sizeof 返回的都是常量

	//声明一组常量，除第一项外等号右边表达式可以省略
	const (
		a = 1
		b
		c = 2
		d
	)
	fmt.Println(a, b, c, d) //1  1  2  2

	//常量生成器  iota

}

//常量生成器  iota  （枚举类型） iota 机制存在局限
type Weekday int

const (
	Sunday Weekday = iota // sunday值为 0
	Monday                // Monday值为 1  以此类推
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)
