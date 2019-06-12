package main

import "fmt"

func main() {

	//1.Print every number from 1 to 10,000
	for x:=0; x<=10000; x++{
		fmt.Println(x)
	}

	//2.Print every rune code point (e.g. print ASCII) of the uppercase alphabet three times.
	for x := 65; x <= 90; x++ {
		fmt.Printf("%v\n", x)
		for i:=0; i<3; i++ {
			fmt.Printf("\t%#U\n",x)
		}
	}

	//3.Create a for loop using this syntax for condition { }
	//Have it print out the years you have been alive.
	yearBorn := 1987
	for yearBorn <= 2019{
		fmt.Println(yearBorn)
		yearBorn++
	}

	//4.Create a for loop using this syntax for { }
	//Have it print out the years you have been alive.
	for {
		if yearBorn > 2019 {
			break
		}
		fmt.Println(yearBorn)
		yearBorn++
	}

	////5.Print out the remainder (modulus) which is found for each number
	//// between 10 and 100 when it is divided by 4.
	for x:=10; x<=100; x++ {
		divisor:= x%4
		fmt.Println(x,divisor)
	}

	//6.Create a program that shows the “if statement” in action.
	name1:= "Cindy"
	if name1 == "Cindy" {
		fmt.Println("Found Cindy")
	}

	//7.Building on the previous hands-on exercise, create a program that uses “else if” and “else”.
	name2:= "Pat";
	if name2 == "Cindy" {
		fmt.Println("Found Cindy")
	} else if name2 == "Roy" {
		fmt.Println("Found Roy")
	} else {
		fmt.Println("Pat wasn't found!")
	}

	//8.Create a program that uses a switch statement with no switch expression specified.
	switch {
			case true:
				fmt.Println("This is going to print!")
				fallthrough
			case 2 > 5:
				fmt.Println("This won't print! unless fallthrough")
			default:
				fmt.Println("This will only file with 'fallthrough' being prior or if everything else is false")
	}

	//9.Create a program that uses a switch statement with the switch expression
	// specified as a variable of TYPE string with the IDENTIFIER “favSport”.
	favSport := "Golf"

	switch favSport {
	case "Golf":
		fmt.Println("Favorite Sport is Golf!")
	case "Baseball":
		fmt.Println("Favorite sport is Baseball")
	default:
		fmt.Println("No favorite sport!")
	}

	//10.Write down what these print:
	//fmt.Println(true && true) --> true
	//fmt.Println(true && false) --> false
	//fmt.Println(true || true)  --> true
	//fmt.Println(true || false) -->true
	//fmt.Println(!true) --> false

	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)

	
}
