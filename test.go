// Golang program to illustrate the usage of
// AfterFunc() function

// Including main package
package main

// Importing fmt and time
import (
	"fmt"
	"time"
)

// Main function
func xxx() {
	// Creating channel using
	// make keyword
	mychan:= make(chan int)

	// Calling AfterFunc() method
	// with its parameters
	time.AfterFunc(6*time.Second, func() {

		// Printed after stated duration
		// by AfterFunc() method is over
		fmt.Println("6 seconds over....")

		// loop stops at this point
		mychan <- 30
	})

	// Calling for loop
	for {

		// Select statement
		select {

		// Case statement
		case n:= <-mychan:

			// Printed after the loop stops
			fmt.Println(n, "is arriving")
			fmt.Println("Done!")
			return

		// Returned by default
		default:

			// Printed until the loop stops
			fmt.Println("time to wait")

			// Sleeps for 3 seconds
			time.Sleep(3 * time.Second)
		}
	}
}