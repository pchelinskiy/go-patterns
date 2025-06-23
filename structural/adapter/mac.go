package adapter

import "fmt"

type Mac struct {
}

func (m *Mac) InsertIntoLightningPort() {
	fmt.Println("lightning connector is plugged into mac machine")
}
