package master

import (
	"errors"
	//"fmt"
	"strings"
)

type Directory struct {
	subdir map[string]*Directory
	files  map[string]*File
}

func NewDirectory() *Directory {
	return &Directory{make(map[string]*Directory), make(map[string]*File)}
}

func (d *Directory) recursiveFindDirectory(subpath string) *Directory {
	slice := strings.SplitN(subpath, "/", 2)
	subdir := d.subdir[slice[0]]
	if subdir != nil {
		return subdir.recursiveFindDirectory(slice[1])
	}
	return nil
}

type Namespace struct {
	rootdir *Directory
}

func NewNamespace() *Namespace {
	return &Namespace{NewDirectory()}
}

func (ns *Namespace) findFile(path string) (*File, error) {
	lastSlash := strings.LastIndex(path, "/")
	filename := path
	d := ns.rootdir
	if lastSlash != -1 {
		//slice := strings.Split(path, "/")
		d = ns.rootdir.recursiveFindDirectory(string(path[0:lastSlash]))
		filename = string(path[lastSlash+1:])
		if d == nil {
			return nil, errors.New("No Such File or Directory")
		}
	}

	msg := d.files[filename]
	if msg == nil {
		return nil, errors.New("No Such File")
	}

	return msg, nil
}

func (ns *Namespace) createFile(path string, flag int, perm uint32) (*File, error) {
	lastSlash := strings.LastIndex(path, "/")
	if lastSlash != -1 {
		slice := strings.Split(path, "/")
		d := ns.rootdir.recursiveFindDirectory(string(path[0:lastSlash]))
		if d == nil {
			return nil, errors.New("No Such File of Directory")
		}
		filename := slice[len(slice)-1]
		file := NewFile(filename)
		d.files[filename] = file
		return file, nil
	} else {
		file := NewFile(path)
		ns.rootdir.files[path] = file
		return file, nil
	}
}
