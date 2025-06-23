package bridge

import "fmt"

type Epson struct {
}

func (p *Epson) PrintFile() {
	fmt.Println("printing by a EPSON printer")
}
