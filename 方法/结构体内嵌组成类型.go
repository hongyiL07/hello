package main

import (
	"fmt"
	"image/color"
	"sync"
)

func main()  {
	var cp ColoredPoint
	cp.x=1
	fmt.Println(cp.Point.x)
	cp.Point.y=2
	fmt.Println(cp.y)

	p:=ColoredPoint1{&Point{1,1},red}
	q:=ColoredPoint1{&Point{5,4},blue}
	fmt.Println(p.Point)        // 5
	q.Point = p.Point           // q 和 p 共享一个 Point
	p.ScaleBy(2)
	fmt.Println(*p.Point,*q.Point)
}

type Point struct {
	x,y float64
}
type ColoredPoint struct {
	Point
	Color  color.RGBA
}   //只想定义一个有三个字段的结构体 ColoredPoint ，实际上内嵌了一个 Point 类型以提供字段 x 和 y

// Point 的方法都被纳入到 ColoredPoint 类型中，内嵌允许构成复杂的类型，该类型由许多字段构成，每个字段提供一些方法

// ColoredPoint 并不是 Point ，但是它包含一个 Point ，并且它有两个另外的方法 Distance 和 ScaleBy 来自 Point

//匿名字段类型可以是个指向命名类型的指针，字段和方法间接地来自于所指向的对象
//让共享通用的结构体以及使对象之间的关系更加动态 多样化
// ColoredPoint1 声明内嵌了 *Point
type  ColoredPoint1 struct {
	*Point
	Color color.RGBA
}

//结构体可以拥有多个匿名字段

//方法只能在命名的类型（比如 Point）和指向它们的指针（*Point）中声明，但是内嵌帮助我们能够在未命名的结构体类型中声明方法
//简单的缓存实现，其中使用了两个包级别的级别的变量——————互斥锁 和 map，互斥锁将会保护 map 数据
var (
	mu sync.Mutex      //保护 mapping
	mapping = make(map[string]string)
)

func Lookup(key string) string {
	mu.Lock()
	v:= mapping[key]
	mu.Unlock()
	return v
}

//下面这个功能和上面完全相同，但是将两个相关变量放到了一个包级别的变量 cache 中
var cache = struct {
	sync.Mutex
	mapping map[string]string
}{
	mapping:make(map[string]string),
}

func Lookup1(key string) string {
	cache.Lock()
	v := cache.mapping[key]
	cache.Unlock()
	return  v
}   //新的变量名更加亲切，而且 sync.Mutex 是内嵌的，它的 Lock 和 Unlock 方法也包含进了结构体中，允许直接使用 cache 变量本身进行加锁