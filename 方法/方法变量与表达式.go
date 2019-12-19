package main

import (
	"fmt"
	"math"
)

func main()  {
	//通常都在相同的表达式里使用和调用方法，就像在 p.Diatance（）中，但是把两个操作分开也是可以的。
	//选择子 p.Diatance 可以赋予一个方法变量，它是一个函数，把方法（p.Diatance ）绑定到一个接收者 p 上。
	//函数只需要提供实参，而不需要提供接收者就能够调用
	p:=Point{1,2}
	q:=Point{4,6}
	distanceFromP:=p.Diatance            //方法变量
	fmt.Println(distanceFromP(q))        // 5
	var origin Point                     // { 0 ， 0}
	fmt.Println(distanceFromP(origin))
	scaleP := p.ScaleBy                  //方法变量
	scaleP(2)
	scaleP(3)
	scaleP(10)

	//与方法变量相关的是方法表达式，和调用一个普通的函数不同，在调用方法的时候必须提供接收者，并且按照选择子的语法进行调用
	//方法表达式写成 T.f 或者 (*T).f ,其中 T 是类型，是一种函数变量，把原来方法的接收者替换成函数的第一个形参，因此它可以像平常的函数一样调用
	distance := Point.Diatance          //方法表达式
	fmt.Println(distance(p,q))          // 5
	fmt.Println(distance)               // func(Point ,Point) float64

	scale := Point.ScaleBy
	scale(&p,2)                         // func(*Point , float64)
	fmt.Println(p)
	fmt.Println(scale)


}
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
func (p *Point) ScaleBy(factor float64)  {
	p.x *=factor
	p.y *=factor
}

//如果需要用一个值代表多个方法中的一个而方法都属于同一个类型，方法变量可调用这个值所对应的方法来处理不同的接收者
//变量 op 代表加法或者减法，二者都属于 Point 类型的方法，Path.TranslateBy 调用了它计算路径上的每一个点
func (p Point) Add (q Point) Point {
	return Point{p.x+q.x,p.y+q.y}
}
func (p Point) Sub (q Point) Point {
	return Point{p.x-q.x,p.y-q.y}
}

type Path []Point

func (path Path) TranslateBy(offset Point,add bool) {
	var op func(p,q Point) Point
	if add{
		op = Point.Add
	}else {
		op = Point.Sub
	}
	for i:= range path{
		//调用 path[ i ].Add(offset) 或者 path[ i ].Sub(offset)
		path[i] = op(path[i] , offset)
	}
}
