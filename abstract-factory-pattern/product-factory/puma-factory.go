package productfactory

import (
	"fmt"
	pant "git/go-design-pattern/abstract-factory-pattern/product/pant"
	shoe "git/go-design-pattern/abstract-factory-pattern/product/shoe"
)

type PumaFactory struct {
}

func (p *PumaFactory) CreateShoe() {
	pumaShoe := shoe.NewPumaShoe("puma", 42)
	fmt.Printf("%s --- %d\n", pumaShoe.GetBrand(), pumaShoe.GetSize())
}

func (p *PumaFactory) CreatePant() {
	pumaPant := pant.NewPumaPant("puma", "pant 323")
	fmt.Printf("%s --- %s\n", pumaPant.GetBrand(), pumaPant.GetName())
}
