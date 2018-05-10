package main

import (
	"net/http"
	"log"
	"net/url"
)

func main() {
	values := url.Values{
		"test": {"value"},
	}
	// POSTなのでEncodeせず、渡している。
	resp, err := http.PostForm("http://localhost:18888", values)
	if err != nil {
		panic(err)
	}
	log.Println("Status:", resp.Status)
	log.Println("Headers:", resp.Header)
	log.Println("Content-Length:", resp.Header.Get("Content-Length"))
}
