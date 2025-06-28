package builderpattern

import "fmt"

type HouseWoodBuilder struct {
	materialWall  string
	windowColor   string
	materialRoof  string
	materialFloor string
}

func NewHouseWood() HouseBuilder {
	return &HouseWoodBuilder{}
}

func (h *HouseWoodBuilder) BuildWall(material string) HouseBuilder {
	h.materialWall = material
	return h
}

func (h *HouseWoodBuilder) BuildRoof(materialRoof string) HouseBuilder {
	h.materialRoof = materialRoof
	return h
}

func (h *HouseWoodBuilder) BuildWindow(color string) HouseBuilder {
	h.windowColor = color
	return h
}

func (h *HouseWoodBuilder) BuildFloor(materialFloor string) HouseBuilder {
	h.materialFloor = materialFloor
	return h
}
func (h *HouseWoodBuilder) Build() *House {
	fmt.Printf("%s - %s - %s - %s \n", h.materialWall, h.materialFloor, h.materialRoof, h.windowColor)
	return &House{
		materialWall:  h.materialWall,
		windowColor:   h.windowColor,
		materialRoof:  h.materialRoof,
		materialFloor: h.materialFloor,
	}
}
