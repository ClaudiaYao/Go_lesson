package main

import "fmt"

type Game struct {
	title string
	price float64
}

func (game Game) Get_price() {
	fmt.Printf("game title: %s, price is $%.2f\n", game.title, game.price)
}
