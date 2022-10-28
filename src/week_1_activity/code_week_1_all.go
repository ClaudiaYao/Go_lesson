package main

import (
	"bufio"
	"fmt"
	"math"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// Integer, Floating and Complex Types
// Activity #1: Simple Calculator Program in Go - Basic calculation
func simple_calculater() {

	var num_1, num_2 int
	var input string
	var operation string
	var result int

	// check if the first input is correct.
	for {
		fmt.Println("Input an integer value between 0-9: ")
		fmt.Scanln(&input)
		i, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("invalid input value.")
		} else {
			if i < 0 || i > 9 {
				fmt.Println("Input is not within the expected range.")
			} else {
				num_1 = i
				break
			}
		}
	}

	// check if the operation is valid
	for {
		fmt.Println("input your operation Add/Subtract/Multiplication/Divide: (a/s/m/d): ")

		_, err := fmt.Scan(&operation)
		if err != nil {
			fmt.Println("invalid input")
		} else {
			operation = strings.ToLower(operation)
			if operation == "a" || operation == "s" || operation == "m" || operation == "d" {
				break
			} else {
				fmt.Println("invalid input")
			}
		}

	}

	// check if the second number is correct
	for {
		fmt.Println("Input the second integer value between 0-9: ")
		fmt.Scan(&input)
		j, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("invalid input value.")
		} else {
			if j < 0 || j > 9 {
				fmt.Println("Input is not within the expected range.")
			} else {
				num_2 = j
				break
			}
		}
	}

	// doing the operation and print out result
	switch operation {
	case "a":
		result = num_1 + num_2
		fmt.Printf("%d + %d = %d\n", num_1, num_2, result)
	case "s":
		result = num_1 - num_2
		fmt.Printf("%d - %d = %d\n", num_1, num_2, result)
	case "m":
		result = num_1 * num_2
		fmt.Printf("%d * %d = %d\n", num_1, num_2, result)
	case "d":
		if num_2 == 0 {
			fmt.Println("Could not conduct the divide operation since 2nd number is zero.")

		} else {
			result = num_1 / num_2
			fmt.Printf("%d / %d = %d\n", num_1, num_2, result)
		}
	}

}

// convert temperature between Keivin, C and F. It is not the best solution, yet.
// complex data type: Activity #2: Temperature Converter - Kelvin to Degrees to Fahrenheit
func temp_converter() {
	var temp_unit int
	var cur_temp float64
	var temp_c, temp_k, temp_f float64
	fmt.Println("The unit of the temperature (1 for Kelvin, 2 for Celsius and 3 for Fahrenheit): ")

	fmt.Scanln(&temp_unit)

	fmt.Println("Input the current temperature:")
	fmt.Scanln(&cur_temp)

	switch temp_unit {
	case 1:
		temp_k = cur_temp
		temp_f = 9*temp_k/5 - 459.67
		temp_c = (temp_f - 32) * 5 / 9
		fmt.Printf("The current Keivin temperature %.2f is converted to Celsius %.2f and Fahrenheit %.2f.\n", temp_k, temp_c, temp_f)
	case 2:
		temp_c = cur_temp
		temp_f = temp_c*9/5 + 32
		temp_k = (temp_f + 459.67) * 5 / 9
		fmt.Printf("The current Celsius temperature %.2f is converted to Kelvin %.2f  and Fahrenheit %.2f.\n", temp_c, temp_k, temp_f)
	case 3:
		temp_f = cur_temp
		temp_c = (temp_f - 32) * 5 / 9
		temp_k = (temp_f + 459.67) * 5 / 9
		fmt.Printf("The current Fahrenheit temperature %.2f is converted to Kelvin %.2f and Celsius %.2f.\n", temp_f, temp_k, temp_c)
	default:
		fmt.Printf("input temperature unit is invalid.")
	}

}

// calculate total coin values.
// Activity #3: Currency calculator
func calc_coins() {
	// var input string
	var input string
	var total int
	values := []int{100, 50, 20, 10, 5}
	var input_s []string

	fmt.Printf("enter the number of 1-dollar coins, 50-cent coins, 20-cent coins, 10-cent coins and 5-cent coins (separated by space):")

	inputReader := bufio.NewReader(os.Stdin)
	input, err := inputReader.ReadString('\n')

	if err == nil {
		input_s = strings.Fields(input)
		for i, v := range input_s {
			num, err := strconv.Atoi(v)
			if err == nil {
				total += num * values[i]
			}
		}
	}

	fmt.Println("total cents: ", total)
	n := int(total / 200)
	m := float64(total-n*200) / 100.0
	fmt.Printf("give %d 2-dollar notes and then return charge %.2f dollars. \n", n, m)

}

