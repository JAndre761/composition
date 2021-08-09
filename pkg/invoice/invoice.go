package invoice

import (
	"github.com/JAndre761/composition/pkg/customer"
	"github.com/JAndre761/composition/pkg/invoiceitem"
)

// Invoice is the struct of an invoice
type Invoice struct {
	country string
	city string
	total float64
	client customer.Customer
	items []invoiceitem.Item
}

// New returns a new invoice. Constructor.
func New(country, city string, client customer.Customer, items []invoiceitem.Item) Invoice {
	return Invoice{
		country: country,
		city:    city,
		client:  client,
		items:   items,
	}
}

// SetTotal setter of Invoice.total
func (i *Invoice) SetTotal() {
	for _, item := range i.items {
		i.total += item.Value()
	}
}
