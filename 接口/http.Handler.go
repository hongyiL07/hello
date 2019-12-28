package main

import (
	"fmt"
	"log"
	"net/http"
)

func main()  {
	db := database{"shoes":50,"socks":5}
	log.Fatal(http.ListenAndServe("localhost",db))

	mux:=http.NewServeMux()
	mux.Handle("/list",http.headlerFunc(db.list))
	mux.Handle("/price",http.headlerFunc(db.price))
	log.Fatal(http.ListenAndServe("localhost",mux))



	//服务器的主函数可以进一步简化
	//db := database{"shoes":50,"socks":5}
	//http.HeadlerFunc(db.list)
	///http.HeadlerFunc(db.price)
	//log.Fatal(http.ListenAndServe("localhost",nil))
}

type dollars float32

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f",d)
}

type database map[string]dollars        //用一个 map 类型（命名为 database）来代表仓库

//处理函数基于 URL 的路径部分（req.URL.Path ）来决定执行哪些逻辑
func (db database) ServeHTTP (w http.ResponseWriter,req *http.Request){
	switch req.URL.Path {
	case "/list":
		for item, price := range db {
			fmt.Fprintf(w, "%s :%s \n", item, price)
		}
	case "/price":
		item := req.URL.Query().Get("item")
		price, ok := db[item]
		if !ok {
			w.WriteHeader(http.StatusNotFound) //404
			fmt.Fprintf(w, "no such item : %q\n", item)
			return
		}
		fmt.Fprintf(w, "%s\n", price)
	default:
		w.WriteHeader(http.StatusNotFound)   //404
		fmt.Fprintf(w, "no such page : %s\n",rep.URl)

		//
		//msg := fmt.Sprintf("no such page : %s\n",rep.URl")
		//http.Error(w, msg, http.StatusNotFound)
	}
}




func (db database) list(w http.ResponseWriter,rep *http.Request) {
	for item, price := range db{
		fmt.Fprintf(w, "%s :%s\n",item,price)
	}
}
func (db database)price ((w http.ResponseWriter,rep *http.Request)  {
	item := req.URL.Query().Get("item")
	price , ok := db[item]
	if ok{
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item : %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}//调用 db.list 时，等价于以 db 为接收者调用 database.list 方法，所以 db.list 是一个实现了处理功能的函数