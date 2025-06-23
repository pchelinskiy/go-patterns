package abstractfactory

type Adidas struct{}

func (a *Adidas) MakeShoe() IShoe {
	return &AdidasShoe{
		Shoe: Shoe{
			logo: AdidasLogo,
			size: 14,
		},
	}
}

func (a *Adidas) MakeShirt() IShirt {
	return &AdidasShirt{
		Shirt: Shirt{
			logo: AdidasLogo,
			size: 14,
		},
	}
}
