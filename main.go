package main

import (
	"fmt"
	"github.com/dobalu/luhn/car"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// random comment
func main() {
	carvar := car.NewCar("Audi", 4)
	// carvar := car.Car{Name: "Audi", Doors: 4}
	fmt.Println(carvar)
	car2 := car.NewCar("Datsun", 2)
	// car2 := &car.Car{Name: "Datsun", Doors: 2}
	fmt.Println(car2.String())
	funkybaby(true, "yeppers")
	http.HandleFunc("/", luhnHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
	// fmt.Println("Hello World")
	// luhn := isLuhn("4242424242424272")
	//
	// if luhn {
	// 	fmt.Println("Valid credit card number")
	// } else {
	// 	fmt.Println("Invalid card number")
	// }
}

func luhnHandler(writer http.ResponseWriter, request *http.Request) {
	// var luhn bool
	creditCardNumber := strings.Trim(request.URL.Path, "/")
	luhn := isLuhn(creditCardNumber)
	message := "The credit card number %s is %s."
	var status string

	if luhn {
		status = "valid"
	} else {
		status = "invalid"
	}

	fmt.Fprintf(writer, message, creditCardNumber, status)
	// fmt.Fprintf(w, "stronk\n")
	// fmt.Fprintf(w, "Hello, %q\n", html.EscapeString(r.URL.Path))
}

func sillystring() string {
	myString := "butts"
	return myString
}

func funkybaby(perchance bool, yep string) {
	fmt.Println(perchance, yep)
	return
}

// Declare function
// Name function
// Specify function arguments and argument type within parens
// Specify function return type
// Insert function code into curly brace
func isLuhn(cardNumber string) bool {
	var sum int
	parity := len(cardNumber) % 2

	// for (initialization; condition; step)
	// for i = 0; i < 10; i++
	for i := len(cardNumber) - 1; i >= 0; i-- {
		digit, _ := strconv.Atoi(string(cardNumber[i]))

		if i%2 == parity {
			digit = digit * 2
		}

		if digit > 9 {
			digit -= 9
		}

		sum += digit
	}

	return sum%10 == 0
}
