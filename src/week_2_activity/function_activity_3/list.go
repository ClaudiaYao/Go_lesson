package main

import (
	"fmt"
)

func Show_game_info(game_slice []Game) {
	for _, item := range game_slice {
		fmt.Printf("game tile: %s, price is : $%.2f\n", item.title, item.price)
	}
}
