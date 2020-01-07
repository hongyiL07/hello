package main

import (
	"database/sql"
	"expvar"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

// Go 语言运行时和工具链装备了一个精致并易于使用的动态分析工具 ： 竟态检测器
//简单的把 -race 命令行参数加到 go build · go run  · go  test 命令里面即可使用该功能

func main() {
	fmt.Println("并发非阻塞缓存")
	//并发非阻塞的缓存系统： 解决在并发实战很常见但已有的库
	//但是不能解决一个问题:函数记忆问题，即函数缓存的结果达到多次调用但只须计算一次的效果

	//对于一串请求 URL 中的每个元素，首先调用 Get ，记录延时和它返回的数据长度
	m := memo.New(httpGetBody)
	for url := range incomingURLs() {
		start := time.Now()
		value, err := m.Get(url)
		if err != nil {
			log.Print(err)
		}
		fmt.Println(url, time.Since(start), len(value.([]byte)))
	}

	////修改测试来让所有请求并发进行，使用 sync.WaitGroup 来做到等最后一个请求完成后再返回的效果
	m1 := memo.New(httpGetBody)
	var n sync.WaitGroup
	for url := range incomingURLs {
		n.Add(1)
		go func(url string) {
			start := time.Now()
			value, err := m.Get(url)
			if err != nil {
				log.Print(err)
			}
			fmt.Println(url, time.Since(start), len(value.([]byte)))
			n.Done()
		}(url)
	} //运行快 但是它并不是每次都能正常运行
}
func httpGetBody(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
	//返回两个结果  一个[ ]byte 和一个 error
	//因为它们可以直接赋给 httpGetBody 声明的结果类型 interface{}  和 一个 error，所以可以直接返回这个结果而不用额外的处理
}

/*让缓存并发安全的最简单的方法就是基于一个监控的同步机制， Memo 加一个互斥量，
并在 Get 函数开头获取互斥锁，在 返回前释放互斥锁，这样两个 cache 相关的操作就发生在临界区域
type Memo struct {
	f Func
	mu  sync.Mutex  //保护 cache
	cache map[string]result
}
 Get 是并发安全的
func (memo *Memo) Get(key string) (value interface{}, err error){
	memo.mu.Lock()
	res ,ok:= memo.cache[key]
	if ! ok{
		res.value,res.err= memo.f(key)
		memo.cache[key]=res
	}
	memo.mu.Unlock()
	return res.value,res.err
}*/

//我们需要一个 非阻塞的缓存，一个不会把它需要记忆的函数串行运行的缓存
//主调 goroutine 会分两次获取锁，第一次用于查询，第二次用于查询五返回结果时进行更新，两次之间，其他 goroutine 也可以使用缓存
