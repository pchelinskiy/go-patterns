package prototype

import "fmt"

type File struct {
	name string
}

func (f *File) Print(indentation string) {
	fmt.Println(indentation + f.name)
}

func (f *File) Clone() INode {
	return &File{name: f.name + "_clone"}
}
