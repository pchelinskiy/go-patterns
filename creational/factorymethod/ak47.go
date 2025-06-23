package factorymethod

type AK47 struct {
	gun
}

func newAK47() IGun {
	return &AK47{
		gun: gun{
			name:  AK47Name,
			power: 4,
		},
	}
}
