package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//直接返回hello world路径
		fmt.Fprintf(w, "Hello World!")
	})
	//可以发现,在设置路径为/time的时候,如果请求的url为/time/1的话,则因为/time代表的是叶子节点
	//只能匹配自己,所以被/节点锁匹配,返回hello world,在将其换成/time/之后,由于匹配的是子节点,所以
	//当访问/time/1的时候,匹配到的节点为/time/返回的为系统的当前时间;
	http.HandleFunc("/time/", func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		//输出对应的时间格式
		timeStr := fmt.Sprintf("{\"time\": \"%s\"}", t)
		w.Write([]byte(timeStr))
	})

	http.ListenAndServe(":8080", nil)
}
