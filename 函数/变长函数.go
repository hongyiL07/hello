package main

import (
	"fmt"
	"os"
)

//变长函数被调用的时候可以有可变的参数个数

func main() {
	fmt.Println(sum())
	fmt.Println(sum(3))

	fmt.Println(sum(1, 2, 3, 4))
	//等价于以下
	values := []int{1, 2, 3, 4}
	fmt.Println(sum(values...))

	linenum, name := 12, "count"
	errorf(linenum, "undefined:  %s", name)

}
func sum(vals ...int) int { //在参数列表最后的类型名称之前使用省略号“...” 表示声明一个变长函数，调用这个函数时，可以传递该类型任意数目的参数
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
} // sum 函数返回零个或者多个 int 参数。在函数体内，vals 是一个 int 类型的 slice 。调用 sum 的时候任意数量的参数都将提供给 vals 参数
//尽管 ...int 参数就像函数体内部的 slice ，但变长函数的类型和一个带有普通 slice 参数的函数的类型不相同

//变长函数通常用于格式化字符串
// errorf 函数构建一条格式化的错误信息，在消息的开头带有行号。函数的后缀 f 是广泛使用的命名习惯，用于可变长 Printf 风格的字符串格式化输出函数
func errorf(linenum int, format string, args ...interface{}) { // interface{} 类型意味着这个函数的最后一个参数可以接受任何值
	fmt.Fprintf(os.Stderr, "Line %d:", linenum)
	fmt.Fprintf(os.Stderr, format, args...)
	fmt.Fprintln(os.Stderr)
}
