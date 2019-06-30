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

//Exercise2:
//This exercise will reinforce our understanding of method sets:
//create a type person struct
type Person struct {
	Name string
	Age int
}
//attach a method speak to type person using a pointer receiver

func (p *Person) speak(){
	fmt.Println("Hello from ", p.Name, "!")
}
//	create a type human interface
//	to implicitly implement the interface, a human must have the speak method
type Human interface {
	speak()
}
//	create func “saySomething”
//	have it take in a human as a parameter
//	have it call the speak method
func saySomething(h Human){
	h.speak()
}

//Exercise 3:
//See udemyAssignment9/raceCondition

//Exercise 4:
//See udemyAssignment9/mutexSynchronization

//Exercise 5:
//See udemyAssignment9/atomicSynchronization

//Exercise 6:
//Print OS and ARCH


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

	//Exercise 2:
	//	show the following in your code
	//	you CAN pass a value of type *person into saySomething
	p1 := Person{
		"Pat",
		32,
	}
	saySomething(&p1)
	//	you CANNOT pass a value of type person into saySomething
	//saySomething(p1)
	p1.speak()

	//Exercise 6:
	//Print OS and ARCH
	fmt.Println("ARCH: ",runtime.GOARCH)
	fmt.Println("OS: ", runtime.GOOS)
	
}
