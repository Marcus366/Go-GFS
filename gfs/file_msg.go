package gfs

import (
	"container/list"
)

type FileMsg struct {
	name   string
	blocks *list.List
}

func NewFileMsg(name string) *FileMsg {
	f := new(FileMsg)

	f.name = name
	f.blocks = list.New()

	return f
}
