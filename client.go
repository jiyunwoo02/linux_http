package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	resp, err := http.Get("https://velog.io/@jiyunwoo/Go-vs.-Rust-1") // GET 요청
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close() // 응답 본문 닫기

	body, err := io.ReadAll(resp.Body) // 응답 본문 읽기
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	fmt.Println("Response Status:", resp.Status) // 응답 상태 출력
	fmt.Println("Response Body:", string(body))  // 응답 본문 출력

	/* 출력 예시
	Response Status: 200 OK
	Response Body: <!doctype html> ~~
	*/
}
