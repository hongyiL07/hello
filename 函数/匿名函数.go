package main

import (
	"fmt"
)

//命名函数只能在包级别的作用域进行声明，但我们能够使用 函数字面量 在任何表达式内指定函数变量

//函数字面量就像函数声明，但在 func 关键字后面没有函数的名称。它是一个表达式，它的值称作 匿名函数
//函数字面量在我们需要的时候才定义

//函数变量类似于使用 闭包 方法实现的变量，Go 程序员通常把函数变量称为 闭包

func main() {
	//里层的函数可以使用外层函数中的变量
	f := sq()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())

	//为了爬虫开始工作，使用命令行参数指定开始的 URL
	//breadthFirst(crawl,os.Args[1:])

	//捕获迭代遍历
	//假设一个程序必须创建一系列的目录之后又要删除它们，可以使用一个包含函数变量的 slice 进行清理操作
	//var rmdirs []func()
	//for _, d:= range tempDirs(){
	//	dir := d                             //注意，这一行是必须的
	//	os.MkdirAll(dir,0755)         //创建父目录
	//	rmdirs = append(rmdirs, func() {
	//		os.RemoveAll(dir)
	//	})
	//}
	//for _, rmdir := range rmdirs{
	//	rmdir()  //清理
	//}   dir 变量的实际取值是最后一次迭代时的值并且所有的 os.RemoveAll 调用最终都试图删除同一个目录

}
func sq() func() int {
	var x int
	return func() int {
		x++
		return x * x // sq 函数返回另一个函数，类型是 func（）int，调用 sq 创建一个局部变量 x 而且返回一个匿名函数
		//每次调用 sq 都会递增 x 的值然后返回 x 的平方，第二次调用 sq 函数将创建第二个变量 x 然后返回一个递增 x 值的新匿名函数
	}
}

//网页爬虫的核心是解决图的遍历（广度优先遍历）
func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...) // f(item)...  将会把 f 返回的列表中所有项添加到 worklist 中
			}
		}
	}
}

//在爬虫里，项节点都是 URL 。我们提供 crawl 函数给 breadthFirst 以输出 URL ，解析链接然后将它们返回，标记为已访问
//func crawl( url string) []string {
//	fmt.Println(url)
//	list,err := links.Extract(url)      links 包提供解析链接的函数
//	if err != nil{
//		log.Printf(err)
//	}
//	return list
//}
