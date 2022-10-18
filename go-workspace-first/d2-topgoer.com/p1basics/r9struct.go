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
}

func main() {

	s1 := Person{
		ID:     1,
		Gender: "nan",
		name:   "xxx",
	}
	data, err := json.Marshal(s1)
	if err != nil {
		return
	}
	fmt.Printf("json str: %s \n", data)

}
