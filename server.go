package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!") // 응답 메시지 작성
}

func main() {
	http.HandleFunc("/", handler)             // 요청 경로와 핸들러 연결
	fmt.Println("Server is running on :8080") // 서버 시작 로그 출력
	http.ListenAndServe(":8080", nil)         // HTTP 서버 실행
}
