package main

import (
	"fmt"
	"math/rand"
	"sync"
)

var random = rand.New(rand.NewSource(42))

type ChopStick struct {
	sync.Mutex
}

type Request int

const (
	GetPermission Request = iota
	Thank
)

type Permission struct{}

type Host struct {
	numPermissions int
	request        chan Request
	permission     chan Permission
}

func (host *Host) GetPermission() {
	host.request <- GetPermission
	<-host.permission
}

func (host *Host) Thank() {
	host.request <- Thank
}

func (host *Host) Serve() {
	var (
		numRequests     int
		currentlyEating int
	)
	for request := range host.request {
		switch request {
		case GetPermission:
			numRequests++
		case Thank:
			currentlyEating--
		}
		for currentlyEating < host.numPermissions && numRequests > 0 {
			currentlyEating++
			numRequests--
			host.permission <- Permission{}
		}
	}
}

type Philosopher struct {
	number         int
	host           *Host
	leftChopStick  *ChopStick
	rightChopStick *ChopStick
}

func (philosopher *Philosopher) Eat(group *sync.WaitGroup) {
	philosopher.host.GetPermission()
	fmt.Println("starting to eat", philosopher.number)
	if random.Float64() < 0.5 {
		philosopher.leftChopStick.Lock()
		philosopher.rightChopStick.Lock()
	} else {
		philosopher.rightChopStick.Lock()
		philosopher.leftChopStick.Lock()
	}
	if random.Float64() < 0.5 {
		philosopher.leftChopStick.Unlock()
		philosopher.rightChopStick.Unlock()
	} else {
		philosopher.rightChopStick.Unlock()
		philosopher.leftChopStick.Unlock()
	}
	fmt.Println("finishing eating", philosopher.number)
	philosopher.host.Thank()
	group.Done()
}

func main() {
	wg := sync.WaitGroup{}
	numPhilosophers := 5
	timesToEat := 3
	host := Host{
		numPermissions: 2,
		request:        make(chan Request, numPhilosophers*timesToEat),
		permission:     make(chan Permission, numPhilosophers*timesToEat),
	}
	chopSticks := make([]*ChopStick, numPhilosophers)
	for i, _ := range chopSticks {
		chopSticks[i] = new(ChopStick)
	}
	philosophers := make([]*Philosopher, numPhilosophers)
	for i, _ := range philosophers {
		philosophers[i] = &Philosopher{
			number:         i + 1,
			host:           &host,
			leftChopStick:  chopSticks[i],
			rightChopStick: chopSticks[(i+1)%numPhilosophers],
		}
	}
	go host.Serve()
	for i := 0; i < timesToEat; i++ {
		for _, philosopher := range philosophers {
			wg.Add(1)
			go philosopher.Eat(&wg)
		}
	}
	wg.Wait()
}
