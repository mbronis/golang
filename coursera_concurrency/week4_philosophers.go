/*
Implement the dining philosopher’s problem with the following
constraints/modifications.

There should be 5 philosophers sharing chopsticks, with one
chopstick between each adjacent pair of philosophers.

Each philosopher should eat only 3 times (not in an infinite
loop as we did in lecture)

The philosophers pick up the chopsticks in any order, not
lowest-numbered first (which we did in lecture).

In order to eat, a philosopher must get permission from a host
which executes in its own goroutine.

The host allows no more than 2 philosophers to eat concurrently.

Each philosopher is numbered, 1 through 5.

When a philosopher starts eating (after it has obtained necessary
locks) it prints “starting to eat <number>” on a line by itself,
where <number> is the number of the philosopher.

When a philosopher finishes eating (before it has released its locks)
it prints “finishing eating <number>” on a line by itself, where
<number> is the number of the philosopher.
*/
package main

import (
	"fmt"
	"sync"
)

type ChopS struct{ sync.Mutex }
type Philo struct {
	no              int
	leftCS, rightCS *ChopS
	eaten           int
}

func (p *Philo) eat(host *Host, wg *sync.WaitGroup) {
	for i := 0; i < 3; i++ {
		p.leftCS.Lock()
		p.rightCS.Lock()
		host.getPermission(p)
		fmt.Println("starting to eat", p.no)

		p.eaten++

		fmt.Println("finishing eating", p.no)
		host.releasePermission()
		p.rightCS.Unlock()
		p.leftCS.Unlock()
	}
	wg.Done()
}

type Host struct {
	dining chan *Philo
}

func (h *Host) getPermission(p *Philo) {
	h.dining <- p
}

func (h *Host) releasePermission() {
	<-h.dining
}

func main() {
	fmt.Println("Dining Philosophers")

	CSticks := make([]*ChopS, 5)
	for i := 0; i < 5; i++ {
		CSticks[i] = new(ChopS)
	}

	philos := make([]*Philo, 5)
	for i := 0; i < 5; i++ {
		leftCS, rightCS := CSticks[i], CSticks[(i+1)%5]
		if i > (i+1)%5 {
			leftCS, rightCS = rightCS, leftCS
		}
		philos[i] = &Philo{i + 1, leftCS, rightCS, 0}
	}

	var wg sync.WaitGroup
	host := Host{make(chan *Philo, 2)}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go philos[i].eat(&host, &wg)
	}
	wg.Wait()

	for _, p := range philos {
		fmt.Printf("%d ate %d\n", p.no, p.eaten)
	}
}
