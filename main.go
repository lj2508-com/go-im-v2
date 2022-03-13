package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/user/login", func(writer http.ResponseWriter, request *http.Request) {
		io.WriteString(writer, "测试接口")
	})
	http.ListenAndServe(":8090", nil)
}
