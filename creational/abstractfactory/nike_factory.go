package abstractfactory

type Nike struct{}

func (a *Nike) MakeShoe() IShoe {
	return &NikeShoe{
		Shoe: Shoe{
			logo: NikeLogo,
			size: 14,
		},
	}
}

func (a *Nike) MakeShirt() IShirt {
	return &NikeShirt{
		Shirt: Shirt{
			logo: NikeLogo,
			size: 14,
		},
	}
}
