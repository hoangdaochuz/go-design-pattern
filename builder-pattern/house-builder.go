package builderpattern

type HouseBuilder interface {
	BuildWall(string) HouseBuilder
	BuildRoof(string) HouseBuilder
	BuildWindow(string) HouseBuilder
	BuildFloor(string) HouseBuilder
	Build() *House
}
