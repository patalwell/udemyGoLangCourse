package main

import "fmt"

func main() {

	// Exercise 1
	//Create your own type “person” which will have an underlying type of “struct” so that it can
	// store the following data: first name, last name, favorite ice cream flavors
	// Create two VALUES of TYPE person. Print out the values, ranging over the elements in the slice
	// which stores the favorite flavors.

	type Person struct{
		firstName string
		lastName string
		favIceCreamFlavors []string
	}

	person1 := Person{
		firstName:"Cindy",
		lastName:"Clark",
		favIceCreamFlavors: []string{"Chocolate","Vanilla","Strawberry"},
	}

	person2 := Person{
		firstName:"Harry",
		lastName:"O'Connor",
		favIceCreamFlavors: []string{"Peach","Mango","Coconut"},
	}

	fmt.Println("First Name:", person1.firstName, "\nLast Name:", person1.lastName)
	for k,v := range person1.favIceCreamFlavors{
		fmt.Println(k,v)
	}
	fmt.Println("---------")

	fmt.Println("First Name:", person2.firstName, "\nLast Name:", person2.lastName)
	for k,v := range person2.favIceCreamFlavors{
		fmt.Println(k,v)
	}
	fmt.Println("---------")


	// Exercise 2
	// Take the code from the previous exercise, then store the values of type person in a map with the key of last name.
	// Access each value in the map. Print out the values, ranging over the slice.

	mapOfPeople := map[string]Person{
		person1.lastName:person1,
		person2.lastName:person2,
	}

	for k,v := range mapOfPeople{
		fmt.Println(k,v)
		for i,v2 := range v.favIceCreamFlavors{
			fmt.Println(i,v2)
		}
		fmt.Println("---------")
	}

	// Exercise 3
	// Create a new struct type: vehicle that has the following fields: doors, color

	type Vehicle struct {
		numDoors int
		color string
	}

	// Create two new struct types: truck & sedan. Embed the “vehicle” type in both truck & sedan.
	// Give truck the field “fourWheel” which will be set to bool.

	type Truck struct {
		Vehicle
		fourWheel bool
	}

	// Give sedan the field “luxury” which will be set to bool.

	type Sedan struct {
		Vehicle
		luxury bool
	}

	// Using the vehicle, truck, and sedan structs:
	// Using a composite literal, create a value of type truck and assign values to the fields;
	// Print out each of these values.
	// Print out a single field from each of these values.

	truck1 := Truck{
		Vehicle: Vehicle{
			numDoors:2,
			color:"black",
		},
		fourWheel: true,
	}
	fmt.Println(truck1)
	fmt.Println("Truck details:", truck1.color,truck1.numDoors,truck1.fourWheel)
	// Using a composite literal, create a value of type sedan and assign values to the fields.
	// Print out each of these values.
	// Print out a single field from each of these values.

	sedan1 := Sedan{
		Vehicle: Vehicle{
			numDoors:4,
			color:"blue",
		},
		luxury:true,
	}
	fmt.Println(sedan1)
	fmt.Println("Sedan details:", sedan1.color,sedan1.numDoors,sedan1.luxury)

	// Exercise 4
	// Create and use an anonymous struct.

	//Student Notes: An anonymous struct is useful for localizing state

	car := struct {
		brand string
		numDoors int8
		body string
		cylinders int8
	}{
		brand:"BMW",
		numDoors: 2,
		body:"Coupe",
		cylinders: 8,
	}

	fmt.Println("Here is the anonymous struct: ", car)




	
}
