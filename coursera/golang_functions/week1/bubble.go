/*
	Write a Bubble Sort program in Go. The program should prompt the user to type
	in a sequence of up to 10 integers. The program	should print the integers out
	on one line, in sorted order, from least to	greatest. Use your favorite search
	tool to find a description of how the bubble sort algorithm works.

	As part of this program, you should write a	function called BubbleSort() which
	takes a slice of integers as an argument and returns nothing. The BubbleSort()
	function should modify the slice so that the elements are in sorted	order.

	A recurring operation in the bubble sort algorithm is the Swap operation which
	swaps the position of two adjacent elements in the slice. You should write a
	Swap() function which performs this operation. Your Swap() function should take
	two arguments, a slice of integers and an index value i which indicates a position
	in the slice. The Swap() function should return nothing, but it should swap	the
	contents of the slice in position i with the contents in position i+1.
*/
package main

import (
	"fmt"
	"strings"
	"os"
	"bufio"
	"strconv"
)

func main(){
	// read input
	var input string
	fmt.Printf("Input up to 10 numbers: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input = scanner.Text()

	// parse input to int slice
	input_splitted := strings.Split(input, " ")
	numbers := make([]int, 0, 10)
	for i, n := range input_splitted {
		if i > 10 {
			break
		}
		if value, err := strconv.Atoi(n); err != nil {
			continue
		} else {
			numbers = append(numbers, value)
		}
	}

	// sort & print results
	fmt.Println("numbers before sort: ", numbers)
	n_iter := BubbleSort(numbers)
	fmt.Printf("numbers after sort (%d iters): %v\n", n_iter, numbers)
}


func BubbleSort(s []int) int{
	var n int
	for n=0; n<len(s); n++ {
		updated := false
		for i:=0; i<len(s)-1-n; i++ {
			if s[i+1] < s[i] {
				Swap(s, i)
				updated = true
			}
		}
		if !updated {			
			return n
		}
	}
	return n
}

func Swap(s []int, i int){
	tmp := s[i]
	s[i] = s[i+1]
	s[i+1] = tmp
}