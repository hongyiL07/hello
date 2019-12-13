package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func main() {
	// javascript 对象表示法（JSON）是一种发送和接收格式化信息的标准

	// JSON 的对象用来编码 GO 里面的 map （键为字符串类型）和结构体
	//boolean   true
	//number    -273.12
	//string    "She said \" hello\""
	//array     ["gold","silver","bronze"]
	//object    { "year"   : 1998
	//			"event"  : "archery"
	//			"medals" : ["gold","silver","bronze"]
	//}

	type Movie struct {
		title  string
		year   int  `json:"released"`
		color  bool `json:"color,omitempty"`
		actors []string
	}
	var movies = []Movie{
		{title: "casablance", year: 1942, color: false, actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
		{title: "uhsfuisuf", year: 1965, color: true, actors: []string{"efwef ffewf"}},
		{title: "wdqwdqd", year: 1985, color: true, actors: []string{"defdwfe bhrhbr", "Ingriwdfqwd Bergmfdqwfan"}},
	}
	// Go 对象和 JSON 的相互转换很容易，把 GO 的数据结构转换为 JSON 称为 marshal 。marshal是通过 json.Marshal 来实现的
	date, err := json.Marshal(movies)
	if err != nil {
		log.Fatal("JSON marshaling failed  : %s", err)
	}
	fmt.Println(date)
	// Marshal 生成一个字节 slice 其中包含一个不带有任何多余空白字符的很长的字符串，把生成的结果折叠一下放进去

	//为了方便阅读， Json.MarshalIndent 的变体可以输出整齐格式化的结果。
	//这个函数有两个参数，一个是定义每行输出的前缀字符串，另一个是定义缩进的字符串
	d, err := json.MarshalIndent(movies, "", "          ")
	if err != nil {
		log.Fatal("JSON marshaling failed  : %s", err)
	}
	fmt.Println("\n", d)
	// marshal 使用 Go 结构体成员的名称作为 JSON 对象里面字段的名称（反射  的方式）
	//成员标签定义可以是任意的  key："value"

	// marshal de 逆操作将 JSON 字符串解码为 GO 数据结构，叫 unmarshal ，这个由 json.Unmarshal 实现的
	var titles []struct{ title string }
	if err := json.Unmarshal(date, &titles); err != nil {
		log.Fatalf("JSON unmarshaling failed  : %s", err)
	}
	fmt.Println(titles)

	//很多 Web 服务都提供 JSON 接口，通过发送 HTTP 请求来获取想要得到的 JSON 信息
	//我们通过查询 GitHub 提供的 issue 跟踪接口来演示一下，首先定义需要的类型和常量

}

const IssuesUrl = "http://api.github.com/search/issues"

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}
type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Tiltle    string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string
}
type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

//即使对应的 JSON 字段的名称不是首字母大写，结构体成员名称也必须首字母大写
//函数 SearchIssues 发送 HTTP 请求并将回复解析为 JSON
//有些字符在 URL 中是特殊字符，（？   &），因此使用 url.QueryEscape 函数来确保它们拥有正确的含义

//查询 GitHub 的 issue 跟踪接口
func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, "   "))
	resp, err := http.Get(IssuesUrl + "?q=" + q)
	if err != nil {
		return nil, err
	}
	//我们必须在所有的可能分支上关闭 resp.Body
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}
	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil { //流式解码器（json.Decoder）
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}
