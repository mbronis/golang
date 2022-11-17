/*
	Write a program which reads information from a file and represents
	it in a slice of structs. Assume that there is a text file which
	contains a series of names. Each line of the text file has a first
	name and a last name, in that order, separated by a single space
	on the line. 

	Your program will define a name struct which has two fields, fname
	for the first name, and lname for the last name. Each field will
	be a string of size 20 (characters).

	Your program should prompt the user for the name of the text file.
	Your program will successively read each line of the text file and
	create a struct which contains the first and last names found in
	the file. Each struct created will be added to a slice, and after
	all lines have been read from the file, your program will have
	a slice containing one struct for each line in the file. After
	reading all lines from the file, your program should iterate through
	your slice of structs and print the first and last names found in
	each struct.
*/
package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

type Name struct {
	fname string
	lname string
}

func main(){
	// get file name
	var fileName string
	fmt.Printf("Input file name: ")
	fmt.Scan(&fileName)

	// read file and split lines
	f, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Failed to read file..")
		fmt.Println(err)
	}
	defer f.Close()

	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)

	// parse lines
	var names []Name
	for fileScanner.Scan() {
		line := fileScanner.Text()
		tokens := strings.Split(line, " ")
		name := Name{fname: tokens[0], lname: tokens[1]}
		names = append(names, name)
    }
	f.Close()

	for _, n := range names {
		fmt.Println(n)
	}
}