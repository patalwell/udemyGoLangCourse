package main

/*
 * We are preventing a race condition with a mutex type to show what happens when goRoutines need to share a variable
 * and attempt to increment said shared variable
 *
 * Note: A Mutex is a mutual exclusion lock. The zero value for a Mutex is an unlocked mutex.
 * see https://godoc.org/sync#Mutex
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

	//Create a mutex to prevent issues with sharing a variable like counter
	var mutex sync.Mutex

	for i:= 0; i < GOROUTINES; i++{
		go func() {
			//lock the counter to a routine at a time
			mutex.Lock()

			//create a local value and assign it to the counter
			instanceValue := counter

			//yield the routine to allow other routines to run
			//we could also sleep the routine time.Sleep(time.Second)
			runtime.Gosched()
			instanceValue ++

			//increment and reassign counter
			//Note: Each goRoutine is going to have a different instance and counter value
			counter = instanceValue

			//unlock the counter for use by another routine
			mutex.Unlock()

			//notify the wg we are done with this task
			waitGroup.Done()
		}()
		fmt.Println("numGoRoutines", runtime.NumGoroutine())
	}
	//Wait for my routines to finish
	waitGroup.Wait()

	fmt.Println("numGoRoutines", runtime.NumGoroutine())
	fmt.Println("Count: ", counter)

}

