package main

import "fmt"

func main() {

	// Exercise 1
	// Create an array which holds 5 values of type int
	// Assign values to each index position
	// range over the array and print the values
	// using format printing print the type of the Array

	 var firstArray [5]int
	 	firstArray[0] = 1
	 	firstArray[1] = 2
	 	firstArray[2] = 3
	 	firstArray[3] = 4
	 	firstArray[4] = 5

	 	for i, v := range firstArray {
	 		fmt.Println(i,v)
		}

	 	fmt.Printf("Here is the Array's type: %T",firstArray)

	// Exercise 2
	// Create a slice of type int and assign ten values
	// range over the slice and print the values
	// using format printing print the type of the slice

	firstSlice := []int {42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	for i, v := range firstSlice {
		fmt.Println(i,v)
	}

	fmt.Printf("Here is the Slice's type: %T\n", firstSlice)


	// Exercise 3
	// Using the code from the previous example, use SLICING to create the following
	// new slices which are then printed:
	//[42 43 44 45 46]
	secondSlice := firstSlice[0:5]
	fmt.Println(secondSlice)
	//[47 48 49 50 51]
	thirdSlice := firstSlice[5:10]
	fmt.Println(thirdSlice)
	//[44 45 46 47 48]
	fourthSlice := firstSlice[2:7]
	fmt.Println(fourthSlice)
	//[43 44 45 46 47]
	fifthSlice := firstSlice[1:6]
	fmt.Println(fifthSlice)

	// Exercise 4
	// start with this slice x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	// append 52 to that slice
	x = append(x, 52)
	// print out the slice
	fmt.Println(x)
	// in ONE STATEMENT append 53 54 55
	x = append(x, 53, 54, 55)
	// print out the slice
	fmt.Println(x)
	// append to the slice y := []int{56, 57, 58, 59, 60}
	y := []int{56, 57, 58, 59, 60}
	x = append(x, y...)
	// print out the slice
	fmt.Println(x)

	// Exercise 5
	// To DELETE from a slice, we use APPEND along with SLICING. For this hands-on exercise, follow these steps:
	//start with this slice x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x = []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	//use APPEND & SLICING to get these values here which you should ASSIGN to a variable “y” and then print:
	y = append(x[:3],x[6:10]...)
	//[42, 43, 44, 48, 49, 50, 51]
	fmt.Println(y)

	// Exercise 6
	// Create a slice to store the names of all of the states in the United States of America.
	// What is the length of your slice? What is the capacity?
	// Print out all of the values, along with their index position in the slice,
	// without using the range clause.

	states := []string{` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`,
		` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`,
		` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`,
		` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`,
		` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`,
		` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, ` Utah`,
		` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`}

	fmt.Println("Here is the length of the state slice: ", len(states))
	fmt.Println("Here is the capacity of the state slice: ", cap(states))

	for i := 0; i < len(states); i++ {
		fmt.Println(i, states[i])
	}

	// Exercise 7
	// Create a slice of a slice of string ([][]string).
	// Store the following data in the multi-dimensional slice:
	// "James", "Bond", "Shaken, not stirred"
	// "Miss", "Moneypenny", "Helloooooo, James."
	// Range over the records, then range over the data in each record.

	jamesSlice := []string{"James", "Bond"  ,"Shaken, not stirred"}
	moneyPennySlice := []string{"Moneypenny", "Helloooooo, James."}
	multiDimensionalSlice := [][]string{jamesSlice,moneyPennySlice}

	for index1,v := range multiDimensionalSlice {
		fmt.Println("Record: ", index1)
		for index2,val := range v{
			fmt.Println("Index: ",index2,val)
		}
	}

	// Exercise 8
	// Create a map with a key of TYPE string which is a person’s “last_first” name,
	// and a value of TYPE []string which stores their favorite things.
	// Store three records in your map. Print out all of the values, along with their index position in the slice.

	firstMap := map[string][]string{
		"Alwell_Patrick":{"sushi","baseball","babes"},
		"Doe_John": {"pizza","beer","surf"},
	}

	for k, v := range firstMap {
		fmt.Println(k,":",v)
		for i, val := range v {
			fmt.Println(i,":",val)
		}
	}

	// Exercise 9
	// Using the code from the previous example, add a record to your map.
	// Now print the map out using the “range” loop

	firstMap["New_Person"] = []string{"liquor","cookies","cake"}

	for k,v := range firstMap{
		fmt.Println(k,v)
	}

	// Exercise 10
	// Using the code from the previous example, delete a record from your map.
	// Now print the map out using the “range” loop

	delete(firstMap,"New_Person")

	for k,v := range firstMap{
		fmt.Println(k,v)
	}

}
