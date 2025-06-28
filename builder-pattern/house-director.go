package builderpattern

type HouseDirector struct {
	builder HouseBuilder
}

func NewHouseDirector(builder HouseBuilder) *HouseDirector {
	return &HouseDirector{
		builder: builder,
	}
}

func (d *HouseDirector) BuildHouse(material, color string) *House {
	return d.builder.BuildWall(material).BuildRoof(material).BuildWindow(color).BuildFloor(material).Build()
}
