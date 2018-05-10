package main

import (
	"net/http"
	"log"
	"strings"
)

func main() {

	reader := strings.NewReader("ホゲホゲ")
	//reader, err := os.Open("main.go")
	//if err != nil {
	//	panic(err)
	//}

	// Postの3つ目の引数には送信する内容を、io.Reader形式で渡す。
	resp, err := http.Post("http://localhost:18888", "text/plain", reader)
	if err != nil {
		panic(err)
	}
	log.Println("Status:", resp.Status)
	log.Println("Headers:", resp.Header)
	log.Println("Content-Length:", resp.Header.Get("Content-Length"))
}
