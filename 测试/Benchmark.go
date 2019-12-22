package main

import (
	"fmt"
	"testing"
	)

func main()  {
	fmt.println("Benchmark")
}



//基准测试是在一定的工作负载之下检测程序性能的一种方法
//Go语言中，基准测试函数看上去像一个测试函数，前缀是 Benchmark 并且拥有一个 *testing.B 参数用来提供大多数和 *testing.T 相同的方法
//额外增加了一些与性能检测相关的方法

//IsPalindrome 判断一个字符串是否位回文字符串
func IsPalindrome(s string) bool {
	for i:= range s{
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return  true
}

//IsPalindrome函数的基准测试，它在一个循环中调用了 IsPalindrome 共 N 次
func BebchmarkIsPalindrome(b *testing.T)  {
	for i:=0;i<b.N;i++{
		IsPalindrome("A man A plan,a canal: panama")
	}
}

//既然有了基准测试和功能测试，就想让程序更快一点，最明显的优化是使得 IsPalindrome 函数的第二次循环在中间停止检测以避免比较两次
func IsPalindrome1(s string) bool {
	letters := "safafafaf"
	n := len(letters)
	for i:= 0;i<n;i++ {
		if letters[i] != letters[len(letters)-1-i]{
			return false
		}
	}
	return  true
}

//性能比较函数只是普通的代码，它们的表现形式通常是带有一个参数的函数，被多个不同的 Benchmark 函数传入不同的值来调用
func benchmark(b *testing.B,size int)  {
	//参数 size 指定了输入的大小，每个 Benchmark 函数传入的值都不同但是在每个函数内部是一个常量
}




