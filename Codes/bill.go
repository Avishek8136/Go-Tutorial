package main

import "fmt"

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

// formatting bill
func (b Bill) format() string {
	fs := "Bill Breakdown\n"
	var total float64 = 0
	fmt.Println("Bill", b.ID)
	for k, v := range b.Items {
		fs += fmt.Sprintf("%-25v ...$%v\n", k+":", v)
		total += v
	}
	fs += fmt.Sprintf("%25v ...$%v\n", "Tip:", b.tip)
	fs += fmt.Sprintf("%25v ...$%v\n", "Total:", total+b.tip)
	return fs
}

//update tip
func (b *Bill) updateTip(tip float64) {
	b.tip = tip
}

//add item
func (b *Bill) addItem(name string, price float64) {
	b.Items[name] = price
}
