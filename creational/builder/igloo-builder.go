package builder

type IglooBuilder struct {
	house *House
}

func newIglooBuilder() IBuilder {
	return &IglooBuilder{
		house: &House{},
	}
}

func (b *IglooBuilder) setWindowType() IBuilder {
	b.house.windowType = "Wooden Window"
	return b
}

func (b *IglooBuilder) setDoorType() IBuilder {
	b.house.doorType = "Wooden Door"
	return b
}

func (b *IglooBuilder) setNumFloor() IBuilder {
	b.house.floor = 2
	return b
}

func (b *IglooBuilder) Build() *House {
	return b.house
}
