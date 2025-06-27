package abstractfactorypattern

type NikeShoe struct {
	BaseShoe
}

func NewNikeShoe(brand string, size int) IShoe {
	return &NikeShoe{
		BaseShoe: BaseShoe{
			brand: brand,
			size:  size,
		},
	}
}
