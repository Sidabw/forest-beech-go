# First 怎么打包
> 参考: 
> [打包,](https://blog.csdn.net/m0_67401228/article/details/126101509)
> [打包带图标](https://blog.csdn.net/weixin_40581279/article/details/121126842?spm=1001.2101.3001.6661.1&utm_medium=distribute.pc_relevant_t0.none-task-blog-2%7Edefault%7ECTRLIST%7ERate-1-121126842-blog-126101509.pc_relevant_aa_2&depth_1-utm_source=distribute.pc_relevant_t0.none-task-blog-2%7Edefault%7ECTRLIST%7ERate-1-121126842-blog-126101509.pc_relevant_aa_2&utm_relevant_index=1)

#### 打mac可执行文件
> Unix Executable File 双击运行，出terminal弹窗
* go env -w CGO_ENABLED==0 
* go env -w GOOS=darwin
* go env -w GOARCH=amd64
* go build -o main-mac main.go

#### 打windows可执行文件
* go env -w CGO_ENABLED==0
* go env -w GOOS=windows
* go env -w GOARCH=amd64
* go build -o main-windows.exe main.go

#### CGO_ENABLED==0 ？
* go env -w CGO_ENABLED=0
* 0表示静态编译，默认为1
* CGO是go语言调用c语言lib库。按照默认启用的话，打出来的可执行文件就会依赖客户电脑上的c库
* 交叉编译时，必须禁用CGO，即设置为0
  * 在mac上编译得到一个exe就叫做交叉编译
* [参考](https://blog.csdn.net/qq_43580193/article/details/120305231)