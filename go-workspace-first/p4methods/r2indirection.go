package main

import "fmt"

/*
补充：
pointer indirection，怎么翻译？迂回指针？说的是"(&v).Scale(2)的缩写"吧？
对于function而言，参数列表要值必须给值，要指针必须给指针，否则编译错误
对于method而言，都可以
使用‘指针接受器（Pointer Receiver）’的原因：1. 改变对象的值 2.避免每次方法调用的时候都进行一次值拷贝
*/

type Vertex3 struct {
	X, Y float64
}

/*
Scale
定义一个method
*/
func (v *Vertex3) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

/*
ScaleFunc
定义一个function
*/
func ScaleFunc(v *Vertex3, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex3{2, 3}
	//这里是(&v).Scale(2)的缩写，method特有的
	v.Scale(2)
	//我要一个指针，你给我一个指针，没毛病
	ScaleFunc(&v, 10)
	fmt.Println(v)

	//创建一个指针
	p := &Vertex3{3, 4}
	//我要一个指针，你给我一个指针，没毛病
	p.Scale(2)
	//我要一个指针，你给我一个指针，没毛病
	ScaleFunc(p, 8)
	fmt.Println("指针p：", p)
	fmt.Println("值p：", *p)
}
