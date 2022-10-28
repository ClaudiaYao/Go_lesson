package main

import "fmt"

type Game struct {
	title string
	price float64
}

func (game Game) Get_price() {
	fmt.Printf("price %.2f of the game %s: ", game.price, game.title)
}
