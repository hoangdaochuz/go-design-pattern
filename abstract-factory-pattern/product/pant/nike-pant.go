package abstractfactorypattern

type NikePant struct {
	BasePant
}

func NewNikePant(brand, name string) IPant {
	return &NikePant{
		BasePant: BasePant{
			brand: brand,
			name:  name,
		},
	}
}
