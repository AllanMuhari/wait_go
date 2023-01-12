package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	// Set the number of Goroutines we want to wait for
	wg.Add(3)

	// Start three Goroutines
	go func() {
		fmt.Println("Goroutine 1 starting")
		time.Sleep(1 * time.Second)
		fmt.Println("Goroutine 1 finished")
		wg.Done()
	}()
	go func() {
		fmt.Println("Goroutine 2 starting")
		time.Sleep(2 * time.Second)
		fmt.Println("Goroutine 2 finished")
		wg.Done()
	}()
	go func() {
		fmt.Println("Goroutine 3 starting")
		time.Sleep(3 * time.Second)
		fmt.Println("Goroutine 3 finished")
		wg.Done()
	}()

	// Wait for all the Goroutines to finish
	wg.Wait()

	fmt.Println("All Goroutines finished")
}
