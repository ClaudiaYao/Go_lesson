package main

import (
	"fmt"
)

func main() {

	customer1 := Customer{"Micheal", "Jordan", "MJ2020", "1234567", "MJ2020@gmail.com", 12345678, "18227 Capstan Greens Road Cornelius, NC 28031."}

	customer1.PrintAllUserInformation()
	fmt.Println(customer1.UserCredentials())
	fmt.Println(customer1.UserAddress())

}
