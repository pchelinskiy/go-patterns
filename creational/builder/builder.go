package builder

import "errors"

var ErrWrongBuilderType = errors.New("wrong builder type passed")

type IBuilder interface {
	setWindowType() IBuilder
	setDoorType() IBuilder
	setNumFloor() IBuilder
	Build() *House
}

func NewBuilder(t BuilderType) (IBuilder, error) {
	switch t {
	case NormalBuilderType:
		return newNormalBuilder(), nil
	case IglooBuilderType:
		return newIglooBuilder(), nil
	default:
		return nil, ErrWrongBuilderType
	}
}
