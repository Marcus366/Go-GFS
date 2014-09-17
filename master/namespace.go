package master

import (
    "fmt"
    "errors"

    "github.com/sysu2012zzp/Go-GFS/utils"
)

type Namespace struct {
   rootdir, rootfile *utils.RBTree
}

func NewNamespace() *Namespace {
    return &Namespace{ utils.NewTree(), utils.NewTree() }
}
