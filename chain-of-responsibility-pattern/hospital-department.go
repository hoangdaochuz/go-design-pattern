package chainofresponsibilitypattern

type Department interface {
	Execute(Patient)
	SetNext(Department)
}
