package decoratorpattern

// Base Pizza
type BasePizza struct {
	description string
	price       int32
}

// implement Pizza interface

func NewBasePizza(description string, price int32) Pizza {
	return &BasePizza{
		description: description,
		price:       price,
	}
}

func (bp *BasePizza) GetDescription() string {
	return bp.description
}

func (bP *BasePizza) GetPrice() int32 {
	return bP.price
}
