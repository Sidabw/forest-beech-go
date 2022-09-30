package main

import "fmt"

type Vertex2 struct {
	l1, l2 float64
}

//声明一个map，key的类型是string，value的类型是Vertex2
var m map[string]Vertex2

func main() {
	//初始化一个map
	m = make(map[string]Vertex2)
	//赋值
	m["a"] = Vertex2{
		1.2, 3.4,
	}
	fmt.Println(m["a"])

	//不用make创建map
	var m2 = map[string]Vertex2{
		"a1": Vertex2{1.2, 4.4},
		"a2": Vertex2{333.222, 12123.456},
	}
	fmt.Println(m2)

	//简写
	var m3 = map[string]Vertex2{
		"a3": {1.1, 2.2},
		"a4": {3.3, 4.4},
	}
	fmt.Println(m3)

	m4 := make(map[string]int)
	//map put
	m4["an"] = 1
	//map get
	fmt.Println(m4["an"])
	//map del
	delete(m4, "an")
	//map contains
	v, ok := m4["an"]
	fmt.Println("The Value:", v, "Present?", ok)
}
