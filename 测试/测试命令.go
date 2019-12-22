package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
)

// go test 工具测试库代码包很有用，但是也可以将它用于测试命令。包名 main 一般会产生可执行文件，但是也可以当作库来导入

// echo 执行逻辑，而 main 用来读取和解析命令行参数以及报告 echo 函数可能返回的错误
var (
	n = flag.Bool("n",false,"omit trailing newline")
	s = flag.String("s"," ","separator")
)
var out io.Writer = os.Stdout   //测试过程中被更改
func main()  {
	flag.Parse()
	if err := echo(!*n,*s,flag.Args());err!=nil{
		fmt.Fprintf(os.Stderr,"echo: %v\n", err)
		os.Exit(1)
	}
}
func echo(newline bool,sep string,args []string) error {
	fmt.Fprint(out,strings.Join(args,sep))
	if newline{
		fmt.Fprintln(out)
	}
	return nil
}

//测试用例还可以用其他的代替 Writer 实现来记录写入的内容以便于后面检查
func TestEcho(t *testing.T)  {
	var tests = []struct{
		newline bool
		sep string
		args []string
		want string
	}{
		{true,"",[]string{},"\n"},
		{false,"",[]string{},""},
		{true,"\t",[]string{"one","two","three"},"one\ttwo\n"},
	}
	for _, test := range tests{
		descr := fmt.Sprintf("echo(%v,%q,%q)",test.newline,test.sep,test.args)

		out = new(bytes.Buffer)     //捕获输出
		if err := echo(test.newline,test.sep,test.args);err!=nil{
			t.Errorf("%s failed: %v",descr,err)
			continue
		}
		got := out.(*bytes.Buffer).String()
		if got != test.want{
			t.Errorf("%S = %q,want %q",descr,got,test.want)
		}
	}
}
//测试代码和产品代码在一个包里，尽管包的名称叫做 main ，并且里面定义了一个 main 函数，
//但是在测试过程中，这个包当作库来测试， 并且将函数 TestEcho 递送到测试驱动程序，而 main 函数则被忽略了






