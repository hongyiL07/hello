package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		links, err := findLinks(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "findlijnks2: %v\n", err)
			continue
		}
		for _, link := range links {
			fmt.Println(link)
		}
	}
}

// findLinks 发起一个HTTP 的 get 请求，解析返回的 HTML 页面，并返回所有链接
func findLinks(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}
	return visit(nil, doc), nil
}

//返回一个多值结果可以是调用另一个多值返回的函数，多了一个记录参数的动作
func findLinksLog(url string) ([]string, error) {
	log.Printf("findLinks %s", url)
	return findLinks(url)
}

//以下两个输出语句效果一致
//log.Printf（findLinks(url)）

//links,err := findLinks(url)
//log.Println(links, err)

//一个函数如果有命名的返回值，可以省略 return 语句的操作数， 称为 裸返回
//裸返回 可以消除重复代码
//裸返回是将每个命名返回结果按照顺序返回的快捷方法
