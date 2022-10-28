package main

import "fmt"

type Book struct {
	title string
	price float64
}

func (bk Book) Get_price() {
	fmt.Printf("book title: %s, price is $%.2f\n", bk.title, bk.price)
}
