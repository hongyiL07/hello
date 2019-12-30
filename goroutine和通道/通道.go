package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

//如果说 goroutine 是 Go 程序并发的执行体，通道就是它们之间的连接。通道是可以让一个 goroutine 发送特定值到另一个 goroutine 的通道机制
//每个通道是一个具体类型的导管，叫做通道的元素类型。   一个有 int 类型的通道写为 chan int

func main() {
	fmt.Println("通道")

	//使用内置的 make 函数来创建一个通道
	ch := make(chan int) // ch 的类型是  chan int
	//使用简单的 make 调用创建的通道叫无缓冲通道，但 make 还可以接受第二个可选参数，一个表示通道容量的整数，如果是 0 ，make 创建一个无缓冲通道
	ch = make(chan int, 9) //容量为 9 的缓冲通道

	//使用无缓冲进行的通信导致发送和接收 goroutine 同步化。因此，无缓冲通道也称为 同步通道。
	//当一个值在无缓冲通道上传递时，接收值后发送方 goroutine 才被再次唤醒
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn)
		log.Println("done")
		done <- struct{}{} //指示主 goroutine
	}()
	mustCopy(conn, os.Stdin)
	conn.Close()
	<-done //等待后台 goroutine 完成
	//当用户关闭标准输入流时，mustCopy 返回，主 goroutine 调用 conn.Close()来关闭两端网络连接

	//通道有两个主要操作：发送（send）和 接收（receive），两者统称通信。
	//send 语句从一个 goroutine 传输一个到另一个在执行接收表达式的 goroutine
	ch <- x  //发送语句
	x = <-ch //赋值语句中的接收表达式
	<-ch     //接收语句，丢弃结果

	//通道的第三个操作就是：关闭（close），它设置一个标志位来指示值当前已经发送完毕，这个通道后面没有值了，关闭后的发送操作将导致宕机
	//调用内置的 close 函数来关闭通道：
	close(ch)
}
