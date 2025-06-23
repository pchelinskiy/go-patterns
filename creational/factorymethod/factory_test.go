package factorymethod_test

import (
	"errors"
	"go-patterns/creational/factorymethod"
	"reflect"
	"testing"
)

type testCaseFactory struct {
	name         string
	gunName      factorymethod.GunName
	power        int
	expectedType reflect.Type
	err          error
}

func TestFactory(t *testing.T) {
	tc := []testCaseFactory{
		{
			name:         "ak47 creation testing",
			gunName:      factorymethod.AK47Name,
			power:        4,
			expectedType: reflect.TypeOf(&factorymethod.AK47{}),
		},
		{
			name:         "musket creation testing",
			gunName:      factorymethod.MusketName,
			power:        1,
			expectedType: reflect.TypeOf(&factorymethod.Musket{}),
		},
		{
			name:         "invalid gun type creation testing",
			gunName:      "",
			power:        0,
			expectedType: nil,
			err:          factorymethod.ErrWrongGunType,
		},
	}

	for _, tt := range tc {
		gun, err := factorymethod.NewGun(tt.gunName)
		if tt.err != nil && err != nil {
			if !errors.Is(tt.err, err) {
				t.Errorf("%s\nunexpected error\n\texpected:\n\t\t%#v\n\treceived:\n\t\t%#v", tt.name, tt.err, err)
			}
			continue
		}

		if tt.err != nil && err == nil {
			t.Errorf("%s\nexpected error, but received nil\n\texpected:\n\t\t%#v\n\t", tt.name, tt.err)
			continue
		}

		if tt.err == nil && err != nil {
			t.Errorf("%s\nunexpected error: %#v", tt.name, err)
			continue
		}

		if receivedType := reflect.TypeOf(gun); receivedType != tt.expectedType {
			t.Errorf("%s\nunexpected struct type\n\texpected:\n\t\t%v\n\treceived:\n\t\t%v", tt.name, tt.expectedType, receivedType)
			continue
		}

		if receivedType := gun.GetName(); receivedType != tt.gunName {
			t.Errorf("%s\nunexpected gun name\n\texpected:\n\t\t%v\n\treceived:\n\t\t%v", tt.name, tt.gunName, receivedType)
			continue
		}

		if receivedPower := gun.GetPower(); receivedPower != tt.power {
			t.Errorf("%s\nunexpected gun power\n\texpected:\n\t\t%v\n\treceived:\n\t\t%v", tt.name, tt.power, receivedPower)
			continue
		}
	}
}
