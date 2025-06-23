package visitor

type Circle struct {
	Radius int
}

func (c *Circle) Accept(v Visitor) {
	v.visitForCircle(c)
}

func (c *Circle) GetType() string {
	return "Circle"
}
