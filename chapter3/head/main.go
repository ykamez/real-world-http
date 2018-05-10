package main

import (
	"net/http"
	"log"
)

func main() {
	// HEADメソッドを使っているため、bodyを読み込んだとしても長さゼロのバイト配列が返ってくる。
	resp, err := http.Head("http://localhost:18888")
	if err != nil {
		panic(err)
	}
	log.Println("Status:", resp.Status)
	log.Println("Headers:", resp.Header)
}
