package prototype

import "fmt"

type Folder struct {
	name   string
	childs []INode
}

func (f *Folder) Print(indentation string) {
	fmt.Println(indentation + f.name)
	for _, node := range f.childs {
		node.Print(indentation + indentation)
	}
}

func (f *Folder) Clone() INode {
	copy := &Folder{name: f.name + "_clone"}
	for _, node := range f.childs {
		copy.childs = append(copy.childs, node.Clone())
	}
	return copy
}
