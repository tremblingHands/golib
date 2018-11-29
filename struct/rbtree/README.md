
自定义结构体接口实现如下

```go

type Item interface {
    Cmp(i Item) bool
}

type Int int

func (x Int) Cmp(y Item) bool {
    return x<y.(Int)
}


```


```go

package main

import (
	"fmt"
	"github.com/tremblingHands/golib/struct/rbtree"
)


func main() {
	rbt := rbtree.NewRbtree()

	m := 0
	n := 100
    x := 0
    a := []int{1,3,2,6,5,7,4,9,8,10}
    k := 0
	for m < n {
        if m %10 == 0 {
            x = a[k] * 10
            k++
        }
        rbt.Insert(rbtree.Int(x))
        x++
		m++
	}

    pp(rbt.Root)
    fmt.Println("BlackCount :", rbt.BlackCount())

	m = 0
	for m < n {
		if m%2 == 0 {
			rbt.Delete(rbtree.Int(m))
		}

		m++
	}

    pp(rbt.Root)
    fmt.Println("BlackCount :", rbt.BlackCount())
}

func pp(now *rbtree.Node) {
    if now.Item == nil {
        return
    }
    pp(now.Left)
    fmt.Println(*now, now.Parent)
    pp(now.Right)

}
```
