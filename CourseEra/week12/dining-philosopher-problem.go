// Implement the dining philosopher’s problem with the following constraints/modifications.
// 1. There should be 5 philosophers sharing chopsticks,
//		with one chopstick between each adjacent pair of philosophers.
// 2. Each philosopher should eat only 3 times (not in an infinite loop as we did in lecture)
// 3. The philosophers pick up the chopsticks in any order, not lowest-numbered first (which we did in lecture).
// 4. In order to eat, a philosopher must get permission from a host which executes in its own goroutine.
// 5. The host allows no more than 2 philosophers to eat concurrently.
// 6. Each philosopher is numbered, 1 through 5.
// 7. When a philosopher starts eating (after it has obtained necessary locks) it prints “starting to eat <number>” on a line by itself, where <number> is the number of the philosopher.
// 8. When a philosopher finishes eating (before it has released its locks) it prints “finishing eating <number>” on a line by itself, where <number> is the number of the philosopher.
// Submission: Upload your source code for the program.

package main

import (
	"fmt"
	"sync"
)

const NUM_DIET = 3

var wg sync.WaitGroup

var this_chan = make(chan int, 2)

type ChopStick struct{ sync.Mutex }

type Philosopher struct {
	identifier      int
	leftCS, rightCS *ChopStick
}

func pick_chopstick(c *ChopStick, this_wg *sync.WaitGroup) {
	c.Lock()
	this_wg.Done()
}

func drop_chopstick(c *ChopStick, this_wg *sync.WaitGroup) {
	c.Unlock()
	this_wg.Done()
}

func (p Philosopher) eat() {
	var this_wg sync.WaitGroup
	for i := 0; i < NUM_DIET; i++ {
		this_chan <- p.identifier
		this_wg.Add(2)
		pick_chopstick(p.leftCS, &this_wg)
		pick_chopstick(p.rightCS, &this_wg)
		this_wg.Wait()
		fmt.Printf("starting to eat %d\n", p.identifier)

		fmt.Printf("finishing eating %d\n", p.identifier)
		this_wg.Add(2)
		drop_chopstick(p.leftCS, &this_wg)
		drop_chopstick(p.rightCS, &this_wg)
		this_wg.Wait()
		<-this_chan
	}
	wg.Done()
}

func main() {
	num_philosophers := 5
	num_of_chopsticks := 5

	all_chopsticks := make([]*ChopStick, num_of_chopsticks)
	all_philosophers := make([]*Philosopher, num_philosophers)

	for i := 0; i < num_of_chopsticks; i++ {
		all_chopsticks[i] = new(ChopStick)
	}
	for i := 0; i < num_philosophers; i++ {
		all_philosophers[i] = &Philosopher{i, all_chopsticks[i], all_chopsticks[(i+1)%num_of_chopsticks]}
	}

	// Start eating
	wg.Add(num_philosophers)
	for i := 0; i < num_philosophers; i++ {
		go all_philosophers[i].eat()
	}

	wg.Wait()
}
