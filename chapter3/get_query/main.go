package main

import (
	"net/http"
	"io/ioutil"
	"log"
	"net/url"
)

func main() {
	values := url.Values{
		"query": {"hello, world"},
	}

	resp, _ := http.Get("http://localhost:18888" + "?" + values.Encode())
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	log.Println(string(body))
	log.Println("Headers:", resp.Header)
	log.Println("Content-Length:", resp.Header.Get("Content-Length"))
}
