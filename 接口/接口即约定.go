package main

import "fmt"

//接口类型是对其他类型行为的概括和抽象
func main() {
	fmt.Println("接口即约定")

	//具体类型指定了它所含数据的精准布局，还暴露了基于这个精准布局的内部操作
	// GO 语言还有一种类型称作接口类型：接口是一种抽象类型，它并没有暴露所含数据的布局或者内部结构
	//当然也没有那些数据的基本操作，提供的仅仅是一些方法而已

	var c ByteCounter
	c.Write([]byte("hello"))
	fmt.Println(c)
	c = 0 //重置计数器
	var name = "Dolly"
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c)
}

//统计传入数据的字节数
type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, err) {
	*c += ByteCounter(len(p)) //转换 int 为 ByteCounter 类型
	return len(p), nil
} // *ByteCounter 满足 io.Writer 接口的约定，可以在 Fprintf 中使用它， Fprintf 察觉不到这种类型差异， ByteCounter 也能正确的积累格式化后结果的长度

//接口类型

//一个接口类型定义一套方法，如果一个具体类型要实现该接口，那必须实现接口类型定义中的所有方法

type Reader interface {
	Read(p []byte) (n int, err error)
}
type Closer interface {
	Close() error
}

//通过已有接口组合得到的新接口
type ReadWriteClose interface {
	Reader
	Closer
} //语法称为嵌入式接口，与嵌入式结构类似，可以直接使用一个接口，而不是逐一写出这个接口所包含的方法
