package main

import (
	"fmt"
	"log"
	"os"
)

var g = "g"

func main() {
	f := "f"
	fmt.Println(f) //局部变量 f 覆盖了 包级函数 f
	//  fmt.Println(g)   包级变量
	//  fmt.Println(h)   编译错误：未定义 h

	x := "hello!"
	for i := 0; i < len(x); i++ {
		x := x[i]
		if x != '!' {
			x := x + 'A' - 'a'
			fmt.Printf("%c", x)
		}
	}
	fmt.Println("\n")
	for _, x := range x {
		x := x + 'A' - 'a'
		fmt.Printf("%c", x)
	}
	fmt.Println("\n")

}

//var cwd string

func init() {
	cwd, err := os.Getwd() //包级变量有 cwd 不能重复声明
	if err != nil {
		log.Fatal("os.Geted failed: %v", err)
	}
	log.Printf("Working Directory= %s", cwd)
}
