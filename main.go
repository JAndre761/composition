package main

import (
	"fmt"
	"github.com/JAndre761/composition/pkg/customer"
	"github.com/JAndre761/composition/pkg/invoice"
	"github.com/JAndre761/composition/pkg/invoiceitem"
)

func main() {
	i := invoice.New(
		"Perusalen",
		"Lima",
		customer.New("Tones", "Av. This is a test", "999888777"),
		[]invoiceitem.Item{
			invoiceitem.New(1,"Course 1", 10.2),
			invoiceitem.New(2,"Course 2", 10.1),
			invoiceitem.New(2,"Course 3", 10.1),
		},
	)
	i.SetTotal()
	fmt.Printf("%+v", i)
}
