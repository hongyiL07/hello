package main

import "fmt"

//需要互斥锁的两个原因
//第一：防止 Balance 插到其他操作中间
//第二：因为同步不仅涉及多个 goroutine 的执行顺序问题，同步还会影响到内存


func main()  {
	fmt.Println("内存同步")


	//由于两个 goroutine 并发运行且没有使用互斥锁的情况下访问共享变量，会有数据竟态
	var x,y int
	go func() {
		x=1
		fmt.Println("y:",y,"  ")
	}()
	go func() {
		y=1
		fmt.Println("x:",x,"  ")
	}()

	//在单个 goroutine 内，每个语句的效果保证按照执行的顺序发生，也就是说 goroutine 是串行一致的

	// y 输出的可能是过期的值
}
