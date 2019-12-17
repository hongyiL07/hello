package main

import (
	"fmt"
	"os"
)

//函数可以递归调用，意味函数可以直接或者间接的调用自己（实用，可以处理许多带有递归特性的数据结构）

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Println(os.Stderr, "findlinks1 : %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}
