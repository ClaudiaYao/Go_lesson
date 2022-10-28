package main

import (
	"fmt"
	"reflect"
)

type customer struct {
	fName        string
	lName        string
	userID       int
	invoiceTotal float64
}

func inspect(dt interface{}) {
	switch dt.(type) {
	case customer:
		xt := reflect.TypeOf(dt)
		fmt.Println("type:", xt, "\tkind:", xt.Kind(), "\tNum of field:", xt.NumField())
	default:
		fmt.Println("type:", reflect.TypeOf(dt), "\tvalue:", reflect.ValueOf(dt))
	}

}

func main() {
	fmt.Println("")

	inspect("this is a string")
	inspect(12345)
	inspect(1.2345)
	inspect(true)

	inspect(customer{"John", "John", 123123123, 10000})

}
