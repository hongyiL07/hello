package main

import (
	"fmt"
	"math"
)

func main() {

	// 溢出：运算结果所需的位超出该类型的范围
	var u uint8 = 255
	fmt.Println(u, u+1, u*u) //255  0  1
	var i uint8 = 127
	fmt.Println(i, i+1, i*i) //127  -128  1

	//谓词指定进位制 输出格式
	o := 0666
	q := "时间"
	fmt.Printf("%d %[1]o %#[1]o\n", o) //438  666   0666
	x := int64(0xdeadbeef)
	fmt.Printf("%d %#[1]x %#[1]X\n", x) //3735928559  deadbeef 0xdeadbeef   0Xdeadbeef

	// %c 输出文字符号  加单引号用  %q
	fmt.Printf("%q\n", q)

	//输出 e 的幂方，保留三位小数
	for x := 0; x < 8; x++ {
		fmt.Printf("x = %d e^x = %8.3f\n", x, math.Exp(float64(x)))
	}

	//math.IsNaN 函数判断其参数是否是非数值
	nan := math.NaN()
	fmt.Println(nan == nan, nan < nan, nan > nan) //false   false   false

}
