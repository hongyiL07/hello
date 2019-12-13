package main

func main() {
	//结构体嵌套机制： 将一个命名结构体当作另一个结构体类型的 匿名成员 使用
	//并提供一种方便的语法，使用简单的表达式（x.y）就可以代表连续的成员（x.e.d.y）
	type Circle struct {
		x, y, Radius int
	}
	type Wheel struct {
		x, y, Radius, Spokes int
	}
	var w Wheel
	w.x = 8
	w.y = 8
	w.Radius = 5
	w.Spokes = 20
	//在支持的类型变多时，意识到它们之间的相似性和重复性。所有要重构相同的部分
	type Point struct {
		x, y int
	}
	//type Circle struct {
	//	Center Point
	//	Radius int
	//}
	//type Wheel struct {
	//	Circle Circle
	//	Spokes int
	//}
	//这个看上去变得清晰，但是访问 Wheel 的成员变麻烦了
	//var w Wheel
	//w.Circle.Center.x = 8
	//w.Circle.Center.y = 8
	//w.Circle.Radius = 8

	// GO 允许定义不带名称的结构体成员，只需要指定类型即可，叫做 匿名成员
	//这个结构体成员的类型必须是一个命名类型或者指向命名类型的指针
	// qw 和 as 都拥有一个匿名成员 ，称 Point 被嵌套到 qw 中，qw 被嵌套到 as 中
	type qw struct {
		Point
		Radie int
	}
	type as struct {
		Circle
		Spokes int
	}
	//有了这种结构体嵌套功能，直接访问到我们需要的变量而不是指定一大串中间变量
	var z as
	z.x = 8      //等价于 w.Circle.Point.x = 8
	z.y = 8      //等价于 w.Circle.Point.y = 8
	z.Radius = 5 //等价于 w.Circle.Radius = 5
	z.Spokes = 20

	//结构体字面量并没有什么快捷方式来初始化结构体

	//结构体字面量必须遵循形状类型的定义，以下 两种 等价的方式初始化
	//w = Wheel{Circle{Point{8,8},5},20}
	//w = Wheel{Circle: Circle{
	//	Point:Point{x:8,y:8},
	//	Radius:5,
	//},Spokes:20}
	//fmt.Println(w)

	//即使两个结构体是不可导出的，仍可以使用快捷方式
}
