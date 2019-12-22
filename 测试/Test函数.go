//每个测试文件必须导入 testing 包，这些函数的函数签名如下
//func TestName(t *testing.T){}
//功能测试函数必须以 Test 开头，可选的后缀名称必须以大写字母开头，
//函数 t 提供了汇报失败和日志记录的功能

package word
//包 word 提供文字游戏相关的工具函数

import (
	"testing"
	"unicode"
)

//IsPalindrome 判断一个字符串是否位回文字符串
func IsPalindrome(s string) bool {
	for i:= range s{
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return  true
}


// TestPalindrome 和 TestNonPalindrome 两个肝功能测试函数都检查 isPalindrome 是否单个输入参数给出了正确的结果，并用 t.Error 来报错
func TestPalindrome(t *testing.T)  {
	if  !IsPalindrome("detartrated"){
		t.Error(`IsPalindrome("kayak") = false`)
	}
	if  !IsPalindrome("kayak"){
		t.Error(`IsPalindrome("kayak") = false`)
	}
}
func TestNonPalindrome(t testing.T)  {
	if IsPalindrome("palindrome"){
		t.Error(`IsPalindrome("palindrome") = true`)
	}
}

//特定的小 bug 自然导致了新的参数用例的产生
func TestFrenchPalindrome(t testing.T)  {
	if IsPalindrome("ete"){
		t.Error(`IsPalindrome("ete") = true`)
	}
}

//比较好的实践是先写测试然后发现它触发的错误和用户报告里的一致，只有这个时候，我们才能确信我们修复的内容是针对这个出现的问题

//另外，运行 go test 比手动测试 bug 报告中的内容要快的多，测试可以让我们顺序的检查内容

//选项 -run 的参数是一个正则表达式，它可以使得 go test 只运行那些测试函数名称匹配给定模式的函数

//一旦我们使得选择的测试用例通过之后，在我们提交更改之前，我们必须重新使用不带开关的 go test 来运行一次整个测试套件

//修改先前的函数
func IsPalindrome1(s string) bool {
	var letters []rune
	for _, r:= range s{
		if unicode.IsLetter(r) {
			letters = append(letters,unicode.ToLower(r))
		}
	}
	for i:= range letters{
		if letters[i] != letters[len(letters)-1-i]{
			return false
		}
	}
	return  true
}
//测试用例
func TestPa(t *testing.T)  {
	var tests =[]struct{
		imput string
		want bool
	}{
		{"",true},
		{"a",true},
		{"aa",true},
		{"ab",true},
	}
	for  _, test := range tests{
		if got := IsPalindrome1(test.imput); got != test.want{
			t.Error("IsPalindrome1(%q) = %v",test.imput,got)
		}
	}
}
//基于表的测试方式在 Go 里面很常见