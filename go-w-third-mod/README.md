# go mod 学习笔记
>[参考1](https://zhuanlan.zhihu.com/p/482014524)
* src pkg bin 应该是使用原来GOPATH的结构。现在都用go.mod就不用这个结构了吧
* GO111MODULE=auto，默认。如果当前文件在包含go.mod文件的目录下面，则按照GO111MODULE=on处理
* GO111MODULE=on，go命令行会使用modules，而一点也不会去GOPATH目录下查找。
* 当modules功能启用时，依赖包的存放位置变更为$GOPATH/mog/pkg