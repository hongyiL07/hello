package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("取消")

	//有时候需要让一个 goroutine 停止它当前的工作
	//一个 goroutine无法直接终止另一个，因为这样会让所有的共享变量状态处于不确定状态

	//创建一个读取标准输入的 goroutine，它通常连接到终端
	//一旦开始读取任何输入，这个 goroutine 通过关闭 done 通道来取消事件
	go func() {
		os.Stdin.Read(make([]byte, 1)) //读一个字节
		close(done)
	}()
}

//定义一个工具函数，在被调用的时候检测或轮询取消状态
var done = make(chan struct{})

func cancelled() bool {
	select {
	case <-done:
		return true
	default:
		return false

	}
}

//性能剖析揭示了它的瓶颈在于 dirents 中获取信号量令牌的操作，下面的 select 让取消操作的延迟从白毫秒降到几十毫秒
func dirents(dir string) []os.FileInfo {
	select {
	case sema <- struct{}{}: //获取令牌
	case <-done:
		return nil //取消

	}
	defer func() { <-sema }() //释放令牌
}
