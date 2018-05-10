package main

import (
	"net/http"
	"log"
	"net/http/cookiejar"
	"net/http/httputil"
)

func main() {
	jar, err := cookiejar.New(nil)
	if err != nil {
		panic(err)
	}
	client := http.Client{
		Jar: jar,
	}
	// クッキーは初回アクセスでクッキーを受信し、2回目以降のアクセスでクッキーをサーバーに対して送信するような仕組みなので、二回アクセスします(?)
	// TODO: ここも文章の意味がよくわからなかった。
	for i := 0; i< 2; i++ {
		resp, err := client.Get("http://localhost:18888/cookie")
		if err != nil {
			panic(err)
		}
		dump, err := httputil.DumpResponse(resp, true)
		if err != nil {
			panic(err)
		}
		log.Println(string(dump))
	}
}
