package rbtree
/*
import(
    "fmt"
)
*/
const(
    Red int = iota
    Black
)

type Item interface {
    Cmp(i Item) bool
}

type Int int

//函数参数为 Item
func (x Int) Cmp(y Item) bool {
    return x<y.(Int)
}

type Node struct {
    Left *Node
    Right *Node
    Parent *Node
    Color int
    Item
}

type Rbtree struct {
    Root *Node
    Nil *Node
}
/*
func PP(x *Node) {
    if m, ok := (x.Item).(Int); ok {
        print(m)
    }
}
*/
func Less(x, y *Node) bool {
    f := x.Cmp(y.Item)
    /*
    PP(x)
    PP(y)
    fmt.Println(f)
    */
    return f
}

func NewRbtree() *Rbtree {
    node := &Node{
         Color: Black,
    }
    return &Rbtree{
        Nil: node,
        Root: node,
    }
}

func (t *Rbtree) LeftRotate(x *Node) {
    if x.Right == t.Nil {
        return
    }
    y := x.Right
    p := x.Parent

    if p == t.Nil {
        t.Root = y
    }else if x == p.Left {
        p.Left = y
    }else {
        p.Right = y
    }
    y.Parent = p

    if y.Left != t.Nil {
        y.Left.Parent = x
    }
    x.Right = y.Left
    y.Left = x
    x.Parent = y
}

func (t *Rbtree) RightRotate(x *Node) {
    if x.Left == t.Nil {
        return
    }
    y := x.Left
    p := x.Parent

    if p == t.Nil {
        t.Root = y
    }else if x == p.Left {
        p.Left = y
    }else {
        p.Right = y
    }
    y.Parent = p

    if y.Right != t.Nil {
        y.Right.Parent = x
    }
    x.Left = y.Right
    y.Right = x
    x.Parent = y
}

func (t *Rbtree) Insert(item Item) *Node {
    if item == nil {
        return nil
    }
    return(t.insert(&Node{t.Nil,t.Nil,t.Nil,Red,item}))
}

func (t *Rbtree) insert(x *Node) *Node{
    //fmt.Println("insert x", *x)
    now := t.Root
    p := t.Nil
    for now != t.Nil {
        p = now
        if Less(x, now) {
            now = now.Left
        }else if Less(now, x){
            now = now.Right
        }else {
            return now
        }
    }

    x.Parent = p
    /*
    PP(x)
    PP(p)
    println("before")
    */
    if p == t.Nil {
        t.Root = x
    // 当节点没有子节点时，左右子节点都是 t.Nil，因此不能直接判断 now == p.Left
    }else if Less(x, p) {
        p.Left = x
    }else {
        p.Right = x
    }
    
    //fmt.Println("x.Parent", x.Parent)
    t.InsertFixUp(x)
    /**
    PP(x)
    PP(p)
    println("after")
    fmt.Println("finish insert x", *x)
    fmt.Println("x.Parent", x.Parent)
    fmt.Println()
    **/
    return x
}

func (t *Rbtree) InsertFixUp(x *Node) {
    for x.Parent.Color == Red {
        if x.Parent == x.Parent.Parent.Left {
            // case 1
            y := x.Parent.Parent.Right
            if y.Color == Red {
                x.Parent.Color = Black
                y.Color = Black
                x.Parent.Parent.Color = Red
                x = x.Parent.Parent
            }else {
                // case 2
                if x == x.Parent.Right {
                    x = x.Parent
                    t.LeftRotate(x)
                }
                // case 3
                x.Parent.Color = Black
                x.Parent.Parent.Color = Red
                t.RightRotate(x.Parent.Parent)
            }
        }else {
            // case 1
            y := x.Parent.Parent.Left
            if y.Color == Red {
                x.Parent.Color = Black
                y.Color = Black
                x.Parent.Parent.Color = Red
                x = x.Parent.Parent
            }else {
                // case 2
                if x == x.Parent.Left {
                    x = x.Parent
                    t.RightRotate(x)
                }
                // case 3
                x.Parent.Color = Black
                x.Parent.Parent.Color = Red
                t.LeftRotate(x.Parent.Parent)
            }

        }

    }
    t.Root.Color = Black
}

