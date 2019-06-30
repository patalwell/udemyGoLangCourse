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
	"sync/atomic"
)

func main() {

	//Get Sys info
	fmt.Println("CPUs: ", runtime.NumCPU())
	fmt.Println("GoRoutines: ", runtime.NumGoroutine())

	//Define an int64 counter for pkg atomic
	var counter int64

	//waitGroup
	var waitGroup sync.WaitGroup

	//goRoutine quantity
	const GOROUTINES = 100

	//Add the routines to our waitGroup
	//Will wait to exit main() until we've hit the same number of wg.Done()
	waitGroup.Add(GOROUTINES)

	for i:= 0; i < GOROUTINES; i++{
		go func() {
			//create an atomic object with a pointer to the atomic var
			//and the amount you wish to increment by e.g. 1
			atomic.AddInt64(&counter, 1)
			fmt.Println("Counter: ", atomic.LoadInt64(&counter))
			//yield the routine to allow other routines to run
			//we could also sleep the routine time.Sleep(time.Second)
			runtime.Gosched()
			waitGroup.Done()
		}()
		fmt.Println("numGoRoutines", runtime.NumGoroutine())
	}
	//Wait for my routines to finish
	waitGroup.Wait()

	fmt.Println("numGoRoutines", runtime.NumGoroutine())
	fmt.Println("Count: ", counter)

}
