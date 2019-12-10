package main

import (
	"fmt"
)

const boilingF = 212.0 //声明包级别的常量
func main() {
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("boiling point = %g℉  or %g℃\n", f, c)
	//输出
	//"boiling point =  212℉  or 100℃

	const freezingF, bo = 32.0, 212.0
	fmt.Printf("%g℉  = %g℃\n", freezingF, fToC(freezingF))
	fmt.Printf("%g℉ = %g℃\n", bo, fToC(bo))
	//fToC(freezingF)   bo,fToC(bo)调用 fToC 函数

	//var q,err  =os.Open(name)
	//返回一个文件和一个错误

	//短变量声明  （最少声明一个新变量  否则 代码编译将无法通过）
	i, j := 0, 1
	i, j = j, i //交换i  j 的值

}

func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}
