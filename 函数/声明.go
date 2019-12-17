package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(hypot(3, 4))
}

// x,y 是函数声明中的形参
func hypot(x, y float64) float64 {
	return math.Sqrt(x*x + y*y)
}

//函数的类型称作 函数签名
//当两个函数拥有相同的形参列表和返回列表时，认为这两个函数的类型和签名是相同的

//形参变量都是函数的局部变量，初始值由调用者提供的实参传递
//函数形参以及命名返回值同属于函数最外层作用域的局部变量

//实参是按值传递的，函数接收到的是每个实参的副本，修改函数的形参变量并不会影响到调用者提供的实参

//如果提供的实参包含引用类型，如指针，slice，map，函数，通道，那么当函数使用形参变量时就有可能会间接的修改实参变量
