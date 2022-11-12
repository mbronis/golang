/* 
	Write a program which prompts the user to enter a floating point number
	and prints the integer which is a truncated version of the floating point
	number that was entered. Truncation is the process of removing the digits
	to the right of the decimal place.
*/
package main

import "fmt"

func main() {
	var inputNumber float64
	var truncatedNumber int

	fmt.Printf("Input a float number: ")
	fmt.Scan(&inputNumber)
	truncatedNumber = int(inputNumber)
	fmt.Printf("Truncated: %d\n", truncatedNumber)
}
