package master

import (
  "container/list"

type File struct {
  name string
  blocks *list.List
}

func NewFile(name string) {
  return &File{name:name, blocks:list.NewList() }
}
