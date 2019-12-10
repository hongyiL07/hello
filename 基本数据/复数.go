package main

import "fmt"

func main() {
	//两种大小的复数 complex64 和 complex128，分别由 float32 和 float64 构成
	//complex 函数根据给定的实部和虚部创建复数
	//内置的 real 函数 和 image 函数则分别提取复数的实部和虚部
	var x complex128 = complex(1, 2) // 1+2i   等价于  2i+1
	var y complex128 = complex(3, 4) // 3+4i
	fmt.Println(x * y)               // (-5+10i)
	fmt.Println(real(x * y))         // -5
	fmt.Println(imag(x * y))         // 10

	fmt.Println(1i * 1i) //i^2=-1  (-1+0i)
}
