/*
Write a program to sort an array of integers. The program should partition
the array into 4 parts, each of which is sorted by a different goroutine.
Each partition should be of approximately equal size. Then the main
goroutine should merge the 4 sorted subarrays into one large sorted array.

The program should prompt the user to input a series of integers. Each
goroutine which sorts ¼ of the array should print the subarray that it
will sort. When sorting is complete, the main goroutine should print
the entire sorted list.
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type NotEnoughNumbersError struct{ expected, count int }

func (e *NotEnoughNumbersError) Error() string {
	return "there should be at least " + strconv.Itoa(e.expected) + " numbers, got " + strconv.Itoa(e.count)
}

func get_input(n_chunks, max_len int) ([]int, error) {
	var input string
	var err error

	fmt.Printf("Input %d or more numbers: ", n_chunks)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input = scanner.Text()

	// parse input to int slice
	input_splitted := strings.Split(input, " ")
	numbers := make([]int, 0, max_len)
	for i, n := range input_splitted {
		if i > max_len {
			break
		}
		if value, err := strconv.Atoi(n); err != nil {
			continue
		} else {
			numbers = append(numbers, value)
		}
	}

	l := len(numbers)
	if l >= n_chunks {
		err = nil
	} else {
		err = &NotEnoughNumbersError{n_chunks, l}
	}

	return numbers, err
}

func partition(s []int, n_chunks int) [][]int {
	chunk_size := len(s) / n_chunks
	var chunks [][]int
	for {
		if len(s) == 0 {
			break
		}
		if len(s) < chunk_size {
			chunk_size = len(s)
		}

		chunks = append(chunks, s[0:chunk_size])
		s = s[chunk_size:]
	}

	return chunks
}

func main() {
	n_chunks := 4
	max_len := 100

	// get input
	numbers, err := get_input(n_chunks, max_len)
	if err != nil {
		fmt.Printf("Invalid imput - %s\n", err)
		return
	}
	chunks := partition(numbers, n_chunks)

	// spawn sorting goroutines
	ch := make(chan []int)
	for _, chunk := range chunks {
		go func(s []int, c chan []int) {
			fmt.Println("Chunk for sorting: ", s)
			sort.Ints(s)
			c <- s
		}(chunk, ch)
	}

	// gather and flatten sorted chunks
	var sorted_chunks []int
	for range chunks {
		chunk := <-ch
		sorted_chunks = append(sorted_chunks, chunk...)
	}

	fmt.Println("Flattened sorted slices:", sorted_chunks)
	sort.Ints(sorted_chunks)
	fmt.Println("Sorted slice:", sorted_chunks)
}
