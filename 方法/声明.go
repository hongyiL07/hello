package geometry

import (
	"fmt"
	"math"
)

//方法的声明和普通函数的声明类似，只是在函数名字前面多了个参数，这个参数把这个方法绑定到这个参数对应的类型上

type Point struct {
	x, y float64
}

//普通函数
func Distance(p,q Point) float64 {
	return math.Hypot(q.x-p.x,q.y-p.y)
}

// Point 类型的方法
func (p Point) Diatance(q Point) float64 {
	return  math.Hypot(q.x-p.x,q.y-p.y)
}
//附加的参数 p 称为方法的接收者，它源自最先的面向对象语言，用来描述主调方法就像向对象发送消息
func main()  {
	p := Point{1,2}
	q := Point{4,6}
	fmt.Println(Distance(p,q))    //  “ 5 ” 函数调用
	fmt.Println(p.Diatance(q))    // “ 5 ”  方法调用
}
//上面两个 Diatance 函数声明没有冲突
//第一个声明一个包级别的函数（称为 geometry.Diatance）
//第二个声明一个类型 Point 的方法，名字为 Point.Diatance
//表达式 p.Diatance 称作 选择子，它为接收者 p 选择合适的 Diatance 方法

//每个类型有它自己的命名空间，能够在其他不同的类型中使用名字 Diatance 作为方法名
//定义一个 Path 类型表示一条线段
type Path  []Point
func (path  Path) Diatance() float64 {
	sum := 0.0
	for i := range path{
		if i >0{
			sum += path[i-1].Diatance(path[i])
		}
	}
	return sum
}
// Path 是一个命名的 slice 类型，而非 Point 这样的结构体类型，但依旧可以给它定义方法

//类型拥有的所有方法名都必须是唯一的，但是不同的类型可以使用相同的方法名，没必要使用附加的字段来修饰方法名

