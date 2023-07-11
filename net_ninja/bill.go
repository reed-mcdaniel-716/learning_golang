package main

import (
	"fmt"
	"os"
)

// rather than classes, Go has types and receiver functions

// type definition of a bill
type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// make new bill
func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}

	return b
}

// format the bill
// explicitly associate that we will receive a bill object into the function
// copy of whatever is passed in
// * indicating we want the pointer to the bill rather than a copy
// this allows us to update the item itself as we'll have a reference to the original
func (b *bill) formatBill() string {
	fs := b.name + " bill breakdown: \n"
	total := 0.0

	for key, value := range b.items {
		// adding some spacing to the right of the key
		fs += fmt.Sprintf("%-25v ...$%.2f \n", key+":", value)
		total += value
	}

	total += b.tip

	fs += fmt.Sprintf("%-25v ...$%.2f \n", "tip:", b.tip)
	fs += fmt.Sprintf("%-25v ...$%.2f \n", "total:", total)

	return fs
}

// update tip
func (b *bill) updateTip(tip float64) {
	b.tip = tip
}

// add items to bill
func (b *bill) addItem(name string, price float64) {
	b.items[name] = price
}

// save bill
func (b *bill) save() {
	// saving formatted string as a byte slice
	data := []byte(b.formatBill())

	err := os.WriteFile("bills/"+b.name+".txt", data, 0644)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Bill was saved to %q \n", "bills/"+b.name+".txt")
}
