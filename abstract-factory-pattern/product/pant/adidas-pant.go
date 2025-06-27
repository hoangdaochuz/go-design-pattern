package abstractfactorypattern

type AdidasPant struct {
	BasePant
}

func NewAdidasPant(brand, name string) IPant {
	return &AdidasPant{
		BasePant: BasePant{
			brand: brand,
			name:  name,
		},
	}
}
