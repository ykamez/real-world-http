package main

import (
	"net/http"
	"log"
	"bytes"
	"mime/multipart"
	"os"
	"io"
	"net/textproto"
)

// TODO: P77の内容が理解できていないので復習したい。
func main() {
	// マルチパートフォームの送信コンテンツ
	var buffer bytes.Buffer
	writer := multipart.NewWriter(&buffer)
	// ファイル以外のフィールドはこのメソッドを使って登録する。
	writer.WriteField("name", "Michael Jackson")

	//MIMEタイプ設定をしないバージョン=============
	// ファイル書き込みのwriterを作る。
	//fileWriter, err := writer.CreateFormFile("thumbnail", "photo.jpg")
	//if err != nil {
	//	panic(err)
	//}
	//readFile, err := os.Open("photo.jpg")
	//if err != nil {
	//	panic(err)
	//}
	//defer readFile.Close()
	//io.Copy(fileWriter, readFile)
	//================

	//==================
	// こっちではマルチパートの構成要素であるパートの作成処理を取り出して、任意のContent-Typeを指定できるようにしている。
	part := make(textproto.MIMEHeader)
	part.Set("Content-Type", "image/jpeg")
	part.Set("Content-Disposition", `form-data; name="thumbnail"; filename="photo.jpg"`)
	fileWriter, err := writer.CreatePart(part)
	if err != nil {
		panic(err)
	}
	readFile, err := os.Open("photo.jpg")
	if err != nil {
		panic(err)
	}
	io.Copy(fileWriter, readFile)
	//=====================

	writer.Close()
	// 二つ目の引数であるContent-Typeにはバウンダリーの文字列を入れている。"multipart/form-data; boundary=" + writer.Boundary()と書いても同じ結果が得られる。
	resp, err := http.Post("http://localhost:18888", writer.FormDataContentType(), &buffer)
	if err != nil {
		panic(err)
	}
	log.Println("Status:", resp.Status)
	log.Println("Headers:", resp.Header)
	log.Println("Content-Length:", resp.Header.Get("Content-Length"))
}
