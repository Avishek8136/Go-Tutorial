package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

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

// update tip
func (b *Bill) updateTip(tip float64) {
	b.tip = tip
}

// add item
func (b *Bill) addItem(name string, price float64) {
	b.Items[name] = price
}

// save bill
func (b *Bill) saveBill() {
	if err := os.MkdirAll("bills", 0755); err != nil {
		panic(err)
	}
	data, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		panic(err)
	}
	filename := "bills/" + strconv.Itoa(b.ID) + ".json"
	if err := os.WriteFile(filename, data, 0644); err != nil {
		panic(err)
	}
	fmt.Println("Bill was saved to file successfully")
}

func ReadBill(id int) Bill {
	filename := "bills/" + strconv.Itoa(id) + ".json"
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("No saved bill found, creating a new one")
		return NewBill(id)
	}

	var b Bill
	if err := json.Unmarshal(data, &b); err != nil {
		fmt.Println("Error reading saved bill, creating new one:", err)
		return NewBill(id)
	}

	fmt.Println(b.format())
	return b
}
