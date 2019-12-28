package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"time"
)

//实现睡眠指定时间的功能
var period = flag.Duration("period", 1*time.Second, "sleep period") //默认时间为 1 秒
func main() {
	flag.Parse()
	fmt.Println("Sleeping for %v...", *period)
	time.Sleep(*period)
	fmt.Println()
}

//定义一个满足 flag.Value 接口的类型
//package flag
// Value 接口代表了存储在标志内的值
type Value interface {
	String() string
	Set(string) error
}

//String 方法用于格式化标准对应的值

//一个类型的接口值（简称接口值）其实有两部分：
//一个具体类型和该类型的值/二者称为接口的动态类型和动态值
//用类型描述符来提供每个类型的具体信息
//对于一个接口值，类型部分就用对应的类型描述符来表达

//变量 w 有个不同的值（最初和最后是同一个值）
func main1() {
	var w io.Writer //声明了 w
	//一个接口值是否是 nil 取决于它的动态类型
	//接口的零值就是把它的动态类型和值都设置为 nil
	//可以用 w == nil 或者 w ！= nil 来检测一个接口值是否是 nil
	//调用一个 nil 接口的任何方法都会导致崩溃
	w.Write([]byte("hello")) //崩溃

	//把一个 *os.File 类型的值赋给了 w
	//具体类型隐式转换为一个接口类型，与对应的显示转换 io.Writer(os.Stdout) 等价
	w.Write([]byte("hello")) //hello
	//等价以下语句
	os.Stdout.Write([]byte("hello")) //hello
	//调用该接口的 Write 方法，实际调用（*os.File）.Write 方法，输出  hello
	w = os.Stdout

	w = new(bytes.Buffer)
	//把一个 *bytes.Buffer 类型的值赋给了接口值
	//动态类型现在是 *bytes.Buffer ，动态值现在则是一个指向新分配缓冲区的指针
	//调用 Write 方法的机制和第二句一致
	w.Write([]byte("hello")) //把 hello 写入 bytes.Buffer
	//类型描述符是 *bytes.Buffer ，所以调用的是（*bytes.Buffer）.Write 方法，接收者是缓冲区的地址。调用该方法会追加 “hello” 到缓冲区

	w = nil
	//把 nil 赋给了接口值，把动态类型和动态值都设置为 nil， 把 w 恢复到了刚声明时的状态

	//var buf *bytes.Buffer
	//if debug{
	//	buf = new(bytes.Buffer)  //启用输出收集
	//}
	//f(buf)
	//if debug {
	//...使用 buf ...
	//}
	//修改方案
	var buf io.Writer
	if debug {
		buf = new(bytes.Buffer) //启用输出收集
	}
	f(buf)
}

//一个接口值可以指向多个任意大的动态值
var x interface{} = time.Now()

//含有空指针的非空接口
//空的接口值（其中不包含任何信息）与仅仅动态值为 nil 的接口值是不一样的

//当 debug 设置为 true 时，主函数收集函数 f 的输出到一个缓冲区中
const debug = true

//如果 out 不是 nil ，那么会向其写入输出的数据
func f() {
	if out != nil {
		out.Write([]byte("done!\n"))
	}
}
