# Go Tour

### Tips
* go函数、变量、常量、自定义类型、包命名方式
  *  var（声明变量）, const（声明常量）, type（声明类型） ,func（声明函数）。
* 可见行
  * 声明在函数内部，是函数的本地值
  * 声明在函数外部，是对当前包可见(包内所有.go文件都可见)的全局值
  * 声明在函数外部且首字母大写是所有包可见的全局值.换句话说：
  * 在导入一个包时，你只能引用其中已导出的名字。任何“未导出”的名字在该包外均无法访问。
  * 已导出的都是大写字母开头，例如fmt.Println、math、Pi
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

### 阅读顺序
* packages.go
* imports.go
* functions.go
* variables.go
* type-basics.go
* type-conversions.go
* type-inference.go
* numberic-constants.go