func (t *Rbtree) Min(x *Node) *Node{
    if x == t.Nil {
        return t.Nil
    }
    
    for x.Left != t.Nil {
        x = x.Left
    }
    return x
}

func (t *Rbtree) Max(x *Node) *Node{
    if x == t.Nil {
        return t.Nil
    }
    
    for x.Right != t.Nil {
        x = x.Right
    }
    return x
}

func (t *Rbtree) Search(x *Node) *Node {
    p := t.Root
    for p != t.Nil {
        if Less(x, p) {
            p = p.Left
        }else if Less(p, x){
            p = p.Right
        }else {
            break
        }
    }
    return p
}

func (t *Rbtree) Successor(x *Node) *Node {
    if x == t.Nil {
        return t.Nil
    }
    if x.Right != t.Nil {
        return t.Min(x.Right)
    }
    y := x.Parent
    for y != t.Nil && x == y.Right {
        x = y
        y = y.Parent
    }
    return y
}

func (t *Rbtree) Desuccessor(x *Node) *Node {
    if x == t.Nil {
        return t.Nil
    }
    if x.Left != t.Nil {
        return t.Max(x.Left)
    }
    y := x.Parent
    for y != t.Nil && x == y.Left {
        x = y
        y = y.Parent
    }
    return y
}

func (t *Rbtree) Delete(key Item) *Node {
    if key == nil {
        return nil
    }
    return t.delete(&Node{t.Nil,t.Nil,t.Nil,Red,key})
}


func (t *Rbtree) delete(key *Node) *Node {
    z := t.Search(key)
    if z == t.Nil {
        return t.Nil
    }
    ret := &Node{t.Nil, t.Nil, t.Nil, z.Color, z.Item}
    var y, x *Node

    if z.Left == t.Nil || z.Right == t.Nil {
        y = z
    }else {
        y = t.Successor(z)
    }
    
    if y.Left != t.Nil {
        x = y.Left
    }else{
        x = y.Right
    }
    x.Parent = y.Parent

    if y.Parent == t.Nil {
        t.Root = x
    }else if y == y.Parent.Left {
        y.Parent.Left = x
    }else {
        y.Parent.Right = x
    }
    if y != z {
		z.Item = y.Item
	}

	if y.Color == Black {
		t.DeleteFixUp(x)
	}


    return ret

}

func (t *Rbtree) DeleteFixUp(x *Node) {
    for x != t.Root && x.Color == Black {
        if x == x.Parent.Left {
            y := x.Parent.Right
            //case 1
            if y.Color == Red {
                x.Parent.Color = Red
                y.Color = Black
                t.LeftRotate(x.Parent)
                y = x.Parent.Right
            }
            //case 2
            if y.Left.Color == Black && y.Right.Color == Black{
                y.Color = Red
                x = x.Parent
                //case 3
            }else if y.Left.Color == Red && y.Right.Color == Black {
                y.Color = Red
                y.Left.Color = Black
                t.RightRotate(y)
                //case 4
            }else {
                y.Color = x.Parent.Color
                x.Parent.Color = Black
                y.Right.Color = Black
                t.LeftRotate(x.Parent)

                x = t.Root
            }
        }else {
            y := x.Parent.Left
            if y.Color == Red {
                y.Color = Black
                x.Parent.Color = Red
                t.RightRotate(x.Parent)
                y = x.Parent.Left
            }
            if y.Left.Color == Black && y.Right.Color == Black {
                y.Color = Black
                x = x.Parent
            }else if y.Left.Color == Black && y.Right.Color == Red {
                y.Color = Red
                y.Right.Color = Black
                t.LeftRotate(y)
            }else {
                y.Color = x.Parent.Color
                x.Parent.Color = Black
                y.Left.Color = Black
                t.RightRotate(x.Parent)

                x = t.Root
            }
        }
    }
    x.Color = Black

}

func (t *Rbtree) BlackCount() int{
    p := t.Root
    cnt := 0
    for p != t.Nil{
        if p.Color == Black {
            cnt++
        }
        p = p.Left
    }
    return cnt
}



