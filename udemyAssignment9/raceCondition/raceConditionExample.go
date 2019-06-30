package main

/*
 * We are creating a race condition on purpose to show what happens when goRoutines share the same variable and
 * attempt to increment said shared variable
 *
 */

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	//Get Sys info
	fmt.Println("CPUs: ", runtime.NumCPU())
	fmt.Println("GoRoutines: ", runtime.NumGoroutine())

	//Define a generic counter object
	counter := 0

	//waitGroup
	var waitGroup sync.WaitGroup

	//goRoutine quantity
	const GOROUTINES = 100

	//Add the routines to our waitGroup
	//Will wait to exit main() until we've hit the same number of wg.Done()
	waitGroup.Add(GOROUTINES)

	for i:= 0; i < GOROUTINES; i++{
		go func() {
			//create a local value and assign it to the counter
			instanceValue := counter
			//yield the routine to allow other routines to run
			//we could also sleep the routine time.Sleep(time.Second)
			runtime.Gosched()
			instanceValue ++
			//increment and reassign counter
			//Note: Each goRoutine is going to have a different instance and counter value
			counter = instanceValue
			waitGroup.Done()
		}()
		fmt.Println("numGoRoutines", runtime.NumGoroutine())
	}
	//Wait for my routines to finish
	waitGroup.Wait()

	fmt.Println("numGoRoutines", runtime.NumGoroutine())
	fmt.Println("Count: ", counter)
	
}
