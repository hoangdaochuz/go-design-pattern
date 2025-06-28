package prototypepattern

import (
	"fmt"
)

type File struct {
	name string
}

func NewFile(name string) INode {
	return &File{
		name: name,
	}
}

func (f *File) Print() {
	fmt.Printf("file: %s\n", f.name)
}

func (f *File) Clone() INode {
	return &File{
		name: f.name + "_cloned",
	}
}
