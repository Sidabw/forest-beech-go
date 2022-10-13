package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	url := "http://localhost:9090/hello-go-post-json"
	contentType := "application/json"
	data := `{"aaa":123, "cccc":"d卡卡卡"}`
	postRes, err := http.Post(url, contentType, strings.NewReader(data))
	if err != nil {
		fmt.Println("fail on post proceed ", err)
		return
	}
	defer postRes.Body.Close()

	allBytes, err := ioutil.ReadAll(postRes.Body)
	if err != nil {
		return
	}
	fmt.Println("res read: ", string(allBytes))
}
