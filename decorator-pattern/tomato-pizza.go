package decoratorpattern

type TomatoPizza struct { // extend Decorator <=> embedded Decorator in golang
	PizzaDecorator
	// other fields
}

func NewTomatoPizza(pizza Pizza) *TomatoPizza {
	return &TomatoPizza{
		PizzaDecorator: *NewPizzaDecorator(pizza),
	}
}

func (tp *TomatoPizza) GetDescription() string {
	return tp.PizzaDecorator.GetDescription() + ", Tomato"
}

func (tp *TomatoPizza) GetPrice() int32 {
	return tp.PizzaDecorator.GetPrice() + 100
}
