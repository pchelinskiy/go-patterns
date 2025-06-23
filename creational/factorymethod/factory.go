package factorymethod

import "errors"

var (
	ErrWrongGunType = errors.New("wrong gun type passed")
)

type IGun interface {
	SetPower(power int)
	GetName() GunName
	GetPower() int
}

func NewGun(t GunName) (IGun, error) {
	switch t {
	case AK47Name:
		return newAK47(), nil
	case MusketName:
		return newMusket(), nil
	default:
		return nil, ErrWrongGunType
	}
}
