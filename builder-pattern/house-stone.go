package builderpattern

import "fmt"

type HouseStoneBuilder struct {
	materialWall  string
	windowColor   string
	materialRoof  string
	materialFloor string
	statueNumber  int
}

func NewHouseStone() HouseBuilder {
	return &HouseStoneBuilder{}
}

func (h *HouseStoneBuilder) BuildWall(materialWall string) HouseBuilder {
	h.materialWall = materialWall
	return h
}
func (h *HouseStoneBuilder) BuildRoof(materialRoof string) HouseBuilder {
	h.materialRoof = materialRoof
	return h
}
func (h *HouseStoneBuilder) BuildWindow(windowColor string) HouseBuilder {
	h.windowColor = windowColor
	return h
}

func (h *HouseStoneBuilder) BuildFloor(materialFloor string) HouseBuilder {
	h.materialFloor = materialFloor
	return h
}

func (h *HouseStoneBuilder) BuildStatue(number int) HouseBuilder {
	h.statueNumber = number
	return h
}

func (h *HouseStoneBuilder) Build() *House {
	h.BuildStatue(10)
	fmt.Printf("%s - %s - %s - %s - %d\n", h.materialWall, h.materialFloor, h.materialRoof, h.windowColor, h.statueNumber)
	return &House{
		materialWall:  h.materialWall,
		windowColor:   h.windowColor,
		materialRoof:  h.materialRoof,
		materialFloor: h.materialFloor,
		statueNumber:  h.statueNumber,
	}
}
