package main

import(
    "github.com/tremblingHands/golib/DistributedSystems/hashring"
)


func main(){
    hr := hashring.NewHashRing()
    hr.InsertNode("wahaha", 2)
    hr.InsertNode("lalala", 3)
    hr.InsertItem("aasd")
    hr.InsertItem("b213")
    hr.InsertItem("cwer")
    hr.InsertItem("dtr")
    hr.InsertItem("e34")
    hr.InsertItem("mbd")
    hr.InsertItem("qeqweqwe")
    hr.InsertItem("cdasd")
    println("wahaha:")
    hr.PrintNode("wahaha")
    println("lalala:")
    hr.PrintNode("lalala")
}
