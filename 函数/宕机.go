package main

import (
	"fmt"
	"os"
	"runtime"
)

// GO 语言的类型系统会捕获许多编译时的错误，但有些其他的错误（比如数组越界访问或者引用空指针）都需要在运行时进行检查。当 Go 语言运行时检测到这些错误，它就会发生宕机
//一个典型的宕机发生时，正常的程序执行会终止， goroutine 中的所有延迟函数会执行，然后程序会异常退出并留下一条日志信息
//日志消息包括宕机的值，这往往代表某种错误消息，每个 goroutine 都会在宕机的时候显示一个函数调用的 栈跟踪 消息
//并不是所有宕机都是在运行时发生的
func main() {
	fmt.Println("444")

	defer printStack()
	f(3)
}

//当宕机发生时，所有的延迟函数以倒叙执行，从栈的最上面函数开始一直返回至 main 函数
func f(x int) {
	fmt.Println(x + 2/x) // if x == 0. 则发生宕机
	defer fmt.Println(x)
	f(x - 1)
}

//函数是可以从宕机状态恢复至正常运行状态而不让程序退出
// runtime 包提供了转储栈的方法使程序员可以诊断错误
func printStack() {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false) //代码在 main 函数中延迟 printStack 的执行
	os.Stdout.Write(buf[:n])
}

//设置函数的断言是个良好的习惯，但是这样也会带来多余的检查
//除非你能够提供有效的错误消息或者能够很快的检查出错误，否者在运行时检测断言条件就毫无意义
//func Reset(x *Buffer)  {
//	if x== nil{
//		panic("x is nil")   //没必要
//	}
//	x.elements = nil
//}
