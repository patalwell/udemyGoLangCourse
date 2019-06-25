package main

import (
	"fmt"
	"math"
)

//Challenge: Create a function that has a callback to add up odd numbers
func odd(f func(xi ...int) int, vi ...int) int {
	var yi []int
	for _, v := range vi {
		if v%2 != 0 {
			yi = append(yi, v)
		}
	}
	//we are returning the values of the yi slice since we are calling the function's first argument
	return f(yi...)
}

//Exercise 1
//create a func with the identifier foo that returns an int
func foo() int {
	return 23
}
//create a func with the identifier bar that returns an int and a string
func bar() (int, string) {
	return 21, "Twenty-One"
}
//call both funcs
//print out their results

//Exercise 2
//create a func with the identifier sum that
//takes in a variadic parameter of type int
func sum(xi ...int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}
//create a func with the identifier sum2 that
//takes in a parameter of type []int
func sum2( sliceX []int) int {
	total := 0
	for _, v := range sliceX {
		total += v
	}
	return total
}
//returns the sum of all values of type int passed in

//Exercise 3:

func fooDefered() {
	defer func() {
		fmt.Println("foo DEFER ran")
	}()
	fmt.Println("foo ran")
}

//Exercise 4:
type Person struct {
	firstName string
	lastName string
	age int
}

func (p Person) speak() {
	fmt.Println("The person's name is :", p.firstName, p.lastName, "and the person's age is :", p.age)
}

//Exercise 5:
//create a type SQUARE
type Square struct{
	length float64
	width float64
}
//create a type CIRCLE
type Circle struct{
	raidus float64
}
//attach a method to each that calculates AREA and returns it
//square area = L * W

func (s Square) area() float64 {
	return s.length * s.width
}

//circle area= π r 2
func (c Circle) area() float64 {
	return math.Pi * (math.Pow(c.raidus,2))

}
//create a type SHAPE that defines an interface as anything that has the AREA method
type Shape interface {
	area() float64
}
//create a func INFO which takes type shape and then prints the area
func info(s Shape){
	fmt.Println("Here is the area: ", s.area())
}

//Exercise 8:
func returnAFunc() func() int{
	return func() int {
		return 34
	}
}

//Exercise 9:
//A “callback” is when we pass a func into a func as an argument. For this exercise,
//pass a func into a func as an argument

func callBack( f func(x int) int, argVal int) int {
	total := f(argVal)
	return  total
}

//Exercise 10:
//Closure is when we have “enclosed” the scope of a variable in some code block. For this hands-on exercise,
// create a func which “encloses” the scope of a variable:

func enclosing( x int) int {
	var y int
	fmt.Println("Y will remain as 43 within the function")
	y = 43
	return y + x
}

func main() {

	//Callback Challenge
	intSlice := []int{1,2,3,4,5,6,7,8,9}
	f := sum(intSlice...)
	fmt.Println("Here is the sum of the slice",f)
	f2 := odd(sum,intSlice...)
	fmt.Println("Here is the sum of the odd numbers: ", f2)

	//Exercise 1:
	//call foo() and bar()
	//print out their results
	fmt.Println("Here are the results of foo: ", foo())
	x, s := bar()
	fmt.Println("Here are the results of bar: ", x , s)

	//Exercise 2:
	//pass in a value of type []int into your func sum() (unfurl the []int)
	//returns the sum of all values of type int passed in
	sumAnswer := sum(intSlice...)
	fmt.Println("Here is the sum of the ints unfurled", sumAnswer)
	sum2Answer := sum2(intSlice)
	fmt.Println("Here is the sum of the slice", sum2Answer)

	//Exercise 3:
	//Use the “defer” keyword to show that a deferred func runs after the func containing it exits.
	defer fooDefered()

	//Exercise 4:
	//Create a user defined struct with
	//	the identifier “person”
	//	the fields: first, last, age
	//	attach a method to type person with
	//	the identifier “speak”
	//	the method should have the person say their name and age
	//	create a value of type person
	//	call the method from the value of type person

	person1 := Person{
		firstName:"Pat",
		lastName:"Alwell",
		age:32,
	}

	person1.speak()

	//Exercise 5:
	//create a value of type square
	square := Square{
		length:12,
		width:13,
	}
	//create a value of type circle
	circle := Circle{
		raidus:3.45,
	}
	//use func info to print the area of square
	info(square)
	//use func info to print the area of circle
	info(circle)

	//Exercise 6:
	//Build and use an anonymous func
	func() {
		fmt.Println("I am an anonymous function")
	}()

	//Exercise 7:
	//Assign a func to a variable, then call that func
	fooFunc := foo()
	fmt.Println("Here is an assigned function", fooFunc)

	//Exercise 8:
	//Create a func which returns a func
	//assign the returned func to a variable
	//call the returned func
	answerToReturnAFunc := returnAFunc()
	fmt.Println("Here is the value of the func within a func ", answerToReturnAFunc())

	//Exercise 9:
	//A “callback” is when we pass a func into a func as an argument. For this exercise,
	//pass a func into a func as an argument
	callBackFunc := func(x int) int{
		return x
	}
	fmt.Println("Here is the value of the callback:", callBack(callBackFunc,12))


	//Exercise 10:
	a := enclosing(100)
	fmt.Println("This is 143: ",a)
	b := enclosing(0)
	fmt.Println("This is 43: ", b)

	}

