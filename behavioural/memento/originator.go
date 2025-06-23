package memento

type Originator struct {
	state string
}

func (e *Originator) CreateMemento() *Memento {
	return &Memento{state: e.state}
}

func (e *Originator) RestoreMemento(m *Memento) {
	e.state = m.getSavedState()
}

func (e *Originator) SetState(state string) {
	e.state = state
}

func (e *Originator) GetState() string {
	return e.state
}
