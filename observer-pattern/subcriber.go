package observerpattern

type Subcriber interface {
	Update(string)
	GetId() string
}
