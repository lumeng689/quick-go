package test

import (
	"fmt"
	"log"
	"net/http"
)

type Hello struct {
}

func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello")
}

// 模块中需要导出的方法，首字母必须大写
func TestHttp() {

	fmt.Println("start server......")

	var h Hello

	err := http.ListenAndServe("localhost:4000", h)

	if err != nil {
		log.Fatal(err)
	}
}
