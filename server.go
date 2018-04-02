package main

import (
	"fmt"
	"github.com/urfave/negroni"
	"net/http"
	"router"
	"time"
)

func main() {

	n := negroni.Classic()
	// 添加全局handler
	n.UseHandler(handler())
	// 添加全局handlerFunc
	n.UseHandlerFunc(handlerFunc)
	// 添加全局 Func
	n.UseFunc(Func)
	// 挂载根路由
	n.UseHandler(router.RootRouter)
	// 监听端口
	n.Run(":3001")
}

/*
  全局中间件 handler
*/
func handler() http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type", "application/json")
		fmt.Println("<----middleware---handler-->")
		fmt.Println(r.Header.Get("User-Agent"))

	})
}

/*
  全局中间件 handlerFunc
*/
func handlerFunc(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("<----middleware---handlerFunc-->")
}

/*
  全局中间件 Func
*/
func Func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	then := time.Now()
	defer func() {
		if rec := recover(); rec != nil {
			fmt.Println("Recovered in f", rec)
			fmt.Fprintf(rw, "Recovered in f :%v", rec)
		}
		diff := time.Now().Sub(then)
		fmt.Println("req cost ", diff)
	}()
	fmt.Println("<----middleware---Func-->")
	next(rw, r)
}
