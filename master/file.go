package master

import (
	"container/list"
)

type File struct {
	name   string
	chunks *list.List
}

func NewFile(name string) *File {
	f := new(File)

	f.name = name
	f.chunks = list.New()

	return f
}
