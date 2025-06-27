package abstractfactorypattern

type PumaPant struct {
	BasePant
}

func NewPumaPant(brand, name string) IPant {
	return &PumaPant{
		BasePant: BasePant{
			brand: brand,
			name:  name,
		},
	}
}
