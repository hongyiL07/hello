package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

//类型断言是一个作用在接口值上的操作，写出来类似于 x.(T),其中 x 是一个接口类型的表达式，而 T 是一个类型（称为断言类型）
//类型断言会检查作为操作数的动态类型是否满足指定的断言类型

func main() {
	fmt.Println("类型断言")

	//首先，如果断言类型 T 是一个具体类型，那么类型断言会检查 x 的动态类型是否就是 T 。
	//如果检查成功，类型断言的结果就是 x 的动态值，类型当然就是 T
	var w io.Writer
	w = os.Stdout
	f := w.(*os.File)

	//其次，如果断言类型 T 是一个接口类型，那么类型断言检查 x 的动态类型是否满足 T 。
	//如果成功，动态值并没有提取出来，结果仍然是一个接口值，接口值的类型和值部分也没有变更，只是结果的类型为接口类型 T
	rw := w.(io.ReadWriter)

	//经常无法确定一个接口值的动态类型，需要检测它是否为一个特定类型。
	//如果类型断言现在需要两个结果的赋值表达式，那么断言不会在失败时崩溃，而是多返回一个布尔型的返回值来指示断言是否成功
	var q io.Writer = os.Stdout
	f, ok := q.(*os.File)      //成功
	e, ok := q.(*bytes.Buffer) //失败

}

//  I/O 会因为很多原因失败，但有三类原因通常必须单独处理：文件已存储（创建操作），文件找不到（读取操作）以及  权限不足
//但由于处理 I/O 错误的逻辑会随着平台的变化而变化，因此这种方法很不健壮，同样的可能用完全不同的错误消息来报告
//可靠的方法是用专门的类型来表示结构化的错误值

type PathError struct {
	Op   string
	Path string
	Err  error
}

func (e *PathError) Error() string {
	return e.op + "" + e.Path + ":" + e.Err.Error()
}

//很多客户端忽略了 PathError ，改用一种统一的方法来处理所有的错误，即调用 Error 方法。
// PathError 的 Error 方法只是拼接了所有字段，而 PathError 的结构则保留了错误所有的底层信息

// Write 方法需要一个字节 slice ，而我们想写入的是一个字符串，所以 []byte(...) 转换就是必需的
