package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// 获取参数
	name := ps.ByName("name")
	fmt.Fprintf(w, "Hello, %s!\n", name)
}

func main() {
	// 通过httprouter创建路由
	r := httprouter.New()

	// 通过路由设置规则（注册函数）
	// http://localhost:8080/
	r.GET("/", Index)
	r.GET("/Hello/:name", Hello)

	// 启动
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
