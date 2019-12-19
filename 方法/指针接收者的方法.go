package main

import "fmt"

func main()  {
	//通过提供 *Point 能够调用 （*Point）.ScaleBy ） 方法
	r := &Point{1,2}
	r.ScaleBy(2)
	fmt.Println(*r)
}

//由于主调函数会复制每一个实参变量，如果函数需要更新一个变量或者如果一个实参太大而我们希望避免复制整个实参
//因此我们必须使用指针来传递变量的地址，同样适用于更新接收者：我们将它绑定到指针类型
func (p *Point) ScaleBy(factor float64)  {
	p.x *=factor
	p.y *=factor
}   //这个方法名字是 （*Point）.ScaleBy ）。圆括号是必须的，没有圆括号，表达式会解析为 *（Point.ScaleBy）

//不允许本身是指针的类型进行方法声明

//实参接收者是 *Point 类型，以 Point.ScaleBy 的方法调用 Point 类型的方法是合法的因为我们有办法从地址中获取 Point 的值
//只要解引用指向接收者的指针即可，编译器自动插入一个隐式的 * 操作符
//pptr.Distance(q)
//(*pptr).Diatance    调用效果一样


// nil 是一个合法的接收者
//就像一些函数允许 nil 指针作为实参，方法的接收者也一样，尤其是当 nil 是类型中有意义的零值（如 map 和 slice 类型）时，更是如此
//在简单的整形数链表中， nil 代表空链表
//
// IntList 是整形链表
// *IntList 的类型 nil 代表空列表
type IntList struct {
	Value int
	Tail  *IntList
}
//
// Sum 返回列表元素的总和
func (list *IntList) Sum() int {
	if list == nil{
		return 0
	}
	return list.Value + list.Tail.Sum()
}   //当定义一个类型允许 nil 作为接收者时，应当在文档注释中显示的标明
