package main

import (
	"net/http"
	"context"
	"strconv"
)

func main() {
	//启动一个Web服务
	http.Handle("/", http.HandlerFunc(myHander))
	http.Handle("/getUser", http.HandlerFunc(userHander))
	http.ListenAndServe(":1234", nil)
}
func myHander(rw http.ResponseWriter, r *http.Request) {
	//模拟为Request附加值，这里附加了2个
	userContext := context.WithValue(context.Background(), "user", "张三")
	ageContext := context.WithValue(userContext, "age", 18)
	rContext := r.WithContext(ageContext)
	//这个模拟一个方法或者函数的调用，大部分情况下可能不在一个包里
	doHander(rw, rContext)
}
func doHander(rw http.ResponseWriter, r *http.Request) {
	//我们从这个Request里取出对应的值。
	user := r.Context().Value("user").(string)
	age := r.Context().Value("age").(int)
	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte("the user is " + user + ",age is " + strconv.Itoa(age)))
}
func userHander(rw http.ResponseWriter, r *http.Request) {
	rw.WriteHeader(200)
	rw.Write([]byte("user is hkk ,password is **********"))

}
