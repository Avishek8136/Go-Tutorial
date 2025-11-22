package main

type Bill struct {
	ID    int
	Items map[string]float64
	tip   float64
}

// bill generation
func NewBill(id int) Bill {
	return Bill{
		ID:    id,
		Items: map[string]float64{},
		tip:   0,
	}
}
