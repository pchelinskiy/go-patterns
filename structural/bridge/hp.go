package bridge

import "fmt"

type HP struct {
}

func (p *HP) PrintFile() {
	fmt.Println("printing by a HP printer")
}
