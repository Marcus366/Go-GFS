package gfs

import (
	"container/list"
)

type FileMsg struct {
	name   string
	chunks *list.List
}

func NewFileMsg(name string) *FileMsg {
	f := new(FileMsg)

	f.name = name
	f.chunks = list.New()

	return f
}
