package main

import (
	"fmt"
	"sync"
	"time"
)

const numPhilosophers = 5
const numEating = 2
const numMeals = 3

var chopsticks [numPhilosophers]sync.Mutex
var eatingHost = make(chan struct{}, numEating)

func philosopher(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < numMeals; i++ {
		// think(id)
		eat(id)
	}
}

// func think(id int) {
// 	fmt.Printf("Philosopher %d is thinking.\n", id+1)
// 	time.Sleep(time.Millisecond * 50) // Simulating thinking time
// }

func eat(id int) {
	// Request permission to eat from the host
	eatingHost <- struct{}{}

	// Pick up chopsticks
	chopsticks[id].Lock()
	chopsticks[(id+1)%numPhilosophers].Lock()

	fmt.Printf("Philosopher %d is starting to eat.\n", id+1)

	// Simulate eating time
	time.Sleep(time.Millisecond * 100)

	// Release chopsticks
	chopsticks[id].Unlock()
	chopsticks[(id+1)%numPhilosophers].Unlock()

	fmt.Printf("Philosopher %d is finishing eating.\n", id+1)

	// Release permission to eat
	<-eatingHost
}

func main() {
	var wg sync.WaitGroup

	// Initialize chopsticks
	for i := 0; i < numPhilosophers; i++ {
		chopsticks[i] = sync.Mutex{}
	}

	// Start philosophers
	for i := 0; i < numPhilosophers; i++ {
		wg.Add(1)
		go philosopher(i, &wg)
	}

	wg.Wait()
}
