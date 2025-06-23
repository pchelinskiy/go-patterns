package factorymethod

type Musket struct {
	gun
}

func newMusket() IGun {
	return &Musket{
		gun: gun{
			name:  MusketName,
			power: 1,
		},
	}
}
