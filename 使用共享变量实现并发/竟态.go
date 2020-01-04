package main

import (
	"fmt"
	bank "github.com/hongyiL07/hello/使用共享变量实现并发"
	"image"
)

//定义：数据竟态发生于两个 goroutine 并发读取同一个变量并且至少其中一个是写入时

//无法自信的说一个事件肯定先于另一个事件，那么这两个事件就是并发的
//这个函数在并发调用时仍然能正确工作，那么这个函数是并发安全的
//一个类型的所有可访问方法和操作都是并发安全时，则可称为并发安全的类型

func main()  {
	fmt.Println("竟态")

	//竟态是指在多个 goroutine 按某些交错顺序执行时程序无法给出正确的结果
	//竟态对于程序是致命的，因为它们可能会潜伏在程序中，出现频率也很低，
	//有可能仅在高负载环境或者在使用特定的编译器 平台 架构时才出现。这些都让竟态很难再现和分析

	go func() {
		bank.Deposit(200)            //A1
		fmt.Println("=",bank.Balance())  //A2
	}()
	go bank.Deposit(100)   //B
}

//避免竟态的第一种方法
//不要修改变量
var icons = make(map[string]image.Image)

func loadIcon(name string) image.Image {
	icon,ok:=icons[name]
	if !ok{
		icon=loadIcon(name)
		icons[name]=icon
	}
	return icon
}//如果在创建其他 goroutine 之前就用完整的数据来初始化 map ，并且不再修改，
// 那么无论再多的 goroutine 也可以安全的并发调用 Icon，因为每个 goroutine 都只读取这个 map

//第二种是避免从多个 goroutine 访问同一个变量
type Cake struct {
	state string
}

func baker(cooked chan <- *Cake)  {
	for{
		cake := new(Cake)
		cake.state="cooked"
		cooked <- cake   // 不再访问 cake 变量
	}
}
func icer(iced chan<- *Cake,cooked <-chan *Cake)  {
	for cake:=range cooked{
		cake.state="iced"
		iced<-cake   //iced 不再访问 cake 变量
	}
}

//第三种方法就是 允许多个 goroutine 同时访问同一个变量，但同一时间只有一个 goroutine 可以访问（互斥机制）


