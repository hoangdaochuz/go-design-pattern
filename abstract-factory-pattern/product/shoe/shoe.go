package abstractfactorypattern

type IShoe interface {
	GetBrand() string
	GetSize() int
}

type BaseShoe struct {
	brand string
	size  int
}

func NewBaseShoe(brand string, size int) IShoe {
	return &BaseShoe{
		brand: brand,
		size:  size,
	}
}

func (b *BaseShoe) GetBrand() string {
	return b.brand
}

func (b *BaseShoe) GetSize() int {
	return b.size
}
