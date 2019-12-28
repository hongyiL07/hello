package main

import (
	"database/sql"
	"fmt"
	"go/types"
)

//接口有两种不同的风格
//第一种，典型的比如 io.Reader  io.Write 等等，接口上的各种方法突出了满足这个接口的具体类型之间的相似性，
//但隐藏了各个具体类型的布局和各自特有的功能，这种风格强调了方法，而不是具体类型
//第二种，充分利用了接口值能够容纳各种具体类型的能力，它把作为这些类型的联合来使用，
//强调的是满足这个接口的具体类型，而不是这个接口的方法（何况经常没有），也不注重信息隐藏，使用方式称为 可识别联合

//以上两种分别对应 子类型多态 和 特设多态

func main() {
	fmt.Println("类型分支")
}

//与其它语言一样， Go 语言的数据库 SQL 查询 API 也允许干净的分离查询中的不变部分和可变部分
func listTracks(db sql.DB, artist string, minYear, maxYear int) {
	result, err := db.Exec(
		"select * from tracks where artist = ? and ?<=year and year >=?,artist,minYear,maxYear")
	//...
} // Exec 方法把查询字符串中的每一个 “ ？ ” 都替换为与相应参数值对应的 SQL 字面量，这些参数可能是布尔型 数字 字符串或者 nil

//将每个参数值转为对应的 SQL 字面量
func sqlQuote(x interface{}) string {
	if x == nil {
		return "NULL"
	} else if _, ok := x.(int); ok {
		return fmt.Sprintf(x)
	} else if _, ok := x.(uint); ok {
		return fmt.Sprintf(x)
	} else if _, ok := x.(bool); ok {
		if b {
			return "TRUE"
		}
		return "FALSE"
	} else if _, ok := x.(string); ok {
		return sqlQuoteString(s)
	} else {
		panic(fmt.Sprintf("unexpected type %T:%v", x, x))
	}
}

// sqlQuote 的分支语句也可以如下
//switch x.(type){
//case nil:
//	case int,unit:
//		case bool:
//}

//类型分支不允许使用 fallthrough 。
//在代码中， bool 和 string 分支的逻辑需要访问由类型断言提取出来的原始值，这个需求比较典型
//所以类型分支语句也有一种扩展形式，它用来把每个分支提取出来的原始值绑定到一个新的变量

//用类型分支的扩展形式重写后的 sqlQuote 就更加清晰可读了
func sqlQuote1(x interface{}) string {
	switch x := x.(type) {
	case nil:
		return "NULL"
	case int, uint:
		return fmt.Sprintf(x) //这里 x 的类型为 interface{}
	case bool:
		if x {
			return "TRUE"
		}
		return "FALSE"
	case string:
		return sqlQuotestring(x)
	default:
		panic(fmt.Sprintf("unexpected type %T : %v", x, x))
	}
}
