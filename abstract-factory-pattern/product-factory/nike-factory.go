package productfactory

import (
	"fmt"
	pant "git/go-design-pattern/abstract-factory-pattern/product/pant"
	shoe "git/go-design-pattern/abstract-factory-pattern/product/shoe"
)

type NikeFactory struct {
}

func (n *NikeFactory) CreateShoe() {
	nikeShoe := shoe.NewNikeShoe("nike", 42)
	fmt.Printf("%s --- %d\n", nikeShoe.GetBrand(), nikeShoe.GetSize())
}

func (n *NikeFactory) CreatePant() {
	nikePant := pant.NewNikePant("nike", "pant 323")
	fmt.Printf("%s --- %s\n", nikePant.GetBrand(), nikePant.GetName())
}
