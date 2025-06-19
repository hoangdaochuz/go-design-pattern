package decoratorpattern

type CheesePizza struct {
	PizzaDecorator
	// other fields
}

func NewCheesePizza(pizza Pizza) *CheesePizza {
	return &CheesePizza{
		PizzaDecorator: *NewPizzaDecorator(pizza),
		// other fields
	}
}

func (cp *CheesePizza) GetDescription() string {
	return cp.PizzaDecorator.GetDescription() + ", Cheese"
}

func (cp *CheesePizza) GetPrice() int32 {
	return cp.PizzaDecorator.GetPrice() + 150
}
