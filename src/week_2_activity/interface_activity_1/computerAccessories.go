package main

import "fmt"

type ComputerAccessories struct {
	title string
	price float64
}

func (ca ComputerAccessories) Get_price() {
	fmt.Printf("computer accessory: %s, price is $%.2f\n", ca.title, ca.price)
}
