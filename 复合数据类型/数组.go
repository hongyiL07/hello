package main

import "fmt"

func main() {
	var p [3]int             //定义 3 个整数的数组
	fmt.Println(p[0])        //输出第一个元素
	fmt.Println(p[len(p)-1]) //输出最后一个元素 即  a[2]

	//仅输出元素
	for _, v := range p {
		fmt.Printf("%d\n", v)
	}

	//新数组中的元素初始值为元素类型的 0 值
	var w [3]int = [3]int{1, 2}
	fmt.Println(w[2])

	//若省略号“..."出现在数组长度的位置，数组的长度由初始化数组的元素个数决定
	q := [...]int{1, 2, 3, 4}
	fmt.Println(len(q))
	fmt.Println(q)

	//数组  slice  map 和结构体的字面语法都是相似的
	type currency int
	const (
		USD currency = iota //0
		EUR                 //1
		GBP                 //2
		RMB                 //3
	)
	//初始任意元素，其他默认为 0 值
	symbol := [...]string{USD: "美元", RMB: "人民币"}
	fmt.Println(RMB, symbol[RMB])
	fmt.Println(EUR, symbol[USD])

	//如果一个数组的元素类型是 可比较 的，那这个数组也是可比较的
	a := [2]int{1, 2}
	b := [...]int{1, 2}
	c := [2]int{1, 3}
	fmt.Println(a == b, a == c, b == c) // T F F
	//d:=[3]int{1,2}
	//fmt.Println(a==d)    编译错误：无法比较 [2]int == [3]int

}

//将一个数组的元素清零  (两种方法)
func zero1(ptr *[32]byte) {
	for i := range ptr {
		ptr[i] = 0
	}
}
func zero2(prt *[32]byte) {
	*prt = [32]byte{}
}
