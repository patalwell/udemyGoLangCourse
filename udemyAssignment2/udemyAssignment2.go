package main

import "fmt"

func main() {

	//1. Write a program that prints a number in decimal, binary, and hex
	x :=42
	fmt.Printf("%d\n",x)
	fmt.Printf("%b\n",x)
	fmt.Printf("%#x\n",x)

	//2. Using the following operators, write expressions and assign their values to variables:
	// ==
	a := 4 == 2
	// <=
	b:= 3 <= 5
	// >=
	c:= 45 >= 44
	// !=
	d:=  1 != 9
	// <
	e:= 1 <= 10
	// >
	f:= 22 > 45
	//Now print each of the variables.
	fmt.Println(a,b,c,d,e,f)

	//3. Create TYPED and UNTYPED constants. Print the values of the constants.
	const (
		untypedConstant = 0
		typedConstant int = -12
	)
	fmt.Println(untypedConstant,typedConstant)

	//4. Write a program that
	// assigns an int to a variable
	var myVar uint = 42;
	// prints that int in decimal, binary, and hex
	fmt.Printf("%d\t%b\t#%x\t\n",myVar,myVar,myVar)
	// shifts the bits of that int over 1 position to the left, and assigns that to a variable
	myNewVar :=  myVar << 1
	// prints that variable in decimal, binary, and hex
	fmt.Printf("%d\t%b\t#%x\t",myNewVar,myNewVar,myNewVar)

	//5.Create a variable of type string using a raw string literal. Print it.
	myString := "Hello World"
	fmt.Println(myString)

	//6.Using iota, create 4 constants for the NEXT 4 years. Print the constant values.
	const(
		_ = iota
		firstYear = 2019 + iota
		secondYear = 2019 + iota
		thirdYear = 2019 + iota
		fourthYear = 2019 + iota
	)
	fmt.Println(firstYear,secondYear,thirdYear,fourthYear)
}
