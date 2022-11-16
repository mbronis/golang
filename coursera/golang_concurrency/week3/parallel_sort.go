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

func getInput(n_chunks, max_len int) ([]int, error) {
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

func partition(s []int, nChunks int) [][]int {
	chunkSize := len(s) / nChunks
	var chunks [][]int
	for {
		if len(s) == 0 {
			break
		}
		if len(s) < chunkSize {
			for i, v := range s {
				chunks[i] = append(chunks[i], v)
			}
			return chunks
		}

		chunks = append(chunks, s[0:chunkSize])
		s = s[chunkSize:]
	}

	return chunks
}

func mergeTwo(a, b []int) []int {
	i, j := 0, 0
	var result []int

	for (i < len(a)) && (j < len(b)) {
		if a[i] < b[j] {
			result = append(result, a[i])
			i++
		} else {
			result = append(result, b[j])
			j++
		}
	}
	if i < len(a) {
		result = append(result, a[i:]...)
	}
	if j < len(b) {
		result = append(result, b[j:]...)
	}
	return result
}

func merge(s [][]int) []int {
	var result []int

	for _, ss := range s {
		result = mergeTwo(result, ss)
	}

	return result
}

func main() {
	nChunks := 4
	maxLen := 100

	// get input
	numbers, err := getInput(nChunks, maxLen)
	if err != nil {
		fmt.Printf("Invalid imput - %s\n", err)
		return
	}

	// split to chunks and sort each in goroutine
	chunks := partition(numbers, nChunks)
	ch := make(chan []int)
	for _, chunk := range chunks {
		go func(s []int, c chan []int) {
			fmt.Println("Chunk for sorting: ", s)
			sort.Ints(s)
			c <- s
		}(chunk, ch)
	}

	// gather sorted chunks
	var sortedChunks [][]int
	for range chunks {
		chunk := <-ch
		sortedChunks = append(sortedChunks, chunk)
	}

	// merge sorted chunks
	fmt.Println("Flattened sorted slices:", sortedChunks)
	result := merge(sortedChunks)
	fmt.Println("Result:", result)
}
