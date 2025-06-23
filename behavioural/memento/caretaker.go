package memento

type Caretaker struct {
	mementoArray []*Memento
}

func (c *Caretaker) AddMemento(m *Memento) {
	c.mementoArray = append(c.mementoArray, m)
}

func (c *Caretaker) GetMemento(index int) *Memento {
	return c.mementoArray[index]
}
