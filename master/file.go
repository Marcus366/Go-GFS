package master

import (
  "container/list"
)

type File struct {
  name string
  blocks *list.List
}

func NewFile(name string) *File {
  return &File {name:name, blocks:list.New() }
}
