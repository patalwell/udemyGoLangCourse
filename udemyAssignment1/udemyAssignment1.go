package main

import "fmt"


//Exercise 2: Use var to DECLARE three VARIABLES. The variables should have package level scope. Do not assign VALUES to the variables.
// Use the following IDENTIFIERS for the variables and make sure the variables are of the following
// TYPE (meaning they can store VALUES of that TYPE).
//NOTE: Using a,b,c since we are using the same package in directory

//identifier “x” type int
var a int
//identifier “y” type string
var b string
//identifier “z” type bool
var c bool

//Exercise 3: Using the code from the previous exercise,
//At the package level scope, assign the following values to the three variables
//for x assign 42
var x1 = 42
//for y assign “James Bond”
var y1 = "James Bond"
//for z assign true
var z1 = true

//Exercise 4: Create your own type. Have the underlying type be an int.
//create a VARIABLE of your new TYPE with the IDENTIFIER “xType” using the “VAR” keyword

type myType int
var xType myType

//Exercise 5: at the package level scope, using the “var” keyword, create a VARIABLE with the IDENTIFIER “y”.
//The variable should be of the UNDERLYING TYPE of your custom TYPE “x”.

var yType int

var abd int



func main() {
	//Exercise 1: Using the short declaration operator, ASSIGN these VALUES to VARIABLES with the IDENTIFIERS “x” and “y” and “z”
	//42
	x := 42
	//"James Bond"
	y := "James Bond"
	//true
	z := true

	//Exercise 1: Now print the values stored in those variables using
	//a single print statement
	fmt.Print(x, y, z,"\n")
	// print statements
	fmt.Println("Here is the value of x:", x)
	fmt.Printf("Here is the value of y: %v and y's type: %T\n",y,y)
	fmt.Printf("Here is the value of z: %v and the binary value of z: %b",z,z)

	//Exercise 2: in func main
	//print out the values for each identifier
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	//The compiler assigned values to the variables. What are these values called?
	//They are known as zero values since they ARE NOT assigned a VALUE during ASSIGNMENT DECLARATION

	//Exercise 3: in func main
	//use fmt.Sprintf to print all of the VALUES to one single string. ASSIGN the returned value of TYPE string using
	//the short declaration operator to a VARIABLE with the IDENTIFIER “s”. print out the value stored by variable “s”

	s := fmt.Sprintf("%v\t%v\t%v\t",x1,y1,z1)
	fmt.Println(s)

	//Exercise 4: in func main
	//print out the value of the variable “x”
	fmt.Println(xType)
	//print out the type of the variable “x”
	fmt.Printf("%T\n",xType)
	//assign 42 to the VARIABLE “x” using the “=” OPERATOR
	xType = 42
	//print out the value of the variable “x”
	fmt.Println(xType)

	//Exercise 5: now use CONVERSION to convert the TYPE of the VALUE stored in “x” to the UNDERLYING TYPE
	//then use the “=” operator to ASSIGN that value to “y”
	yType = int(xType)
	//print out the value stored in “y”
	fmt.Println(yType)
	//print out the type of “y”
	fmt.Printf("%T",yType)

}
