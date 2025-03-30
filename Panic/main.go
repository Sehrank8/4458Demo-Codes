package main

import (
	"fmt"
	"time"
)

func recoverFromPanic() {
	// recover function recovers a panicking goroutine in a defered function,
	// it doesn't work outside defered functions

	if r := recover(); r != nil {
		fmt.Println("Recovered from panic:", r)
	}
}

func worker(id int) {
	// Defered functions runs when the surrounding function is exiting,
	// or the goroutine is panicking
	// they are pushed into and popped from a stack
	defer recoverFromPanic()

	fmt.Printf("Worker %d started\n", id)
	time.Sleep(time.Second)

	if id == 2 {
		panic(fmt.Sprintf("Worker %d encountered an error!", id))
	}

	fmt.Printf("Worker %d finished successfully\n", id)
}

func main() {
	fmt.Println("Starting concurrent workers...")

	for i := 1; i <= 3; i++ {
		// go command starts executing a function concurrently,
		// in this case the workers start concurrently doing their jobs.
		// every worker has its own goroutine
		go worker(i)
	}

	time.Sleep(3 * time.Second)

	fmt.Println("All workers completed (or recovered).")
}
