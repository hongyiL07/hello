//随机测试 ：通过构建随机输入来扩展测试的覆盖范围
package word

import (
	"math/rand"
	"testing"
	"time"
)
//输入是随机的，函数输出什么内容：
//第一种：额外写个函数，使用低效但是清晰的算法，然后检查这两种实现的输出是否一致
//第二章：构建符合某种格式的输入，这样可以知道它们对应的输出是什么
//randomPalindrome 函数产生一系列回文字符串，这些输出在构建的时候就确定是回文字符串了
func randomPalindrome(rng *rand.Rand) string {  //randomPalindrome返回一个回文字符串，它的长度和内容都是随机数生成器//rng 生成的
	n:=rng.Intn(25)   //随机字符串最大长度为 24
	runes := make([]rune,n)
	for i := 0;i<(n+1)/2;i++{
		r:=rune(rng.Intn(0x1000))
		runes[i] = r
		runes[n-i-1] = r
	}
	return string(runes)
}
func TestRandompalindromes(t *testing.T)  {
	//初始化一个伪随机数生成器
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d",seed)
	rng := rand.New(rand.NewSource(seed))
	for i:=0;i<1000;i++{
		p := randomPalindrome(rng)
		if !IsPalindrome1(p){
			t.Errorf("IsPalindrome1(%q) = false",p)
		}
	}
}
//由于随机测试的不确定性，在遇到测试用例失败的情况下，一定要记录足够的信息以便于重现这个问题