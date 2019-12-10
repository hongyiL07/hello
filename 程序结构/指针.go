package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separtor")

func main() {
	x := 1
	p := &x // 整型指针 p 指向 x 的地址
	fmt.Println(*p)
	*p = 2 //通过 p 指向的变量写出 *p
	fmt.Println(x)

	//两个指针当且仅当指向同一个变量或者两者都是nil的情况下才相等
	var i, j int
	fmt.Println(&i == &j, &i == &j, &i == nil)

	//间接传递变量值
	v := 1
	incr(&v)       //     v=2
	fmt.Println(v) //    v=3
	fmt.Println(incr(&v))

	// sep 使用 sep 替换默认参数输出时使用空格分隔符
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}

}
func incr(p *int) int {
	*p++ //递增 怕所指向的值；p 本身保持不变
	return *p
}