// Activity #4: Triangle side calculator
func tiangle_calc() {
	var side_b, side_c, angle float64
	fmt.Println("Input the length of one side :")
	fmt.Scanln(&side_b)

	fmt.Println("Input the length of another side: ")
	fmt.Scanln(&side_c)

	fmt.Println("Input the angle between the two sides: ")
	fmt.Scanln(&angle)

	n := math.Pow(side_b, 2) + math.Pow(side_c, 2) - 2*side_b*side_c*math.Cos(math.Pi*angle/180)
	if n <= 0 {
		fmt.Println("the input values and angle could not form a triangle.")
	} else {
		fmt.Printf("the length of the third side is %.2f \n", math.Sqrt(n))
	}

}

// string concatenation
// Activity #1: Simple Concatenation
func simple_concat() {
	list1 := [4]string{"ans", "wer", "is", "of"}
	list2 := [4]string{"-", "+", "*", "/"}
	list3 := [4]string{"5", "10", "4", "0"}

	m, _ := strconv.Atoi(list3[0])
	n, _ := strconv.Atoi(list3[2])

	fmt.Printf("%s%s %s %s %s %s %s %d", list1[0], list1[1], list1[3], list3[0], list2[1], list3[2], list1[2], m+n)

}

// If statement practice
// Activity #5: Compound If...else statements
func check_leap_year() {
	var year int
	fmt.Println("Input a year:")
	fmt.Scanln(&year)

	if (year%4 == 0 && year%100 != 0) || (year%400 == 0) {
		fmt.Println("it is leap year!")
	} else {
		fmt.Println("it is not leap year.")
	}
}

// For statement practice
// Activity #2: Odd or even numbers
func check_odd_even() {
	var num int
	fmt.Println("please input an integer:")
	fmt.Scanln(&num)

	if num%2 == 0 {
		fmt.Println("even number.")
	} else {
		fmt.Println("odd number.")
	}
}

// Activity #3: For range
func check_range_numbers() {
	var start_num, end_num int
	fmt.Println("input the starting integer number:")
	fmt.Scanln(&start_num)

	fmt.Println("input the ending integer number: ")
	fmt.Scanln(&end_num)

	for i := start_num; i <= end_num; i++ {
		if i%2 == 0 {
			fmt.Println(i, "is an even number.")
		} else {
			fmt.Println(i, "is an odd number.")
		}
		if i > 9 {
			fmt.Println(i, "has more than one digit.")
		} else {
			fmt.Println(i, "has one digit.")
		}
	}
}

// Section: Basic Data Types
// Activity #1: Simple Types Printout Program in Go
func check_input() {
	text := "The following is the account information."
	firstName, lastName := "Luke", "Skywalkter"
	age, weight, height := 20, 73.0, 1.72
	remaining_credit := 123.45
	accountName, accountPassword := "admin", "password"
	subscribedUser := true

	fmt.Printf("text: %s, type: %T.\n", text, text)
	fmt.Printf("first name: %s, type: %T.\n", firstName, firstName)
	fmt.Printf("last name: %s, type: %T.\n", lastName, lastName)
	fmt.Printf("weight: %.2f, type: %T.\n", weight, weight)
	fmt.Printf("height: %.2f, type: %T.\n", height, height)
	fmt.Printf("age: %d, type: %T.\n", age, age)
	fmt.Printf("remainning credit: %.2f, type: %T.\n", remaining_credit, remaining_credit)
	fmt.Printf("account name: %s, type: %T.\n", accountName, accountName)
	fmt.Printf("account password: %s, type: %T. \n", accountPassword, accountPassword)
	fmt.Printf("sbuscribed user: %t, type: %T.\n", subscribedUser, subscribedUser)
}

// Boolean value: Activity #1: Simple True-False Program
func check_user_input() {
	fmt.Println("Enter a value between 1 - 100:")

	var expNum = rand.Float64() * 100 //create a random float64 number
	var num float64

	// Taking input from user
	fmt.Scanln(&num)
	if num > expNum {
		fmt.Println("Too high, try again next time!")
	} else if num < expNum {
		fmt.Println("Too low, try again next time!")
	} else {
		fmt.Println("Well Done! Your guess is correct.")
	}
}

// For statement practice
// Activity #6: Number guessing game
func guess_number() {
	reader := bufio.NewReader(os.Stdin)

	rand.Seed(time.Now().UnixNano())
	x := rand.Intn(100) // an int between 0 to 100

	for i := 0; i < 7; i++ {
		fmt.Print("Enter a number between 0-100: ")
		text, _ := reader.ReadString('\n')
		//pay attention that the text contains the \n, so we need to remove
		//those white spaces.
		fmt.Printf(text)

		num, err := strconv.Atoi(strings.TrimSpace(text))
		if err != nil {
			fmt.Println("incorrect input.")
		} else {
			if num > x {
				fmt.Println("guess is too big.")
			} else if num < x {
				fmt.Println("guess is too small.")
			} else {
				fmt.Println("bingo!")
				return
			}
		}
	}
	fmt.Println("exceed the 7 times. try next time!")

}

