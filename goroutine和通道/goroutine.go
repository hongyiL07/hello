package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

//在 Go 里面，每一个并发执行的活动称为 goroutine
//当一个程序启动时，只有一个 goroutine 来调用 main 函数，称它为 主goroutine
//语法上，一个 Go 语句是在普通的函数或者方法调用前加上 go 关键字前缀，go 语句使函数在一个新创建的 goroutine 中调用

func main() {
	fmt.Println("goroutine 和通道")

	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) //例如，连接终止
			continue
		}
		handleConn(conn) //一次处理一个连接
	}

	//只读 TCP 客户端
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	mustCopy(os.Stdout, conn)
}
func handleConn(c net.Conn) { //处理一个完整的客户连接
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:08")) // time.Now().Format 方法提供了格式化日期和时间信息的方式
		if err != nil {
			return //例如，连接断开
		}
		time.Sleep(1 * time.Second)
	}
} // Listen 函数创建一个 net.Listener 对象它在一个网络端口上监听进来的连接，这里是 TCP 端口 localhost：8000

//只读 TCP 客户端
func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}

// killall 命令是 UNIX 的一个实用程序，用来终止所有指定名字的进程

//第二个客户端必须等到第一个客户端结束才能正常工作，因为服务器是顺序的，一次只能处理一个客户请求
//让服务器支持并发只需要一个很小的改变：在调用 handleConn 的地方添加一个 go 关键字，使它在自己的 goroutine 内执行
//for {
//	conn,err := listener.Accept()
//	if err!=nil{
//		log.Print(err)   //例如，连接终止
//		continue
//	}
//	go handleConn(conn)   //并发处理连接
