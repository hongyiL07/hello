package main

import (
	"time"
)

func main() {
	//text/template 包和 html/template 包提供一种机制，可以将程序变量的值带入到文本或者 HTML 模板中

	//模板是一个字符串或者文件，它包含一个或者多个两边用大括号包围的单元————{{...}},这称为 操作
	const temp1 = `{{.TotalCount}} issues:
{{range .Items}}----------------------------
Number: {{.Number}}
User:   {{.User.Login}}
Title:  {{.Title | printf "%.64s"}}
Age:    {{.CreatedAt | daysAgo}} days
{{end}}`
	//模板首先输出符合条件的 issue 数量，然后分别输出每个 issue 的序号，用户，标题和距离创建时间已过的天数
	//用点号（ . ）表示当前值的标记，最开始点号表示模板里的参数
	// {{.TotalCount}} 表示 TotalCount 成员的值，直接输出
	// {{range .Items}} 和 {{end}} 创建一个循环，所以内部的值会展开很多次，这时候点号（.）表示 Items 里面连续的元素
	//符号 |  会将前一个操作的结果当作下一个操作的输出
	//通过模板输出结果需要两个步骤，1. 需要解析模板（只需要一次）并转换为内部的方法。 2. 在指定的输入上执行

	//方法的链式调用：template.New 创建并返回一个新模板，funcs 添加 daysAgo 到模板内部可以访问的函数列表中，然后返回这个模板对象；最后调用 Parse 方法
	//report, err := template.New("report").Funcs(template.FuncMap{"daysAgo": daysAgo}).Parse(temp1)
	//if err != nil{
	//	log.Fatal(err)
	//}

	//创建了模板，添加了内部可调用函数 daysAgo ，然后解析，再检查，就可以使用 github.IssuesSearchResult 作为源数据，使用 os.Stdout 作为输出目标执行这个模板
	//var report1 = template.Must(template.New("issuelist").Funcs(template.FuncMap{"daysAgo": daysAgo}).Parse(temp1))
	//func main(){
	//	result ,err := github.IssuesSearch(os.Args[1:])
	//	if err != nil{
	//		log.Fatal(err)
	//	}
	//	if err := report.Execute(os.Stdout,result); err != nil{
	//		log.Fatal(err)
	//	}
	//}

}

//已过去的天数
func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

// time.Time 的 JSON 序列化值就是该类型标准的字符串表示方法
