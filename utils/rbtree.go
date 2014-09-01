package utils

/*
* This is a Red-Black Tree impelementation. It has no document but it
* is just the same as CLRS description. You can easily get the message.
*/


const (
    RED = iota
    BLACK
)

type Comparable interface {
    CompareTo(other Comparable) int
}

type RBNode struct {
    Key   Comparable
    Value interface{}
    color int
    left, right, parent *RBNode
}

func NewNode(key Comparable, value interface{}) *RBNode {
    return &RBNode{key, value, RED, nil, nil, nil};
}

type RBTree struct {
    root, sentinel  *RBNode
}

func NewTree() *RBTree {
    sentinel := NewNode(nil, nil)
    return &RBTree{sentinel, sentinel}
}

func (rbt *RBTree) Find(key Comparable) interface{} {
    x := rbt.root
    for x != rbt.sentinel {
        if ret := key.CompareTo(x.Key); ret > 0 {
            x = x.right
        } else if ret < 0 {
            x = x.left
        } else {
            return x.Value
        }
    }
    return nil
}

func (rbt *RBTree) Insert(newnode *RBNode) {
    newnode.left = rbt.sentinel
    newnode.right = rbt.sentinel
    node, parent := rbt.root, rbt.sentinel
    for node != rbt.sentinel {
        parent = node
        if ret := newnode.Key.CompareTo(node.Key); ret > 0 {
            node = node.right
        } else if ret < 0 {
            node = node.left
        } else {
            node.Value = newnode.Value
            return
        }
    }
    newnode.parent = parent
    if parent == rbt.sentinel {
        rbt.root = newnode
    } else if newnode.Key.CompareTo(parent.Key) > 0 {
        parent.right = newnode
    } else {
        parent.left = newnode
    }
    rbt.insertFixup(newnode)
}

func (rbt *RBTree) Delete(n *RBNode) *RBNode {
    var x, y *RBNode
    if n.left == rbt.sentinel || n.right == rbt.sentinel {
        y = n
    } else {
        y = rbt.successor(n)
    }
    if y.left != rbt.sentinel {
        x = y.left
    } else {
        x = y.right
    }
    x.parent = y.parent
    if y.parent == rbt.sentinel {
        rbt.root = x
    } else if y == y.parent.left {
        y.parent.left = x
    } else {
        y.parent.right = x
    }
    if y != n {
        n.Key = y.Key
        n.Value = y.Value
    }
    if y.color == BLACK {
        rbt.deleteFixup(x)
    }
    return y
}

func (rbt *RBTree) insertFixup(n *RBNode) {
    for n.parent.color == RED {
        if n.parent == n.parent.parent.left {
            y := n.parent.parent.right
            if y.color == RED {
                n.parent.color = BLACK
                y.color = BLACK
                n.parent.parent.color = RED
                n = n.parent.parent
            } else {
                if n == n.parent.right {
                    n = n.parent
                    rbt.leftRotate(n)
                }
                n.parent.color = BLACK
                n.parent.parent.color = RED
                rbt.rightRotate(n)
            }
        } else {
            y := n.parent.parent.left
            if y.color == RED {
                n.parent.color = BLACK
                y.color = BLACK
                n.parent.parent.color = RED
                n = n.parent.parent
            } else {
                if n == n.parent.left {
                    n = n.parent
                    rbt.rightRotate(n)
                }
                n.parent.color = BLACK
                n.parent.parent.color = RED
                rbt.leftRotate(n)
            }
        }
    }
    rbt.root.color = BLACK
}

func (rbt *RBTree) deleteFixup(x *RBNode) {
    for x != rbt.root && x.color == BLACK {
        if x == x.parent.left {
            w := x.parent.right
            if w.color == RED {
                w.color = BLACK
                x.parent.color = RED
                rbt.leftRotate(x.parent)
                w = x.parent.right
            }
            if w.left.color == BLACK && w.right.color == BLACK {
                w.color = RED
                x = x.parent
            } else {
                if w.right.color == BLACK {
                    w.color = RED
                    rbt.rightRotate(w)
                    w = x.parent.right
                }
                w.color = x.parent.color
                x.parent.color = BLACK
                w.right.color = BLACK
                rbt.leftRotate(x.parent)
                x = rbt.root
            }
        } else {
            w := x.parent.left
            if w.color == RED {
                w.color = BLACK
                x.parent.color = RED
                rbt.rightRotate(x.parent)
                w = x.parent.left
            }
            if w.left.color == BLACK && w.right.color == BLACK {
                w.color = RED
                x = x.parent
            } else {
                if w.left.color == BLACK {
                    w.color = RED
                    rbt.leftRotate(w)
                    w = x.parent.left
                }
                w.color = x.parent.color
                x.parent.color = BLACK
                w.right.color = BLACK
                rbt.rightRotate(x.parent)
                x = rbt.root
            }
        }
    }
    x.color = BLACK
}
func (rbt *RBTree) leftRotate(x *RBNode) {
    y := x.right
    x.right = y.left
    if y.left != rbt.sentinel {
        y.left.parent = x
    }
    y.parent = x.parent
    if x.parent == rbt.sentinel {
        rbt.root = y
    } else if x == x.parent.left {
        x.parent.left = y
    } else {
        x.parent.right = y
    }
    y.left = x
    x.parent = y
}

func (rbt *RBTree) rightRotate(y *RBNode) {
    x := y.left
    y.left = x.right
    if x.right != rbt.sentinel {
        x.right.parent = y
    }
    x.parent = y.parent
    if y.parent == rbt.sentinel {
        rbt.root = x
    } else if y == y.parent.left {
        y.parent.left = x
    } else {
        y.parent.right = x
    }
    x.right = y
    y.parent = x
}

func (rbt *RBTree) successor(n *RBNode) (y *RBNode) {
    if n.right != rbt.sentinel {
        for y = n.right; y.left != rbt.sentinel; y = y.left {}
        return
    }
    y = n.parent
    for y != rbt.sentinel && y.right == n {
        n = y
        y = y.parent
    }
    return
}
