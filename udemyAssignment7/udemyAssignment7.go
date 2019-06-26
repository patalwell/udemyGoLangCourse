package main

import "fmt"

//Exercise 2:
//create a person struct
//create a func called “changeMe” with a *person as a parameter
//change the value stored at the *person address

type Person struct {
	firstName string
	lastName string
	age int
}

//Important: to dereference a struct, use (*value).field e.g p1.first -> (*p1).first
//“As an exception, if the type of x is a named pointer type and (*x).f is a valid selector expression denoting
//a field (but not a method), x.f is shorthand for (*x).f.” https://golang.org/ref/spec#Selectors
func changeMe(person *Person) {
	(*person).firstName = "Pat"
	(*person).lastName= "Alwell"
	(*person).age= 21
}


func main() {

	//Exercise 1:
	//Create a value and assign it to a variable.
	//Print the address of that value.
	a := 67
	fmt.Println(&a)

	//Exercise 2:

	//create a value of type person
	//print out the value
	p1 := Person{
		"John",
		"Doe",
		30,
	}

	// in func main
	//call “changeMe”
	//in func main
	//print out the value
	fmt.Println("Here is the Person object we just created:", p1)

	changeMe(&p1)

	fmt.Println("Here is the Person object after we changed its values with a pointer:", p1)


	}
