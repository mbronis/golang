/*
	Write a program which prompts the user to first enter a name,
	and then enter an address. Your program should create a map and
	add the name and address to the map using the keys “name” and “address”,
	respectively.
	Your program should use Marshal() to create a JSON object from the map,
	and then your program should print the JSON object.
*/
package main

import (
	"fmt"
	"encoding/json"
)

func main(){
	var name, address string

	fmt.Printf("Enter user name: ")
	fmt.Scan(&name)
	fmt.Printf("Enter user address: ")
	fmt.Scan(&address)

	user := map[string]string {
		"name": name,
		"address": address,
	}
	barr, _ := json.MarshalIndent(user, "", "    ")
	jsonObj := string(barr)

	fmt.Println("JSON object:")
	fmt.Println(jsonObj)
}

