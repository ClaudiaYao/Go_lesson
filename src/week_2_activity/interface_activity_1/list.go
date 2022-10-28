package main

type equipment interface {
	Get_price()
}

func Show_game_info(e []equipment) {
	for _, item := range e {
		item.Get_price()
	}
}
