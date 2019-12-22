package main

import "fmt"

// go test 子命令是 Go 语言包的测试驱动程序，这些包根据某些约定组织在一起

//在 *—test.go 文件中，三种函数需要特殊对待，即 功能测试函数 ，基准测试函数 ，示例函数
//功能测试函数：
//以 Test 前缀命名的函数，用来检测一些程序逻辑的正确性， go test 运行测试函数，并报告结果是 PASS 还是 FAIL
//基准测试函数：
//名称以 Benchmark 开头，用来测试某些操作的性能， go test 汇报操作的平均执行时间
//示例函数：
//以 Example 开头，用来提供机器检查过的文档

// go test 工具扫描 *—test.go 文件来寻找特殊函数，并生成一个临时的 main 包来调用它们，然后编译和运行，并汇报结果，最后清空临时文件

func main()  {

	fmt.Println("go test 工具")
}
