package main

import "fmt"

func main() {
	/*
	  1. 指针持有某个值的内存地址
	  2. 类型*T 标识 指向T的值的指针，零值为nil
	*/
	var p *int
	//会输出 <nil>
	fmt.Println(p)
	/*
		以下会报错，那这个*T有什么用呢
		*p = 2
		fmt.Println(*p)

		一般用于参数列表声明。要一个指针。
		i := 42
		p = &i
		fmt.Println(*p)

		如果非要这样用，可以加一行；new的作用是分配内存。make同理,只不过make只用于slice、map
		p = new(int)
		*p = 2
		fmt.Println("p = new(int)  ----->    ", *p)
	*/

	//标准用法：【全部是用在变量上】&拿到指针，*设置值，*读取值
	//&运算符生成指向其操作数的指针。
	i := 42
	var p2 = &i
	//直接输出该指针会打印其内存地址，如：0xc0000b2010
	fmt.Println(p2)

	//*运算符表示指针的基础值。打印42
	fmt.Println(*p2)

	//通过指针改变原值
	*p2 = 21
	fmt.Println(i)

}

/*
三种用法
1. &只用在变量前面，得到一个指针，如&i，那么标识生成一个指向操作数(i)的值的指针，直接输出：内存地址
2. *用在变量前面，表示读取指针的值 或 设置指针的值
3. *用在类型前面，得到一个指针，如*int ，那么标识生成一个指向类型的值的指针，直接输出：<nil>

补充：&用在类型前面，得到一个指针，如&T, 参见r2indirection

有关指针使用上的一些简写操作：
r2structs、
*/
