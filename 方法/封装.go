package main

import (
	"fmt"
	"time"
)
//如果变量或者方法是不能通过对象访问到的，这称作 封装 的变量或者方法。
//封装（有时候称作数据隐藏）是面向对象编程中重要的一方面
//Go语言只有一种方式控制命名的可见性：定义的时候，首字母大写的标识符是可以从包中导出的，反之则不能导出
//要封装一个对象，必须使用结构体
//Go语言中封装的单元是包而不是类型
//无论是在函数内的代码还是方法内的代码，结构体类型内的字段对于同一个包中的所有代码都是可见的

//封装提供的三个优点
//第一  ，因为使用方不能直接修改对象的变量，所以不需要更多的语句用来检查变量的值

//第二  ，隐藏现实细节可以防止使用方依赖的属性发生改变，使得设计者可以更加灵活的改变 API 的实现而不破坏兼容性
type Buffer struct {
	buf []byte
	initial [64]byte
}
// Grow 方法按需扩展缓冲区的大小
//保证 n 个字节的空间
func (b *Buffer) Grow(n int) {
	if b.buf == nil{
		b.buf = b.initial[0:]  //最初使用预分配的空间
	}
	if len(b.buf) + n > cap(b.buf){
		buf := make([]byte,b.Len(),2*cap(b.buf) + n)
		copy(buf,b,buf)
		b.buf = buf
	}
}

//第三  ，就是防止使用者肆意的改变对象内的变量
//因为对象的变量只能被同一个包内的函数修改，包的作者能够保证所有的函数都可以维护对象内部的资源
type Counter struct {
	n int
}

func (c *Counter) N ()int {
	return c.n
}
func (c *Counter) Increment() {
	c.n++
}
func (c *Counter) Reset() {
	c.n = 0
}

//仅仅用来获得或者修改内部变量的函数称为 getter 和 setter ，就像 log 包里的 Logger 类型
//然而命名 getter 方法的时候，通常将 Get 前缀省略

// Go 语言允许导出字段，一旦导出就必须面对 API 的兼容性问题，要考虑之后维护的复杂程度，将来发生的可能性，变化对原本代码质量的影响

//封装并不总是必须的。



func main()  {
	fmt.Println("封装")

	//time.Duration 对外暴露 int64 的整形用于获得微秒，这使我们能够对其进行通常的数学运算和比较操作，甚至定义常数
	const day  = 24 * time.Hour
	fmt.Println(day.Seconds())
}
