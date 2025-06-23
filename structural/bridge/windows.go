package bridge

import "fmt"

type Windows struct {
	printer Printer
}

func (m *Windows) Print() {
	fmt.Println("print request for windows")
	m.printer.PrintFile()
}

func (m *Windows) SetPrinter(p Printer) {
	m.printer = p
}
