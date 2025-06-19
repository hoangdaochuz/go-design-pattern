package decoratorpattern

type PizzaDecorator struct {
	pizza Pizza
}

func NewPizzaDecorator(pizza Pizza) *PizzaDecorator {
	return &PizzaDecorator{
		pizza: pizza,
	}
}
func (pd *PizzaDecorator) GetDescription() string {
	return pd.pizza.GetDescription()
}

func (pd *PizzaDecorator) GetPrice() int32 {
	return pd.pizza.GetPrice()
}
