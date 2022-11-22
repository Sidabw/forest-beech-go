package main

import (
	"encoding/json"
	"fmt"
)

//反引号，结构体标签。Tag是结构体的元信息，可以在运行的时候通过反射的机制读取出来

type Person struct {
	ID     int `json:"idd"`
	Gender string
	name   string //私有不能被json包访问
	Email  string `json:"-"` //表示json序列化时忽略该字段
}

func main() {
	//使用内置"encoding/json" 序列化/反序列化
	//原文中还有xml解析，未跟进
	//原文中还有更快更省空间的二进制json，需要msgpack库，未跟进
	//原文：https://www.topgoer.com/常用标准库/数据格式.html

	//part-1 standard
	s1 := Person{
		ID:     1,
		Gender: "nan",
		name:   "xxx",
		Email:  "a@bok.com",
	}
	data, err := json.Marshal(s1)
	if err != nil {
		return
	}
	fmt.Printf("json str: %s \n", data)

	//part-2
	//没啥用吧这个。。。
	data2, err := json.MarshalIndent(s1, "-pre-", "    ")
	if err != nil {
		return
	}
	fmt.Printf("json str2: %s \n", data2)

	//part-3
	m1 := make(map[string]interface{})
	m1["aa"] = "aa-v"
	m1["bb"] = 12
	m1["cc"] = true

	data3, _ := json.Marshal(m1)
	fmt.Printf("json str3: %s \n", data3)

	//part-4
	//反序列化到一个struct
	or1 := `{"id":1,"Gender":"nanaaaaaa"} `
	var or2 Person
	err = json.Unmarshal([]byte(or1), &or2)
	if err != nil {
		return
	}
	fmt.Printf("json str4:: %v \n", or2)

	//part-5
	//反序列化到一个interface{} 实际会自动变成map[string]interface
	var or3 interface{}
	err = json.Unmarshal([]byte(or1), &or3)
	if err != nil {
		return
	}
	fmt.Printf("json str5 %v \n", or3)
	or4 := or3.(map[string]interface{})
	for k, v := range or4 {
		switch v.(type) {
		case string:
			fmt.Printf("im string, k %s, v %v \n", k, v)
		case int:
			fmt.Printf("im int, k %s, v %v \n", k, v)
		case float64:
			//int 会转成 float64
			fmt.Printf("im float64, k %s, v %v \n", k, v)
		default:
			fmt.Printf("其他 \n")
		}
	}

}
