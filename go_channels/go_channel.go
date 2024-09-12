package main

import (
	"fmt"
	"sync"
)

func main() {
	// Create a channel to communicate between the producer and consumer goroutines
	ch := make(chan int)

	// WaitGroup to ensure the main function waits for all goroutines to finish
	var wg sync.WaitGroup
	wg.Add(2) // We have two goroutines to wait for

	// Producer goroutine: generates numbers from 1 to 10
	go func() {
		defer wg.Done() // Mark this goroutine as done when it finishes
		for i := 1; i <= 10; i++ {
			ch <- i // Send the number to the channel
		}
		close(ch) // Close the channel after sending all numbers
	}()

	// Consumer goroutine: reads from the channel and prints the numbers
	go func() {
		defer wg.Done()       // Mark this goroutine as done when it finishes
		for num := range ch { // Continuously read from the channel until it's closed
			fmt.Println(num) // Print the received number
		}
	}()

	// Wait for both goroutines to finish
	wg.Wait()
}
