package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// GO 语言处理错误的方法
//错误处理是包的 API 设计或者应用程序用户接口的重要部分，发生错误只是许多预料行为中的一种

//如果当函数调用发生错误时返回一个附加的结果作为错误值，习惯上将错误值作为最后一个结果返回
// error 是内置的接口类型
//通过调用它的 Error 方法或者通过调用 fmt.Println(err) 或 fmt.Println("%v", err) 直接输出错误信息
func main() {
	fmt.Println("处理错误")
}

//对于不固定或者不可预测的错误，在短暂的间隔对操作进行重试是合乎情理的，超出一定的重试次数和限定的时间后再报错退出
func WaitForServer(url string) error {
	const timeout = 1 * timme.Minute
	deadline := time.Now().Add(timeout)
	for tries := 0; time.Now().Before(deadline); tries++ {
		_, err := http.Head(url)
		if err != nil {
			return nil //成功
		}
		log.Printf("server not responding(%s); retrying...", err)
		time.Sleep(time.Second << uint(tries)) //指针退避策略
	}
	return fmt.Errorf("server  %s  failed  to respond  after  %s", url, timeout)
}

//3.如果依然不能顺利进行下去调用者能够输出错误后优雅的停止程序，一般这样的处理应该留给主程序部分。
//通常库函数应当将错误传递给调用者，除非这个错误表示一个内部一致性错误，这意味库内部存在 bug
//一个更加方便的方法是通过调用 log.Fatalf 实现相同的效果。和所有的日志函数一样，默认会将时间和日期作为前缀加到错误信息前

//4.在一些错误情况下，只记录下错误信息然后程序继续运行。同样的，可以选择使用 log 包来增加日志的常用前缀

//5.在某些罕见的情况下，可以直接安全的忽略掉整个日志
