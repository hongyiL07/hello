package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println("select多路复用")

	//火箭发射进行倒计时
	fmt.Println("Commencing countdown.")
	tick := time.Tick(1 * time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		<-tick
	}

	//通过按下回车键来取消发射过程，第一，启动一个 goroutine 试图从标准输入中读取一个字符，若成功，发生一个值到 abort 通道
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1)) //读取单个字节
		abort <- struct{}{}
	}()

	//现在每一个倒计时迭代需要等待事件到达两个通道中的一个：计时器通道，前提是一切顺利；或者中止事件前提是有“异常”
	//不能只从一个通道上接收，因为哪一个操作都会在完成前堵塞，
	//所以需要多路复用那些操作过程，为了实现这个目的，需要一个 select 语句
	//通用格式
	//select {
	//case <-ch1:
	//case x =: <- ch2:
	//case ch3<-y:
	//default：

	//创建中止通道
	fmt.Println("Commencing countdown.  Press return to abort.")
	select {
	case <-time.After(10 * time.Second):
	//不执行任何操作
	case <-abort:
		fmt.Println("Launch abortted!")
		return

	}
	launch()

	//偶数时发送
	ch := make(chan int, 1)
	for i := 0; i < 10; i++ {
		select {
		case x := <-ch:
			fmt.Println(x)
		case ch <- i:

		}
	}

	//有时候试图在一个通道上发送或接收，但是不想在通道没有准备好的情况下被阻塞————非阻塞通信，使用 select语句能做到
	// select 可以有一个默认的情况，它用来指定在没有其他的通信发生时可以立即执行的操作
}
