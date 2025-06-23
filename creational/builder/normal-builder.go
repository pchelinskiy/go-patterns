package builder

type NormalBuilder struct {
	house *House
}

func newNormalBuilder() IBuilder {
	return &NormalBuilder{
		house: &House{},
	}
}

func (b *NormalBuilder) setWindowType() IBuilder {
	b.house.windowType = "Wooden Window"
	return b
}

func (b *NormalBuilder) setDoorType() IBuilder {
	b.house.doorType = "Wooden Door"
	return b
}

func (b *NormalBuilder) setNumFloor() IBuilder {
	b.house.floor = 2
	return b
}

func (b *NormalBuilder) Build() *House {
	return b.house
}
