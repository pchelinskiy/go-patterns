package factorymethod

type gun struct {
	name  GunName
	power int
}

func (g *gun) GetName() GunName {
	return g.name
}

func (g *gun) SetPower(p int) {
	g.power = p
}

func (g *gun) GetPower() int {
	return g.power
}
