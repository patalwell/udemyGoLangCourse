package main

import (
	"fmt"
	"runtime"
	"sync"
)

//Exercise 1:
//in addition to the main goroutine, launch two additional goroutines
//each additional goroutine should print something out
func routineOne() {
	fmt.Println("goRoutineOne")
}
func routineTwo(){
	fmt.Println("goRoutineTwo")
}
//use waitgroups to make sure each goroutine finishes before your program exists

var waitGroup sync.WaitGroup
func main() {

	//Exercise 1:
	//Add two routines to the waitGroup
	waitGroup.Add(2)
	go func() {
		routineOne()
		waitGroup.Done()
		fmt.Println("GoRoutines: ", runtime.NumGoroutine())
	}()

	go func() {
		routineTwo()
		waitGroup.Done()
		fmt.Println("GoRoutines: ", runtime.NumGoroutine())
	}()

	//Wait till our routines run prior to exiting main()
	waitGroup.Wait()
	fmt.Println("GoRoutines: ", runtime.NumGoroutine())
	
}
