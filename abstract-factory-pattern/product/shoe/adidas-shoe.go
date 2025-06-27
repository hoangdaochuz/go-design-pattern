package abstractfactorypattern

type AdidasShoe struct {
	BaseShoe
}

func NewAdidasShoe(brand string, size int) IShoe {
	return &AdidasShoe{
		BaseShoe: BaseShoe{
			brand: brand,
			size:  size,
		},
	}
}
