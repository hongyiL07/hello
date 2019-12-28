package main

import (
	"fmt"
	"sort"
)

// sort 包中提供了针对任意序列根据任意排序函数原地排序的功能
// Go 语言的 sort.Sort 函数对序列和其中元素的布局无任何要求，使用 sort.Interface 接口来指定通用排序算法和每个具体的序列类型之间的协议（contract）
//这个接口的实现确定了序列的具体布局（经常是一个 slice ），以及元素期望的排序方式

//一个原地排序算法需要知道三个信息：序列长度  比较两个元素的含义  如何交换两个元素  sort.Interface就有三个方法
type Interface interface {
	Len() int
	Less(i, j int) bool // i,  j 是序列元素的下标
	Swap(i, j int)
}

//简单的例子：字符串 slice 定义的新类型 StringSlice 以及它的 Len、Less、Swap 三个方法如下
type StringSlice []string

func (p StringSlice) Len() int {
	return len(p)
}
func (p StringSlice) Less(i, j int) bool {
	return p[i] < p[j]
}
func (p StringSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

//一个 slice 转换为 StringSlice 类型，生成了一个新的 slice
//与原始的 names 有同样的长度，容量，底层数组，不同的是额外增加了三个用于排序的方法
func main() {
	fmt.Println("sort.Interface 排序")

	sort.Sort(StringSlice(names))
}
