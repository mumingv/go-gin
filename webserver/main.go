/**
 * 通过net/http包实现一个Web服务器
 *
 * http:// localhost:8080/hello
 * http:// localhost:8080/hello?name=mumingv&age=30&height=180
 */

package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	// 解析参数，默认是不解析的
	r.ParseForm()
	// k-v数据对
	fmt.Println(r.Form)

	fmt.Println("path = ", r.URL.Path)

	for k, v := range r.Form {
		fmt.Println("key: ", k)
		fmt.Println("val: ", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello Baby")
}

func main() {
	// 注册函数
	http.HandleFunc("/hello", sayHelloName)

	// 设置监听端口
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	// 服务启动后不会有任何打印
}
