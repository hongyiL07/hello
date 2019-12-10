package main

import (
	"fmt"
	"math"
)

func main() {
	//浮点数常量 math.Pi 可用于任何需要浮点数或复数的地方
	var x float64 = math.Pi
	var y float32 = math.Pi
	var z complex128 = math.Pi
	fmt.Println(x, y, z)

	//假 始终需要 float32 或 complex128  则可能需要转换类型
	const Pi64 float64 = math.Pi

	var a float64 = Pi64
	var b float32 = float32(Pi64)
	var c complex128 = complex128(Pi64)
	fmt.Println(a, b, c)

	// true 和 false 是无类型布尔值，字符串字面量是无类型字符串

	//只有常量才可以是无类型的
	//若将无类型常量声明为变量 或者 类型明确的变量赋值的右边出现无类型常量，则常量会被隐式转换成该变量的类型
	//var f float64 = 3+0i
	//f = 2
	//f = le123
	//f = 'a'
	//等价于以下语句
	//var f float64= float64(3+0i)
	//f = float64(2)
	//f = float64(le123)
	//f = float64('a')

	//不论隐式或者显式，常量从一种类型转换成另一种，都要求目标类型能够表示原值。实数和复数允许舍入取整
	//const (
	//deadbeef = 0xdeadbeef //无类型整数，值为 3735928559
	//a=uint32(deadbeef)    //uint32 ,值为 3735928559
	//b=float32(deadbeef)   //float32 ,值为 3735928559(向上取整)
	//c=float64(deadbeef)   //float64 ，值为 3735928559（精确值）
	//d=int32(deadbeef)     //编译错误：溢出 ， int32 无法容纳常量值
	//e=float64(1e309)      //编译错误：溢出 ， float64 无法容纳常量值
	//f=uint(-1)            //编译错误：溢出 ， uint 无法容纳常量值

	//)

	//在将无类型转换为接口时，这些默认类型就分外重要，因为它们决定了接口值的动态类型
	fmt.Printf("%T\n", 0)      // ine
	fmt.Printf("%T\n", 0.0)    // float64
	fmt.Printf("%T\n", 0i)     // copmlex128
	fmt.Printf("%T\n", '\000') // int32  (rune)
}
