package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func getHelloGo(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	fmt.Println("request line: ", r.RequestURI)
	qs := r.URL.Query()
	fmt.Println("k-1 : ", qs.Get("k-1"))
	fmt.Println("k-2 : ", qs.Get("k-2"))
	fmt.Fprintln(w, "Hello 你好阿！")
}

func postJsonHelloGo(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	allBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("read all from req body fail", err)
		return
	}
	fmt.Println("req body: ", string(allBytes))

	answer := `{"status":"oak"}`
	w.Write([]byte(answer))

	//post form 如下：
	//r.ParseForm()
	//fmt.Println(r.PostForm) // 打印form数据
	//fmt.Println(r.PostForm.Get("name"), r.PostForm.Get("age"))
}

func main() {
	http.HandleFunc("/hello-go", getHelloGo)
	http.HandleFunc("/hello-go-post-json", postJsonHelloGo)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("fail....")
		return
	}
}
