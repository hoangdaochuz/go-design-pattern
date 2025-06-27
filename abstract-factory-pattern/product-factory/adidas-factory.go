package productfactory

import (
	"fmt"
	pant "git/go-design-pattern/abstract-factory-pattern/product/pant"
	shoe "git/go-design-pattern/abstract-factory-pattern/product/shoe"
)

type AdidasFactory struct {
}

func (a *AdidasFactory) CreateShoe() {
	adidasShoe := shoe.NewAdidasShoe("adidas", 42)
	fmt.Printf("%s --- %d\n", adidasShoe.GetBrand(), adidasShoe.GetSize())
}

func (a *AdidasFactory) CreatePant() {
	adidasPant := pant.NewAdidasPant("adidas", "pant 323")
	fmt.Printf("%s --- %s\n", adidasPant.GetBrand(), adidasPant.GetName())
}
