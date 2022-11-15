package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func dist_sort(ints_to_sort []int, c chan []int) {
	fmt.Printf("Goroutine before sorting: %v\n", ints_to_sort)
	sort.Ints(ints_to_sort)
	fmt.Printf("Goroutine after sorting: %v\n", ints_to_sort)
	c <- ints_to_sort
}

func main() {
	const chunks_num int = 4
	var fields []string
	var input_ints []int

	for {
		fmt.Print("Please enter a number of integers (more than 4 and space separated), to be sorted: ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		sc_err := scanner.Err()
		if sc_err == nil {
			//pi.Name = scanner.Text()
			fields = strings.Fields(scanner.Text())

			if len(fields) < chunks_num {
				fmt.Println("Wrong input!")
				continue
			}
			input_ints = make([]int, len(fields))
			for k, f := range fields {
				tmp_int, err := strconv.Atoi(f)
				if err != nil {
					fmt.Println("Wrong input!")
					break
				}
				input_ints[k] = tmp_int
			}
			if len(input_ints) == len(fields) {
				fmt.Printf("Ints given as input: %v\n", input_ints)
				break
			}
		}
	}

	var chunks_len int = len(input_ints) / chunks_num
	c := make(chan []int, chunks_num)

	for i := 0; i < chunks_num; i++ {
		var chunk_end int
		if (len(input_ints)-i*chunks_len)/chunks_len == 1 {
			chunk_end = len(input_ints)
		} else {
			chunk_end = i*chunks_len + chunks_len
		}
		chunk := input_ints[i*chunks_len : chunk_end]
		go dist_sort(chunk, c)
	}

	dist_sorted_ints := make([]int, 0, len(input_ints))
	for j := 0; j < chunks_num; j++ {
		sorted_chunk := <-c
		dist_sorted_ints = append(dist_sorted_ints, sorted_chunk...)
	}

	fmt.Printf("Ints after distributed/concurrent sorting and merge: %v\n", dist_sorted_ints)
	sort.Ints(dist_sorted_ints)
	fmt.Printf("Ints after final sorting by main, applied on the merged sorted slice: %v\n", dist_sorted_ints)
}
