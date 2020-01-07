package main

import "fmt"

//可增长的栈：每个 OS 线程都有一个固定大小的栈内存（通常为2MB）
//栈内存区域用于保存在其他函数调用期间那些正在执行或临时暂停的函数中的局部变量

//goroutine 调度：OS 线程由 OS内核来调度每隔几毫秒，一个硬件时钟中断发到 CPU，CPU调用一个叫 调度器的内核函数

// GO 调度器使用 GOMAXPROCS 参数来确定需要多少个 OS 线程来同时执行 Go 代码，默认值是机器上的 CPU 数量
//阻塞在 I/O 和其他系统调用中或调用非 GO 语言写的函数 goroutine 需要一个独立的 OS 线程，但不计算在 GOMAXPROCS 中

//goroutine 没有标识
//当前线程都有一个独特的标识，它通常可以取一个整数或者指针
//这个特性可以轻松构建一个 线程的局部存储，本质上就是一个 map
//以线程的标识作为键，每个线程都可以独立用这个 map 存储和获取值，而不受其他线程干扰

func main() {
	fmt.Println("goroutine 和线程")

	for {
		go fmt.Println(0)
		fmt.Println(1)
	} //无止境的输出 0 和 1
}
