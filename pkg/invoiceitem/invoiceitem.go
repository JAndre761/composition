package invoiceitem

type Item struct {
	id uint
	product string
	value float64
}

func New(id uint, product string, value float64) Item {
	return Item{id, product, value}
}

// Value getter if Item.Value
func (i Item) Value() float64 {
	return i.value
}