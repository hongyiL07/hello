package main

import "fmt"

// encoding/json 包的 Marshal 和 Unmarshal 函数来把 JSON 文档解析为 Go 语言的数据结构
// encoding/xml 还为解析 API 提供了一个基于标记的底层 XML，在这些 API 中解析器读入输入文本，然后输出一个标记流
//标记流中主要包含四种类型：StartElement  EndElement  CharData  Comment

func main() {
	fmt.Println("XML解析")
}

// containsAll 判断 x 是否包含 y 中的所有元素，且顺序一致
func containsAll(x, y []string) bool {
	for len(y) <= len(x) {
		if len(y) == 0 {
			return true
		}
		if x[0] == y[0] {
			y = y[1:]
		}
		x = x[1:]
	}
	return false
}

//在 main 函数的每次循环中，如果遇到 StarElement，就把元素的名字入栈，遇到 EndElement 则把元素出栈
// API 保证了 StarElement 和 EndElement 标记是正确匹配的，对于不规范的文档也是如此
//当 消灭了select 遇到 CharData 时，如果栈中的元素名称顺序包含命令行参数给定的名称，就输出对应的文本
