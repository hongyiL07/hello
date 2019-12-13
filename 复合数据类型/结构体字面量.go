package main

import "fmt"

type Point struct {
	x, y int
}

func main() {
	//结构体类型的值可以通过 结构体字面量 来设置，即通过设置结构体的成员变量来设置

	//要求按正确的顺序，为每个成员变量指定一个值
	p := Point{1, 2}
	fmt.Println(p)

	fmt.Println(Scale(Point{1, 2}, 5)) //{5,10}

	//创建  初始化一个 struct 类型的变量并获取它的地址
	//pp := &Point{1,2}
	//等价于
	//pp := new(Point)
	//*pp = Point{1,2}
	//但是 &Point{1，2} 这种复数可以直接使用在一个表达式中

	//《《《《《《《《《结构体比较》》》》》》》》》》》》》
	//如果结构体的所有成员变量都可以比较，那这个结构体就是可比较的（ ==   != ）
	//其中 == 操作符按照顺序比较两个结构体变量的成员变量，所以下面两个输出语句是dengjiade
	type qwer struct {
		x, y int
	}
	a := qwer{1, 2}
	q := qwer{2, 1}
	fmt.Println(p.x == q.x && p.y == q.y) //false
	fmt.Println(a == q)                   //false

	//和其他可比较的类型一样，可比较的结构体类型都可以作为 map 的键类型
	type address struct {
		hostname string
		port     int
	}
	hits := make(map[address]int)
	hits[address{"golang.org", 443}]++
}

//结构体类型的值可以作为参数传递给函数或者作为函数的返回值
//将 Point 缩放一个比例
func Scale(p Point, factor int) Point {
	return Point{p.x * factor, p.y * factor}
}

//出于效率，大型的结构体通常使用结构体指针的方式直接传递给函数或者从函数中返回
//func Bonus(e *Employee, percent int) int {
//	e.Salary = e.Salary *105/100
//}
