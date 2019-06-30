package main

import (
	"fmt"
	"runtime"
	"sync"
)

//Create a waitGroup
//Synchronization primitive
var wg sync.WaitGroup


func foo() {
	for i:= 0; i < 10; i++ {
		fmt.Println("foo: ", i)
	}
	//Notify the waitGroup that we are done
	wg.Done()
}

func bar() {
	for i:=0; i < 10; i ++ {
		fmt.Println("bar: ",i)
	}
}

func init(){
	fmt.Println("This func executes prior to main.")
}

func main() {

	fmt.Println("OS\t", runtime.GOOS)
	fmt.Println("ARCH\t", runtime.GOARCH)
	fmt.Println("CPUs\t", runtime.NumCPU())
	fmt.Println("GoRountines\t", runtime.NumGoroutine())

	// the go keyword acts like an & at the end of your shell command; e.g. to run the command in the background
	//If we don't create a wait group, func main() will finish execution and exit prior to foo() returning a result

	//wait for one routine
	wg.Add(1)
	go foo()
	bar()

	fmt.Println("CPUs\t", runtime.NumCPU())
	fmt.Println("GoRountines\t", runtime.NumGoroutine())

	//Wait until everything we've added to our waitGroup is done
	wg.Wait()
}
