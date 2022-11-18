package main

import (
	"fmt"
	"time"
)

func receiver(c chan int, n string) {
	fmt.Printf("%s - receiving\n", n)
	x := <-c
	fmt.Printf("%s - received %d\n", n, x)
}

func sender(v int, c chan int, n string) {
	fmt.Printf("%s - sending %d\n", n, v)
	c <- v
	fmt.Printf("%s - send %d\n", n, v)
}

func main() {
	c := make(chan int) // unbuffered channel

	go receiver(c, "X")
	go receiver(c, "Y")
	go sender(1, c, "A")
	// go sender(2, c, "B")

	time.Sleep(time.Second)
	fmt.Println("Done")
}

/*
	Does not make any sense:
	X - receiving
	B - sending 2
	B - send 2
	A - sending 1
	Y - receiving
	Y - received 1
	X - received 2
	A - send 1
	Done
*/
