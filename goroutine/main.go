package main

import (
	"fmt"
	"time"
)
 func sayHello() {	
		fmt.Println("Hello, World!")
		time.Sleep(200 * time.Millisecond)
		fmt.Println("say hello function ended")

	}
func sayHi(){
		fmt.Println("Hi, there!")
		time.Sleep(100 * time.Millisecond)
		fmt.Println("say hi function ended")
}

	func main() {
		fmt.Println("Goroutine Example")
		go sayHello()
		go sayHi()
		time.Sleep(500 * time.Millisecond)
		fmt.Println("Main function ended")

	t := time.Now()
		fmt.Println("Update started at:", t)

		go func() {
			for i := 1; i <= 5; i++ {
				fmt.Printf("Goroutine: %d\n", i)
			}
		}()
		time.Sleep(100 * time.Millisecond)

		t = time.Now()
		fmt.Println("Update ended at:", t)

		// Wait to ensure goroutine finishes before main exits
		time.Sleep(50000 * time.Millisecond)

		// Note: In real applications, use sync.WaitGroup or other synchronization methods to manage goroutines properly.

		// Uncomment the following lines to test CRUD operations
	
	}