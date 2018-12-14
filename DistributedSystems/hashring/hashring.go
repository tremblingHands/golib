package hashring

import(
    //"sync"
    "github.com/tremblingHands/golib/struct/rbtree"
    "strconv"
    "crypto/sha1"
    "fmt"
)

type vnode struct {
    value uint32 
    item map [string]bool
    parent  *node
}

type node struct {
    child []*vnode
}

type HashRing struct {
    vnodeArray *rbtree.Rbtree
    nodeArray map[string]*node
}

func newnode(id string) *node{
    return &node{}
}

func newvnode(value uint32 ) *vnode{
    return &vnode {
        value : value,
        item : make(map[string]bool),
    }
}

func (x *vnode) Cmp(y rbtree.Item) bool {
    return x.value < y.(*vnode).value
}

func NewHashRing() *HashRing {
    return &HashRing {
        vnodeArray : rbtree.NewRbtree(),
        nodeArray : make(map[string]*node),
    }
}

func getHash(id string) uint32 {
    h := sha1.New()
    h.Write([]byte(id))
    value := h.Sum(nil)
    return (uint32(value[3]) << 24) | (uint32(value[2]) << 16) | (uint32(value[1]) << 8) | (uint32(value[0])) 
}

func (r *HashRing) InsertNode(id string, weight int) {
    parent := newnode(id)
    if _, ok := r.nodeArray[id]; ok {
        return
    }
    r.nodeArray[id] = parent
    cnt := 0
    for cnt < weight {
        newid := id + ":" + strconv.Itoa(weight)
        value := getHash(newid)
        child := newvnode(value)
        cnt++
        if _, ok := r.vnodeArray.Search(child); ok {
            continue
        }
        r.vnodeArray.Insert(child)
        parent.child = append(parent.child, child)
    }

}

func (r *HashRing) GetNode(id string) {
    n := r.nodeArray[id]
    for _, child := range n.child {
        for it, _ := range child.item {
            fmt.Print(it)
        }
    }
    fmt.Println()
}

func (r *HashRing) DeleteNode(id string) {
    if  _, ok := r.nodeArray[id]; !ok  {
        return
    }
    parent := r.nodeArray[id]
    for _, child := range parent.child {
        next, flag := r.vnodeArray.Search(child)
        if flag == false {
            next = r.vnodeArray.Successor(next)
        }
        for s, _  :=  range child.item {
            next.Item.(*vnode).item[s] = true
        }
        r.vnodeArray.Delete(child)
    }
    delete(r.nodeArray, id)

}

func (r *HashRing) InsertItem(id string) {
    value := getHash(id)
    child, flag := r.vnodeArray.Search(newvnode(value))
    if flag == false {
        child = r.vnodeArray.Successor(child)
    }
    child.Item.(*vnode).item[id] = true 
}

func (r *HashRing) DeleteItem(id string) {
    value := getHash(id)
    child, flag := r.vnodeArray.Search(newvnode(value))
    if flag == false {
        child = r.vnodeArray.Successor(child)
    }
    child.Item.(*vnode).item[id] = false
}






