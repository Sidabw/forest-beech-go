## before-all
* go函数、变量、常量、自定义类型、包命名方式
    *  var（声明变量）, const（声明常量）, type（声明类型） ,func（声明函数）。
* 可见行
    * 声明在函数内部，是函数的本地值
    * 声明在函数外部，是对当前包可见(包内所有.go文件都可见)的全局值
    * 声明在函数外部且首字母大写是所有包可见的全局值.换句话说：
    * 在导入一个包时，你只能引用其中已导出的名字。任何“未导出”的名字在该包外均无法访问。
    * 已导出的都是大写字母开头，例如fmt.Println、math.Pi
* Go语言四种声明方式
    * var（声明变量）, const（声明常量）, type（声明类型） ,func（声明函数）。
    * 另外文件第一行的package xxx声明，用来说明当前文件属于哪个包。
* 类型
    * 值类型
        * bool
        * int(32 or 64), int8, int16, int32, int64
        * uint(32 or 64), uint8(byte), uint16, uint32, uint64
        * float32, float64
        * string
        * complex64, complex128
        * array    -- 固定长度的数组
    * 引用类型
        * slice   -- 序列数组(最常用)
        * map     -- 映射
        * chan    -- 管道
* 项目结构
    * src 源码
    * pkg 包文件，源码编译完成后的文件的位置。
    * bin 二进制。

## 内置函数 
* append          -- 用来追加元素到数组、slice中,返回修改后的数组、slice
* close           -- 主要用来关闭channel
* delete            -- 从map中删除key对应的value
* panic            -- 停止常规的goroutine  （panic和recover：用来做错误处理）
* recover         -- 允许程序定义goroutine的panic动作
* real            -- 返回complex的实部   （complex、real imag：用于创建和操作复数）
* imag            -- 返回complex的虚部
* make            -- 用来分配内存，返回Type本身(只能应用于slice, map, channel)
* new                -- 用来分配内存，主要用来分配值类型，比如int、struct。返回指向Type的指针
* cap                -- capacity是容量的意思，用于返回某个类型的最大容量（只能用于切片和 map）
* copy            -- 用于复制和连接slice，返回复制的数目
* len                -- 来求长度，比如string、array、slice、map、channel ，返回长度
* print、println     -- 底层打印函数，在部署环境中建议使用 fmt 包

## 内置接口
```
//只要实现了Error()函数，返回值为String的都实现了err接口
type error interface { 
        Error()    String
}
//Stringer同理
```

## init 函数和main函数
##### init
*  1 init函数是用于程序执行前做包的初始化的函数，比如初始化包里的变量等
*  2 每个包可以拥有多个init函数
*  3 包的每个源文件也可以拥有多个init函数
*  4 同一个包中多个init函数的执行顺序go语言没有明确的定义(说明)
*  5 不同包的init函数按照包导入的依赖关系决定该初始化函数的执行顺序
*  6 init函数不能被其他函数调用，而是在main函数执行之前，自动被调用
##### main
* 默认入口函数(主函数)：func main()
##### 异同
* 两个函数在定义时不能有任何的参数和返回值，且Go程序自动调用。
* init可以应用于任意包中，且可以重复定义多个。
* main函数只能用于main包中，且只能定义一个

## 命令
* go env用于打印Go语言的环境信息。
* go run命令可以编译并运行命令源码文件。
* go get可以根据要求和实际情况从互联网上下载或更新指定的代码包及其依赖包，并对它们进行编译和安装。
  * go get github.com/jmoiron/sqlx
  * 下载下来的包则会位于${GOPATH}/src/github.com/jmoiron/sqlx
* go build命令用于编译我们指定的源码文件或代码包以及它们的依赖包。
* go install用于编译并安装指定的代码包及它们的依赖包。
* go clean命令会删除掉执行其它命令时产生的一些文件和目录。
* go doc命令可以打印附于Go语言程序实体上的文档。我们可以通过把程序实体的标识符作为该命令的参数来达到查看其文档的目的。
* go test命令用于对Go语言编写的程序进行测试。
* go list命令的作用是列出指定的代码包的信息。
* go fix会把指定代码包的所有Go语言源码文件中的旧版本代码修正为新版本的代码。
* go vet是一个用于检查Go语言源码中静态错误的简单工具。
* go tool pprof命令来交互式的访问概要文件的内容。

## 运算符
...

## 下划线
* import _ "./hello" 仅执行hello包的init函数，仅仅。hello包下的都用不了
* n, _ := f.Read(buf) 占位，表示另外一个我不要了
  * 这个也叫做匿名变量

## 变量和常量
* 函数外的每个语句都必须以关键字开始（var、const、func等）
* :=不能使用在函数外。【短变量声明】
* iota常量计数器，使用参考：r1iota.go