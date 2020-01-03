package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	worklist := make(chan []string)

	//从命令行参数开始
	go func() { worklist <- os.Args[1:] }()

	//并发爬取web
	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				go func(link string) {
					worklist <- crawl(link)
				}(link)
			}
		}
	}
}

//发送给任务列表的命令行参数必须在它自己的 goroutine 中运行来避免死锁，死锁是种卡住的情况
//其中主 goroutine 和一个爬取 goroutine 同时发生给对方但是双方都没有接收
func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

//可以使用容量为 n 的缓冲通道来建立一个并发原语，称为 计数信号量
