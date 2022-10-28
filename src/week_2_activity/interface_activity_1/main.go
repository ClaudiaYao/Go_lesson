package main

import "fmt"

func main() {
	fmt.Println("")

	equipment_slice := []equipment{Book{"Candle in the tomb", 20}, Book{"Barney and Friends", 19}, ComputerAccessories{"Razer BT earpiece", 159}, ComputerAccessories{"Razer keyboard", 110}, ComputerAccessories{"Logitech Mouse", 80}, Game{"Minecraft", 5}, Game{"World of warcraft", 19}, Game{"Elite Dangerous", 54}}

	Show_game_info(equipment_slice)

}
