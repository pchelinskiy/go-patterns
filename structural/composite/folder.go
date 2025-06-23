package composite

import "fmt"

type Folder struct {
	components []Component
	name       string
}

func (f *Folder) Search(keyword string) {
	fmt.Printf("serching recursively for keyword %s in folder %s\n", keyword, f.name)
	for _, composite := range f.components {
		composite.Search(keyword)
	}
}

func (f *Folder) Add(c Component) {
	f.components = append(f.components, c)
}
