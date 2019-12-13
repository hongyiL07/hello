package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// map 是一个拥有键值对元素的无序集合，键的值是唯一的键对应的值可以通过键来获取 更新 或者 移除
	// map 是散列表的引用， map 的类型是 map[K]V ,其中 K 和 V 是字典的 键 和 值对应的数据类型，
	// map 中所有的键都拥有相同的数据类型，同时所有的值也拥有相同的数据类型，但是 键的类型 和 值的类型 不一定相同
	// map 可以检测一个键是否已经存在

	//内置函数 make 可以用来创建一个 map
	ages := make(map[string]int) //创建一个从 string 到 int 的 map
	//也可以使用 map 的 字面量 来新建带初始化键值对元素的字典
	//ages := map[string]int{
	//	"alice"    : 31,
	//	"charile"  : 34,
	//}
	//等价于以下：
	//ages := make(map[string]int)
	//ages["alice"] = 31
	//ages["24444"] = 34

	//新的空 map 的另一种表达式为  ：  map[string]int{}

	// map de 元素访问也是通过下标的方式
	ages["alice"] = 32
	ages["4646"] = 31
	fmt.Println(ages["alice"]) //32
	fmt.Println(ages)
	//内置函数 delete 来从字典中根据键移除一个元素
	delete(ages, "4646")
	fmt.Println(ages)

	// map 使用给定的键来查找元素，如果对应的元素不存在，就返回值类型的零值
	// ages["bob"] 的值为 0
	ages["bob"] = ages["bob"] + 1
	fmt.Println(ages) // map[ bob : 1]

	//快捷赋值方式（ x+=y 和 x++ ）对 map 中的元素同样适用
	//ages["bob"] += 1
	//ages["bob"] ++

	// map 元素不是一个变量，不可以获取它的地址
	//一个原因是 map 的增长可能会导致已有元素被重新散列到新的储存位置，这样就可能使获取的地址无效

	//可以使用 for 循环（结合 range 关键字） 来遍历 map 中所有的 键 和 对应的值
	for name, age := range ages {
		fmt.Println(name, age)
	} //循环语句的连续迭代将会使得变量 name 和 age 被赋予 map 中的下一对 键 和 值

	// map 中元素的迭代顺序是不固定的，不同的实现方法会使用不同的散列算法，得到不同的元素顺序
	//如果键是字符串类型，可以使用 sort 包中的 Strings 函数来进行键的排序
	//import "sort"
	//var names  []String
	//for name :=range ages{
	//	names = append(names,name)
	//}
	//sort.String(names)
	//for _, name := range names{
	//	fmt.Println(name,ages[name])
	//}

	//创建一个初始元素为空但是容量足够容纳 ages map 中所有键的 slice
	names := make([]string, 0, len(ages))
	fmt.Println(names)

	// map 类型的零值是 nil， 也就是说，没有引用任何散列表
	var qw map[string]int
	fmt.Println(qw == nil)    // true
	fmt.Println(len(qw) == 0) // true

	//设置元素前  必须初始化 map

	//如果元素类型是数值类型，你需要能够辨别一个不存在的元素或者恰好这个元素的值是 0
	//age , ok := ages["bob"]
	//if age , ok := ages["bob"]; !ok {
	//	/*  ...  */
	//}   输出两个值，第二个值是布尔值，用来报告该元素是否存在（ ok ）

}

// map 不可比较，唯一合法的是和 nil 作比较
//判断两个 map 是否拥有相同的键和值，必须写一个循环
func equal0(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}
	return true
}

//使用 map 的键来储存这些已出现的行，来确保接下来出现的相同行不会输出
func equal1() {
	seen := make(map[string]bool) //字符串集合
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		line := input.Text()
		if !seen[line] {
			seen[line] = true
			fmt.Println(line)
		}
	}
	if err := input.Err(); err != nil {
		fmt.Println(os.Stderr, err)
		os.Exit(1)
	}
}

//函数 ReadRune 解码 UFT-8 编码，并返回三个值：解码的字符  UFT-8 编码中字节的长度 和 错误值
//如果输入的不是合法的 UFT-8 的字符，返回的字符 code.ReplacementChar 并且长度为 1

// map 的值类型本身可以是复合数据类型，例如 map 或 slice 。
//变量 graph 的键类型是 string 类型；值类型是 map 类型 map[string]bool ,表示一个字符串集合
var graph = make(map[string]map[string]bool)

//建立了一个从字符串的字符串集合的映射
func addEdge(form, to string) {
	edges := graph[form]
	if edges == nil {
		edges = make(map[string]bool)
		graph[form] = edges
	}
	edges[to] = true
}
func hasEdge(form, to string) bool {
	return graph[form][to]
}
