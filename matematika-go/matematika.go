// operasi matematika pada go-lang

package main

import "fmt"

func main() {
	var result = 10 + 10
	fmt.Println(result)

	var a = 10
	var b = 10
	var c = a * b
	fmt.Println(c)

	var i = 10
	i += 10 //sama dengan i+i = 20
	fmt.Println(i)

	//unary operator
	i++ //i = i + 1
	fmt.Println(i)

	var j = 1
	j++
	j++
	fmt.Println()
}
