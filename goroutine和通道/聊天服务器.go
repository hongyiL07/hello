package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

//主 goroutine 的工作是监听端口，接受连接客服端的网络连接
func main() {
	fmt.Println("聊天服务器")

	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

//下一个是广播器，使用局部变量 clients 来记录当前连接的客户集合
//每个客户唯一被记录的信息是其对外发送消息通道的 ID
type client chan<- string

var (
	entering = make(map[client]bool)
	leaving  = make(chan client)
	messages = make(chan string) //所有接收的客户消息
)

func broadcaster() {
	clients := make(map[client]bool)
	for {
		select {
		case msg := <-messages:
			//把所有接收的消息广播给所有客户
			//发送消息通道
			for cli := range clients {
				cli <- msg
			}
		case cli := <-entering:
			clients[cli] = true
		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		}
	} //广播者监听两个全局的通道 entering 和 leaving ，通过它们通知客户的到来和离开
}

//下面是每个客户自己的 goroutine 。handleConn函数创建一个对外发送消息的新通道，然后通过 entering 通道通知广播者新客户的到来
//接着读取客户发来的每一行文本，通过全局接收消息通道将每一行发送给广播者，发送时在每条信息前面加上发送者 ID 作为前缀
//一旦从客户端读取完毕消息， handleConn 通过 leaving 通道通知客户离开，然后关闭连接
func handleConn(conn net.Conn) {
	ch := make(chan string) //对外发送客户信息的通道
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	ch <- "You are" + who
	messages <- who + "has arrived"
	entering <- ch

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + ":" + input.Text()
	} // 注意，忽略input.Err（） 中可能的错误

	leaving <- ch
	messages <- who + "has left"
	conn.Close()
}

//另外，handleConn 函数还为每个客户创建了写入（clientWriter）goroutine，它接收消息，广播到客户的发送消息通道中，然后将它们写到客户的网络连接中
//客户写入者的循环在广播者收到 leaving 通知并且关闭客户的发送消息通道后终止
func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) //注意网络层面的错误
	}
}
