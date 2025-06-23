package abstractfactory

type IShirt interface {
	SetLogo(logo Logo)
	SetSize(size int)
	GetLogo() Logo
	GetSize() int
}

type Shirt struct {
	logo Logo
	size int
}

func (s *Shirt) SetLogo(logo Logo) {
	s.logo = logo
}

func (s *Shirt) GetLogo() Logo {
	return s.logo
}

func (s *Shirt) SetSize(size int) {
	s.size = size
}

func (s *Shirt) GetSize() int {
	return s.size
}
