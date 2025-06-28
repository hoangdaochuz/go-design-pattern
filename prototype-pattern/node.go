package prototypepattern

type INode interface {
	Print()
	Clone() INode
}
