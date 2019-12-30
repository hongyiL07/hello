package main

import "fmt"

//通道可以用来连接 goroutine，这样一个的输出是另一个的输入。就叫 管道（pipeline）

func main() {
	fmt.Println("管道")

	naturals := make(chan int)
	squares := make(chan int)
	//counter
	go func() {
		for x := 0; x < 100; x++ {
			naturals <- x
		}
		close(naturals)
	}()
	//squares
	go func() {
		for {
			x, ok := <-naturals
			if !ok {
				break //通道关闭并且读完
			}
			squares <- x * x
		}
		close(squares)
	}()
	//printer(在主 goroutine 中)
	for x := range squares {
		fmt.Println(x)
	} //结束时，关闭每个通道不是必须的。只有在通知接收方 goroutine 所有的数据都发送完毕的时候才需要关闭通道

}

// Go 的类型系统提供了单向通道类型，仅仅导出发送或者接收操作。
// 类型 chan<- int 是一个只能发送的通道允许发送但不允许接收
//类型 <- chan int 是一个只能接收的 int 类型通道，允许接收但不允许发送

//缓冲通道有一个元素队列，队列的最大长度在创建的时候通过 make 的容量参数来设置
//ch =make(chan string , 3)
//程序需要知道缓冲通道的容量，可以通过调用内置的 cap 函数获取它
// fmt.Println(cap(ch))    //  3

//如果使用一个无缓冲通道，两个比较慢的 goroutine 将被卡住，因为在它们发送响应结果到通道的时候没有 goroutine 来接收
//这种情况叫做 goroutine 泄漏，它属于一个 bug
//无缓冲和缓冲通道的选择  缓冲通道容量大小的选择，都会对程序的正确性产生影响
