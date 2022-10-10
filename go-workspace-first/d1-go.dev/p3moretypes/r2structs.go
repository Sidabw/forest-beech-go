package main

import "fmt"

/*
结构体就是一组字段的集合
*/

//Vertex 定义一个结构体，名叫Vertex(顶点)，包含两个字段 X Y
type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	v.X = 3
	fmt.Println(v)
	fmt.Println(v.X)

	fmt.Println("----------------")

	var p = &v
	//这里是(*p).X 的简写
	p.X = 1e9
	fmt.Println(*p)

	fmt.Println("----------------")
	var (
		v1 = Vertex{1, 2}
		v2 = Vertex{X: 1} //Y is 0
		v3 = Vertex{}     //X:0 and Y:0
		v4 = &Vertex{3, 4}
	)
	fmt.Println(v1, v2, v3)
	fmt.Println(v4)
	fmt.Println(*v4)
}
