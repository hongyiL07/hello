package main

import (
	"fmt"
	"strconv"
)

func main() {
	//整数转换成字符串 fmt.Sprintf
	//另一种使用函数 strconv.Itoa("integer to ASCII")
	x := 123
	y := fmt.Sprintf("%d", x)
	fmt.Println(y, strconv.Itoa(x)) //123   123

	// Formatter 和 formatUint 可以按不同的进位制格式化数字
	fmt.Println(strconv.FormatInt(int64(x), 2))

	//fmt.Printf 里的谓词 %b %d %o %x 比 Format函数方便 ，若要包含数字以外的附加信息很有用
	fmt.Sprintf("x=%b", x)

	// strconv 包内的 Atoi 函数或者 ParseInt 函数用于解释表示整数的字符串   而 ParseUint 用于无符号整数
	a, err := strconv.Atoi("123")
	b, err := strconv.ParseInt("123", 10, 64) // 10：十进制   64：最长为 64 位
	if err != nil {
		fmt.Println(a, b)
	}

}
