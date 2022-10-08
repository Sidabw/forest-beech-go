package main

import "fmt"

//定义一个接口
//1。 接口： 方法集（a set of method）
//2。 对接口的implements是隐式的，也就是不用手动声明
//3。 Interface values【下面的：var a Abser】可以认为是一个具体值和具体类型的元祖。
//4。 nil receiver

type Abser interface {
	Abs() float64
}

type Myfloat float64

func (f Myfloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex4 struct {
	X int
	Y int
}

func (a Vertex4) Abs() float64 {
	if a.X < 0 {
		return float64(-a.X)
	}
	return float64(a.X)
}

type Vertex5 struct {
	X int
	Y int
}

func (b *Vertex5) Abs() float64 {
	if b == nil {
		fmt.Println("got a nil")
		return 1
	}

	if b.X < 0 {
		return float64(-b.X)
	}
	return float64(b.X)
}

func main() {
	var a Abser
	f := Myfloat(1.2)
	v5 := Vertex5{1, 2}
	v6 := &Vertex5{1, 2}

	a = f
	//以下编译报错，
	//a = v5
	//因为v5.Abs()是(&v5).Abs()的indirection，但接口不行
	//这样就好了
	a = &v5
	a = v6

	v2 := Vertex4{1, 3}
	a = v2
	fmt.Println("Vertex4 Abs()", a.Abs())
	//以上总结：
	//Myfloat implements Abser
	//*VerTex3 implements Abser
	//Vertex implements Abser

	//https://go.dev/tour/methods/11
	fmt.Printf("(%v, %T)\n", a, a)

	//https://go.dev/tour/methods/12
	var v7 *Vertex5
	a = v7
	a.Abs()
	fmt.Printf("(%v, %T)\n", a, a)

	//empty interface
	//类似Java的Object类，
	fmt.Println("//empty interface")
	var aa interface{}
	aa = 1
	fmt.Printf("(%v, %T)\n", aa, aa)
	aa = "aaa"
	//fmt.Println的参数列表就是一个empty interface
	fmt.Println(aa)

	fmt.Println("//Type assertions")
	//Type assertions
	var bb interface{} = "bbb"
	fmt.Printf("(%v, %T)\n", bb, bb)
	b2 := bb.(string) //这就叫Type assertions
	fmt.Printf("(%v, %T)\n", b2, b2)
	//这相当于是一个测试方法，上面是强转，失败就panic
	//这个失败b2Val则被赋零值
	b2Val, isStr := bb.(string)
	fmt.Println(b2Val, isStr)

}
