package main

import (
	"fmt"
	"net/http"
)

func handler2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, Secure World!")
}

func main() {
	http.HandleFunc("/", handler2)                // 요청 처리 핸들러 등록
	fmt.Println("Starting HTTPS server on :8443") // HTTPS 서버 실행 로그
	err := http.ListenAndServeTLS(":8443", "cert.pem", "key.pem", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
