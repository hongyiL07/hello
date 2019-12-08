package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"reflect" //输出变量类型的包
)

func main() {
	for _, ur1 := range os.Args[1:] {
		resp, err := http.Get(ur1)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		b, err := ioutil.ReadAll(resp.Body)
		//var b string
		//io.Copy(b,resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s：%v\n", ur1, err)
			os.Exit(1)
		}
		q := 1
		fmt.Printf("%s", b)
		fmt.Println(reflect.TypeOf(b).Name())
		fmt.Println(q)
		fmt.Printf(reflect.TypeOf(q).Name())
	}
}
