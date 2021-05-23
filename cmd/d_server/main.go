package main

import "net/http"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("欢迎调用C服务"))
	})
	http.ListenAndServe(":3002", nil)
}
