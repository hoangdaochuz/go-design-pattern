package abstractfactorypattern

type PumaShoe struct {
	BaseShoe
}

func NewPumaShoe(brand string, size int) IShoe {
	return &PumaShoe{
		BaseShoe: BaseShoe{
			brand: brand,
			size:  size,
		},
	}
}