// For loop practice
// Activity #7: Verify Vehicle Plate Number Authenticity
func check_vehicle_num() {

	fmt.Println("Enter a vehicle plate number: ")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	// use regular expression to find all the matching parts
	// matches[0] - all matches, matches[1] - group 1 match result, matches[2] - group 2 match result
	// matches[3] - group 3 match result
	r, _ := regexp.Compile("([a-zA-Z]+)([0-9]+)([A-Z])")
	matches := r.FindStringSubmatch(input)
	fmt.Println("match results: ", matches)

	var nums [6]int
	var chars string

	m := len(matches[1])
	chars = strings.ToUpper(matches[1])
	if m == 3 {
		nums[0] = int(chars[1]) - int('A') + 1
		nums[1] = int(chars[2]) - int('A') + 1
	} else if m == 1 {
		nums[0] = 0
		nums[1] = int(chars[0]) - int('A') + 1
	} else {
		nums[0] = int(chars[0]) - int('A') + 1
		nums[1] = int(chars[1]) - int('A') + 1
	}

	// fill in the nums slice for the numbers part. Pay attention that user might input less than 4 numbers.
	ns := matches[2]
	// fmt.Println(ns)
	k := len(ns)
	for i := 0; i < k; i++ {
		nums[5-i] = int(ns[k-1-i]) - int('1') + 1
	}

	var total int
	factors := []int{9, 4, 5, 4, 3, 2}

	for i, v := range factors {
		total += nums[i] * v
	}

	rem := total % 19
	// fmt.Println(rem)
	checkChars := "AZYXUTSRPMLKJHGEDCB"

	// calcuated the expected last letter and compare with matches[3]
	fmt.Println(matches[3])
	if matches[3] == string(checkChars[rem]) {
		fmt.Println("genuine!")
	} else {
		fmt.Println("Fake!")
	}

}

// practice switch statement and fallthrough
func fall_through() {
	num := 2
	switch num {
	case 1:
		fmt.Println("1")
		fallthrough
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3")
		fallthrough

	default:
		fmt.Println("final line.")
	}

}

// array
// Activity #1: walk through each element in the array
// Activity #2: Assess a particular element in an array and update the elem
// Activity #3: Add elements to an array
// Activity #4: Assess a particular segment in an array
func check_array_element() {
	os := []string{"Windows XP", "Linux 1.0", "Raspbian Teensy", "MacOS", "IOS", "Google Android"}
	// activity #1
	for _, v := range os {
		fmt.Println("element", v, "len:", len(v))
	}

	// activity #2
	os[0] = "Windows 10"
	os[1] = "Linux 16.04"
	os[2] = "Raspbian Buster"
	fmt.Println("updated os types:", os)

	// activity #3
	os_slice := make([]string, len(os))
	copy(os_slice, os)
	os_slice = append(os_slice, "Ubuntu")
	os_slice = append(os_slice, "MS-Dos")
	os_slice = append(os_slice, "Solaris")

	fmt.Println("extended array:", os_slice)

	// Activity #4: Assess a particular segment in an array
	fmt.Println("the first 3 elements:", os_slice[:3])
	fmt.Println("the second 3 elements:", os_slice[3:6])
	fmt.Println("the third 3 elements:", os_slice[6:9])

}

// Activity #5: Temperature Sensor
func temp_sensor() {
	temp_slice := []float64{20.1, 24, 27.3, 30.1, 26.4, 22.2, 20.1, 24, 27.3, 30.1,
		26.4, 20.1, 24, 27.3, 30.1, 26.4, 20.1, 24, 27.3, 30.1,
		26.4, 20.1, 24, 27.3}

	sum := 0.0

	for _, v := range temp_slice {
		sum += v

	}
	fmt.Printf("the average temperature is: %.2f \n", sum/float64(len(temp_slice)))

}

// slice practice: activity #1, #2, #3, #4
func slice_operation() {
	spending := make([]float64, 4, 5)
	spending[0] = 9.50
	spending[1] = 8.00
	spending[2] = 10.20
	spending[3] = 7.40

	fmt.Println("the length of slice: ", len(spending), "capacity:", cap(spending))

	fmt.Println("spending of weeek 3:", spending[2])
	spending[2] = 9.80
	fmt.Println("updated spending of week 3:", spending[2])

	spending = append(spending, 8.40)
	spending = append(spending, 9.20)
	spending = append(spending, 7.20)
	fmt.Println("updated length of slice: ", len(spending), "capacity:", cap(spending))

	fmt.Println("subslice from position 3:", spending[2:])

}

