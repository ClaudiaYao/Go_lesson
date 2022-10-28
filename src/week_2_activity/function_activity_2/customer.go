package main

import "fmt"

type Customer struct {
	firstName string
	lastName  string
	userName  string
	password  string
	email     string
	phone     int
	address   string
}

func (cust *Customer) UserCredentials() (string, string) {
	return cust.userName, cust.password
}

func (cust *Customer) UserAddress() string {
	return cust.address

}

func (cust *Customer) PrintAllUserInformation() {
	fmt.Println("User Name: ", cust.firstName, cust.lastName)
	fmt.Println("User Credentials: ", cust.userName, cust.password)
	fmt.Println("Contact Info:", cust.email, cust.phone)
	fmt.Println("User Address: ", cust.UserAddress())

}
