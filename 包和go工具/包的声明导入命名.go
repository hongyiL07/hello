package main

import _ "image/png" //对包级变量执行初始化表达式求值（空导入的作用）
import (             //导入声明
	"fmt"
	"math/rand"
)

func main() {
	fmt.Printf("导入")
}
