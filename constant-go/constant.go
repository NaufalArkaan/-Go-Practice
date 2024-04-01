//Constant go-lang

package main

import "fmt"

func main() {
	const firstName string = "Jhon"
	const lastName = "Dhoe"
	const value = 1000

	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(value)

	const (
		country string = "Indonesia"
		city           = "Jakarta"
	)

	fmt.Println(country)
	fmt.Println(city)
}
