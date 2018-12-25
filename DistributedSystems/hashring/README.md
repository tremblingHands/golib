
# 红黑树实现一致性哈希算法

 示例：

```go

package main

import(
    "github.com/tremblingHands/golib/DistributedSystems/hashring"
)


func main(){

//创建哈希环
    hr := hashring.NewHashRing()

//添加节点
    hr.InsertNode("wahaha", 2)
    hr.InsertNode("lalala", 3)

//插入资源
    hr.InsertItem("aasd")
    hr.InsertItem("b213")
    hr.InsertItem("cwer")
    hr.InsertItem("dtr")
    hr.InsertItem("e34")
    hr.InsertItem("mbd")
    hr.InsertItem("qeqweqwe")
    hr.InsertItem("cdasd")

//输出各个节点存储的资源
    println("wahaha:")
    hr.PrintNode("wahaha")
    println("lalala:")
    hr.PrintNode("lalala")
}
```
