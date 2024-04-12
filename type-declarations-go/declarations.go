// type-declarations go-lang

package main

import "fmt"

func main() {
	type NoKTP string
	type Married bool

	var noKtpEko NoKTP = "16278732948598"
	var marriedStatus Married = true

	fmt.Println(noKtpEko)
	fmt.Println(marriedStatus)
}
