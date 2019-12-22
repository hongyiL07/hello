package main

import "fmt"

func main()  {
	fmt.Println("Example 函数")
}

//IsPalindrome 判断一个字符串是否位回文字符串
func IsPalindrome(s string) bool {
	for i:= range s{
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return  true
}

//被 Go 语言特殊对待的第三种函数就是示例函数，它们的名字以 Example 开头。既没有参数也没有结果

func Example()  {
	fmt.Println(IsPalindrome("a man, a plan,a canal: Panama"))
	fmt.Println(IsPalindrome("palindrome"))
	//输出
	//true
	//false
}
//如果有一个示例函数就叫 Example ，那么它就和包 word 关联在一起