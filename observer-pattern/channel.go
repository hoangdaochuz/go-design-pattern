package observerpattern

type Channel interface {
	Register(Subcriber)
	Unregister(Subcriber)
	Notify(string)
}
