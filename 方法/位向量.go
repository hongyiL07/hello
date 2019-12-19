package main

import (
	"bytes"
	"fmt"
)
// Go 语言的集合通常使用 map[T]bool 来实现，其中 T 是元素类型，使用 map 的集合扩展性良好
//在数据流分析领域，集合元素都是小的非负整数，集合拥有许多元素，而且集合的操作多数是求并集和交集，位向量是个理想的数据结构
//位向量使用一个无符号整形的 slice ，每一位代表集合中的一个元素。
//如果设置第 i 位的元素，则认为集合包含 i 。
type IntSet struct {
	// IntSet 是一个包含非负整数的集合
	//零值代表空的集合
	words []uint64
}
// Has 方法的返回值是否存在非负数 x
func (s *IntSet) Has(x int) {
	word, bit := x/64,uint(x%64)
	return word <len(s.words) && s.words[word]&(1<<bit) != 0
}
// add 添加非负数 x 到集合中
func (s *IntSet) add(x int) {
	word, bit := x/64,uint(x%64)
	for word >= len(s.words){
		s.words = append(s.word,0)
	}
	s.words[word] |= 1 << bit
}
// UnionWith 将会对 s 和 t 做并集且将结果存在 s 中
func (s *IntSet) UnionWith(t *IntSet) {     //使用按位或者操作符 | 来计算一次 64 个元素求并集的结果
	for i, tword := range t.words{
		if i<len(s.words){
			s.wordsp[i] |= tword
		}else {
			s.words = append(s.words,tword)
		}
	}
}
//由于每一个字拥有64位，因此为了定位 x 的位置，使用商数 x/64 作为字的索引，而 x%64 记作该字内部的索引


//以字符串 {1，2，3} 的形式返回集中
func (s *IntSet) String() string  {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i,word :=range s.words{
		if word == 0{
			continue
		}
		for j:=0;j<64;j++{
			if word &(1<<uint(j)) != 0{
				if buf.Len()>len("{"){
					buf.WriteByte(' ')
				}
				fmt.Println(&buf,64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}


func main()  {
	fmt.Println("位向量")
}

