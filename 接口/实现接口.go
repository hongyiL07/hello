package main

//一个接口实现一个接口要求的所有方法，那这个类型实现了这个接口
import (
	"bytes"
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	fmt.Println("实现接口")

	//接口的赋值规则很简单，仅当一个表达式实现了一个接口时，这个表达式才可以赋给该接口
	var w io.Writer
	w = os.Stdout
	w = new(bytes.Buffer)
	w = time.Second // time.Duration 缺少 Write 方法

	var rwc io.ReadWriteCloser
	rwc = os.Stdout
	rwc = new(bytes.Buffer) // *bytes.Buffer 缺少 Close 方法

	//当右侧表达式也是一个接口时，该规则也有效
	w = rwc
	rwc = w

	//接口类型 interface{} 完全不包含任何方法，称为 空接口类型，是不可缺的
	//空接口类型对其实现类型没有任何要求，可以把任何值赋给空接口类型
	var any interface{}
	any = true
	any = 12.34
	any = "heool"

	//判断是否实现接口只需要比较具体类型和接口类型的方法，没必要在具体类型的定义中声明这种关系
	// *bytes.Buffer 必须实现 io.Write
	var q io.Writer = new(bytes.Buffer)
	//不想引用 q ，可替换成 空白标识符
	var _ io.Writer = (*bytes.Buffer)(nil)
	//非空的接口类型通常由一个指针类型来实现，特别是当接口类型的一个或多个方法暗示会修改接收者的情形
	//一个指向结构的指针才是最常见的方法接收者

	//指针类型肯定不是实现接口的唯一类型，即使是那些包含了会改变接收者方法的接口类型，也可以由 Go 的其他引用类型来实现
}

//对每一个具体类型 T ，部分方法的接收者就是 T ，而其他方法的接收者则是 *T 指针
//同时对类型 T 的变量直接调用 *T 的方法也是合法的，只要改变量是可变的，编译器隐式的完成取地址操作

//每一种抽象都用一种接口类型来表示
type Text interface {
	Pages() int
	Words() int
	PageSize() int
} //这些接口只是一种把具体类型分组并暴露它们共性的方式还可以发现其他分组方式
