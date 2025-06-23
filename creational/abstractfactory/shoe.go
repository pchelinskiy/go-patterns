package abstractfactory

type IShoe interface {
	SetLogo(logo Logo)
	SetSize(size int)
	GetLogo() Logo
	GetSize() int
}

type Shoe struct {
	logo Logo
	size int
}

func (s *Shoe) SetLogo(logo Logo) {
	s.logo = logo
}

func (s *Shoe) GetLogo() Logo {
	return s.logo
}

func (s *Shoe) SetSize(size int) {
	s.size = size
}

func (s *Shoe) GetSize() int {
	return s.size
}
