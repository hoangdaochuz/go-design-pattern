package decoratorpattern

type SeafoodPizza struct {
	PizzaDecorator
	// other fields
}

func NewSeafoodPizza(pizza Pizza) *SeafoodPizza {
	return &SeafoodPizza{
		PizzaDecorator: *NewPizzaDecorator(pizza),
	}
}

func (sp *SeafoodPizza) GetDescription() string {
	return sp.PizzaDecorator.GetDescription() + ", Seafood"
}

func (sp *SeafoodPizza) GetPrice() int32 {
	return sp.PizzaDecorator.GetPrice() + 200
}
