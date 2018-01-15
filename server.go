package main

import (
	"github.com/urfave/negroni"
	"router"
	"net/http"
	"fmt"
)

func main() {

	n := negroni.Classic()
	// 添加全局handler
	n.UseHandler(handler())
	// 挂载根路由
	n.UseHandler(router.Routes)
	// 监听端口
	n.Run(":3001")
}

/*
  全局中间件
*/
func handler() http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		fmt.Println("<----middleware----->")
		fmt.Println(r.Header.Get("User-Agent"))

	})
}
