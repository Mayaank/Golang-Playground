/*
Implement the dining philosopherâ€™s problem with the following constraints/modifications.

    There should be 5 philosophers sharing chopsticks, with one chopstick between each adjacent pair of philosophers.

    Each philosopher should eat only 3 times (not in an infinite loop as we did in lecture)

    The philosophers pick up the chopsticks in any order, not lowest-numbered first (which we did in lecture).

    In order to eat, a philosopher must get permission from a host which executes in its own goroutine.

    The host allows no more than 2 philosophers to eat concurrently.

    Each philosopher is numbered, 1 through 5.

    When a philosopher starts eating (after it has obtained necessary locks) it prints â€œstarting to eat <number>â€ on a line by itself, where <number> is the number of the philosopher.

    When a philosopher finishes eating (before it has released its locks) it prints â€œfinishing eating <number>â€ on a line by itself, where <number> is the number of the philosopher.


olaf@oslo /cryptdata6/var/log/tmp/shared/gopath/src/github.com/scottstensland/play/coursera/golangConcurrency $ cp concurrPhilosophers_works_1121_1350.go concurrPhilosophers_works_1121_1350__20221226_01.go


*/

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type philosWatcherHost struct { //

	sync.Mutex
}

type ChopStick struct {
	sync.Mutex
}

type Philo struct {
	name            int
	leftCS, rightCS *ChopStick
	countEat        int
}

func (p Philo) doEat(philoIsFull map[int]int, maxEatCycles int, wg *sync.WaitGroup) {

	defer wg.Done()

	for p.countEat < maxEatCycles {

		p.leftCS.Lock()
		p.rightCS.Lock()

		fmt.Printf("starting to eat %d\n", p.name)
		fmt.Printf("finished eating %d\n", p.name)

		p.rightCS.Unlock()
		p.leftCS.Unlock()

		p.countEat++
	}

	if p.countEat == maxEatCycles {

		philoIsFull[p.name] = p.name
	}
}

func getRandIntFromRange(lowEnd, highEnd int) int {

	return (rand.Intn(highEnd-lowEnd) + lowEnd)
}

func getPairInts(countPhilos int) (int, int) {

	leftInt := getRandIntFromRange(0, countPhilos-1)

	rightInt := getRandIntFromRange(0, countPhilos-1)

	for leftInt == rightInt {

		rightInt = getRandIntFromRange(0, countPhilos-1)
	}

	return leftInt, rightInt
}

func getPhilo(countPhilos int, philoIsFull map[int]int) int {

	availablePhilos := make([]int, 0)
	for i := 0; i < countPhilos; i++ {

		if _, ok := philoIsFull[i]; ok {

			// fmt.Println("seeing full philo", i, " so not added to availablePhilos")
		} else {

			availablePhilos = append(availablePhilos, i)
		}
	}

	currPhiloIndex := rand.Intn(len(availablePhilos)) // randomly pick one Philosopher

	chosenPhiloIndex := availablePhilos[currPhiloIndex]

	return chosenPhiloIndex
}

func host(countPhilos, maxEatCycles int, philoIsFull map[int]int, philos []*Philo, onlyOnePhiloAvailable *bool, wg *sync.WaitGroup) {

	onePhilo := getPhilo(countPhilos, philoIsFull)
	anotherPhilo := getPhilo(countPhilos, philoIsFull)

	size_philoIsFull := len(philoIsFull)
	if size_philoIsFull == countPhilos-1 {

		*onlyOnePhiloAvailable = true
	}

	for onePhilo == anotherPhilo && *onlyOnePhiloAvailable == false { // repeat until we have two distinct Philosophers

		anotherPhilo = getPhilo(countPhilos, philoIsFull) // randomly pick one Philosopher
	}

	wg.Add(1)

	go philos[onePhilo].doEat(philoIsFull, maxEatCycles, wg)

	if *onlyOnePhiloAvailable == false {

		wg.Add(1)

		go philos[anotherPhilo].doEat(philoIsFull, maxEatCycles, wg)
	}

}

func main() {

	var wg sync.WaitGroup

	maxEatCycles := 3 //  Each philosopher should eat only 3 times
	countPhilos := 5  // how many philosophers exist

	rand.Seed(time.Now().UnixNano()) //  get fresh different random seed on each run

	CSticks := make([]*ChopStick, countPhilos)
	// allocate each CS
	for i := 0; i < countPhilos; i++ {
		CSticks[i] = new(ChopStick)
	}

	philos := make([]*Philo, countPhilos)
	// define each philos
	for i := 0; i < countPhilos; i++ {

		philos[i] = &Philo{i, CSticks[i], CSticks[(i+1)%5], 0}
	}

	//  start the dining

	philoIsFull := make(map[int]int, 0) // add Philosopher to this map when reached maximum eat cycles

	for {

		onlyOnePhiloAvailable := false
		host(countPhilos, maxEatCycles, philoIsFull, philos, &onlyOnePhiloAvailable, &wg)

		wg.Wait()

		if onlyOnePhiloAvailable {

			return // Done ... all philosophers have eaten their maximum cycles
		}
	}
}
