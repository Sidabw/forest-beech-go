# 一些资料
> 学了一遍又一遍，复习了一遍又一遍，关键是得用上啊。这是之前的资料

* https://www.topgoer.com/go%E5%9F%BA%E7%A1%80/Go%E8%AF%AD%E8%A8%80%E7%9A%84%E4%B8%BB%E8%A6%81%E7%89%B9%E5%BE%81.html
* http://docscn.studygolang.com/doc/
* https://go.dev/learn/
* https://go.dev/tour/methods/9
* https://go.dev/tour/list


#### why go
> Java有NIO，为啥要用go？说一下自己的理解
* Java只有JIT识别的热点代码才会转成机器码直接执行，其他都是解释执行
* GO直接编译成机器码。这是差距。
* 至于协程 VS NIO，这个争议还是比较大的
* GO的用处，还是在更快的扩容上吧，启动一个GO和冷启动一个Java，时间还是差别很多的
* 听过golang作为web服务器，qps轻轻松松上4w？需要测下
