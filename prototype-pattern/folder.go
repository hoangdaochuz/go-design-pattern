package prototypepattern

import "fmt"

type Folder struct {
	name      string
	childrens []INode
}

func NewFolder(name string, children []INode) INode {
	return &Folder{
		name:      name,
		childrens: children,
	}
}

func (fo *Folder) Print() {
	fmt.Printf("folder: %s\n", fo.name)
	for _, item := range fo.childrens {
		item.Print()
	}
}

func (fo *Folder) Clone() INode {
	folderCloned := &Folder{}
	folderCloned.name = fo.name + "_cloned"
	var folderChildrenCloned []INode
	for _, item := range fo.childrens {
		clonedItem := item.Clone()
		folderChildrenCloned = append(folderChildrenCloned, clonedItem)
	}
	folderCloned.childrens = folderChildrenCloned
	return folderCloned
}
