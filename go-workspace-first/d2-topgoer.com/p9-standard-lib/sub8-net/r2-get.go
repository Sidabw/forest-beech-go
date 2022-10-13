package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {

	apiUrl := "http://localhost:9090/hello-go"
	data := url.Values{}
	data.Set("k-1", "a-v")
	data.Set("k-2", "b-v")
	parsedUrl, err := url.ParseRequestURI(apiUrl)
	if err != nil {
		fmt.Println("fail on ParseRequestURI", err)
		return
	}
	parsedUrl.RawQuery = data.Encode()
	fmt.Println("url:: ", parsedUrl.String())

	getRes, err := http.Get(parsedUrl.String())
	if err != nil {
		fmt.Println("fail on get:: ", err)
		return
	}
	defer getRes.Body.Close()

	allBytes, err := ioutil.ReadAll(getRes.Body)
	if err != nil {
		return
	}
	fmt.Println("res::", string(allBytes))

}
