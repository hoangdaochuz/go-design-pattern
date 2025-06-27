package abstractfactorypattern

type IPant interface {
	GetBrand() string
	GetName() string
}

type BasePant struct {
	brand string
	name  string
}

func NewBasePant(brand, name string) IPant {
	return &BasePant{
		brand: brand,
		name:  name,
	}
}

func (p *BasePant) GetBrand() string {
	return p.brand
}

func (p *BasePant) GetName() string {
	return p.name
}
