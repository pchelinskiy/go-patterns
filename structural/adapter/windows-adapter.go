package adapter

import "fmt"

type WindowsAdapter struct {
	Win *Windows
}

func (wa *WindowsAdapter) InsertIntoLightningPort() {
	fmt.Println("adapter converts Lightning signal to USB")
	wa.Win.InsertIntoUSBPort()
}
