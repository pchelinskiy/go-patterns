package abstractfactory

import "errors"

var ErrWrongBrandName = errors.New("wrong brand name passed")

type IBrandFactory interface {
	MakeShoe() IShoe
	MakeShirt() IShirt
}

func NewBrandFactory(b BrandName) (IBrandFactory, error) {
	switch b {
	case AdidasBrand:
		return &Adidas{}, nil
	case NikeBrand:
		return &Nike{}, nil
	default:
		return nil, ErrWrongBrandName
	}
}
