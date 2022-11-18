/*
	Write two goroutines which have a race condition when executed concurrently.
	Explain what the race condition is and how it can occur.
*/
package main

import (
	"fmt"
	"time"
	"sync"
)

var wg sync.WaitGroup

func add(n *int){
	defer wg.Done()
	time.Sleep(time.Millisecond)
	*n++
	fmt.Println("added")
}

func mul(n *int){
	defer wg.Done()
	time.Sleep(time.Millisecond)
	*n *= 2
	fmt.Println("multiplied")
}

func main() {
	number := 0

	for i:=0; i<10; i++ {
		/*
			In each pass we add and multiply the same number
			with a separate goroutine.
			Race condition occurs as the order of groutines
			execution in each pass is not deterministic.
		*/
		wg.Add(2)
		go add(&number)
		go mul(&number)
		wg.Wait()
	}
	fmt.Println("number:", number)
}