// slice operation
// Activity #5: Temperature Sensor with nested loop and slices
func temp_sensor_multiroom() {
	room_temp := [][]float64{
		{20, 21, 23, 25, 22},
		{27, 23, 25, 20, 30, 24},
		{22, 23, 24, 22},
	}

	var t, num = 0.0, 0

	for _, v := range room_temp {
		for _, m := range v {
			t += m
			num += 1
		}

	}
	fmt.Println("the average temperature of multiple rooms are:", t/float64(num))

}

// struct practice
// Activity #1 Declare a struct type
// Activity #2 Declare 2 users using struct
func practice_struct() {
	type User struct {
		fname            string
		lname            string
		age              int
		subscriber       bool
		homeAddress      string
		phone            int
		creditAvailable  float32
		currentCartCost  float32
		currentOrderCost float32
	}

	student_1 := User{"Annakin", "Skywalker", 45, true, "Death Star", 1234567,
		10000.00, 0.00, 0.00}

	student_2 := User{"Han", "Solo", 50, false, "Tatooine", 4321765, 20000.00,
		0.00, 0.00}

	fmt.Println(student_1, student_2)
}

// map section
// Activity #1 Define a map
// Activity #2 Currency conversion application
func map_operation() {

	// Activity #1 Declare a currency conversion map
	type currency_info struct {
		full_name     string
		exchange_rate float64
	}

	// fill in the data to the map
	current_exchange := map[string]currency_info{"USD": currency_info{"US dollar", 1.1318},
		"JPY": currency_info{"Japanese yen", 121.05},
		"GBP": currency_info{"Pound sterling", 0.90630},
		"CNY": currency_info{"Chinese yuan renminbi", 7.9944},
		"SGD": currency_info{"Singapore dollar", 1.5743},
		"MYR": currency_info{"Malaysian ringgit", 4.8390}}

	// Activity #1 Declare a currency conversion map
	key_string := ""
	for k, v := range current_exchange {
		key_string += "/" + k
		fmt.Println(k, v.full_name, v.exchange_rate)
	}

	var cur_from, cur_to string
	var amount float64
	// Activity #2 Currency conversion application
	fmt.Println("currency to convert from:", key_string)
	fmt.Scanln(&cur_from)

	from_struct, exist := current_exchange[strings.ToUpper(cur_from)]
	if exist == false {
		fmt.Println("currency converted from does not exist.")
		return
	}

	fmt.Println("input the amount of currentcy:")
	fmt.Scanln(&amount)

	if amount < 0 {
		fmt.Println("Could not convert negative amount of currency.")
		return
	}

	fmt.Println("currency to convert to:", key_string)
	fmt.Scanln(&cur_to)

	if cur_from == cur_to {
		fmt.Println("two currency should be different, otherwise, you do no need to convert.")
		return
	}

	to_struct, exist := current_exchange[strings.ToUpper(cur_to)]
	if exist == false {
		fmt.Println("currency converted to does not exist.")
		return
	}

	converted_amount := amount * to_struct.exchange_rate / from_struct.exchange_rate
	fmt.Println("converted from", cur_from, amount, "to", cur_to, ":", math.Round(100*converted_amount)/100)

}

func main() {

	// Basic Data types: Activity #1: Simple Types Printout Program in Go
	check_input()

	// Boolean Value. Activity #1: Simple True-False Program
	check_user_input()

	//complex data type: Activity #1: Simple Calculator Program in Go - Basic calculation
	//Activity #1: simple calculator
	simple_calculater()

	// Activity #2: Temperature Converter - Kelvin to Degrees to Fahrenheit
	temp_converter()

	// Activity #3: Currency calculator
	calc_coins()

	// Activity #4: Triangle side calculator
	tiangle_calc()

	// string concatenation
	// Activity #1: Simple Concatenation
	simple_concat()

	// If statement practice
	// Activity #5: Compound If...else statements
	check_leap_year()

	// For statement practice
	// Activity #2: Odd or even numbers
	check_odd_even()

	//Activity #3: check range numbers
	check_range_numbers()

	// Activity #6: Number guessing game
	guess_number()

	// Activity #7: Verify Vehicle Plate Number Authenticity
	check_vehicle_num()

	//switch statement practice
	fall_through()

	// array practice
	// Activity #1, #2, #3, #4  Declare array variable with a specific size
	// Activity #2: Assess a particular element in an array and update the element
	// add new elements
	check_array_element()

	// Activity #5: Temperature Sensor
	temp_sensor()

	// slice operation, all #1, #2, #3, #4
	slice_operation()

	// slice operation
	// Activity #5: Temperature Sensor with nested loop and slices
	temp_sensor_multiroom()

	//struct section
	// Activity #1 Declare a struct for user
	// Activity #2 Declare 2 users using struct
	practice_struct()

	// Map section
	// Activity #1 Declare a currency conversion map
	// Activity #2 Currency conversion application
	map_operation()

}
