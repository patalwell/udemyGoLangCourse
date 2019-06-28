package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

//Exercise 1:
//marshal the []user to JSON. There is a little bit of a curve ball in this hands-on exercise
// - remember to ask yourself what you need to do to EXPORT a value from a package.

type user struct {
	First string
	Age   int
}

//Exercise 2:
//Struct for rawJson String

type Person struct {
	First string `json:"First"`
	Last string `json:"Last"`
	Age int `json:"Age"`
	Sayings []string `json:"Sayings"`
}

//Exercise 3:
//Starting with this code, encode to JSON the []user sending the results to Stdout.
// Hint: you will need to use json.NewEncoder(os.Stdout).encode(v interface{})

type user2 struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

//Exercise 5:
//sort the []user by
//age
//last
//Also sort each []string “Sayings” for each user
//print everything in a way that is pleasant

type ByAge []user2

func (a ByAge) Len() int { return len(a)}
func (a ByAge) Swap(i,j int) {a[i],a[j] = a[j],a[i]}
func (a ByAge) Less(i,j int) bool {return a[i].Age < a[j].Age}

type ByLast [] user2
func (a ByLast) Len() int { return len(a)}
func (a ByLast) Swap(i,j int) {a[i],a[j] = a[j],a[i]}
func (a ByLast) Less(i,j int) bool {return a[i].Last < a[j].Last}


func main() {
	u1 := user{
		First: "James",
		Age:   32,
	}

	u2 := user{
		First: "Moneypenny",
		Age:   27,
	}

	u3 := user{
		First: "M",
		Age:   54,
	}

	users := []user{u1, u2, u3}

	fmt.Println(users)

	//Exercise 1:
	//Marshall User to Json
	jsonArray, err := json.Marshal(users)
	if err != nil {
		fmt.Println("Something went wrong durng hte marshalling process: ", err)
	}

	fmt.Println("Here is the rawJson []byte: ", jsonArray)

	//Exercise 2:
	//Starting with this code unMarshall the JsonArray into a struct
	s := `[{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]},{"First":"Miss","Last":"Moneypenny","Age":27,"Sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]},{"First":"M","Last":"Hmmmm","Age":54,"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]}]`
	fmt.Println(s)


	//struct for unmarshalling our bytes
	var people []Person

	//Unmarshall bytes into a pointer of type person
	err = json.Unmarshal([]byte(s),&people)
	if err != nil {
		fmt.Println("Something went wrong during the unmarshalling process: ", err)
	}
	for _,v := range people {
		fmt.Println(v.First, v.Last, v.Age, v.Sayings[0:2])
	}

	//Exercise 3:

	u4 := user2{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u5 := user2{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u6 := user2{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users2 := []user2{u4, u5, u6}

	fmt.Println(users2)
	//encode to JSON the []user sending the results to Stdout.
	// Hint: you will need to use json.NewEncoder(os.Stdout).encode(v interface{})
	err = json.NewEncoder(os.Stdout).Encode(users2)
	if err != nil {
		fmt.Println("Something went wrong during the stdOut encoding: ", err)
	}

	//Exercise 4:
	xi := []int{5, 8, 2, 43, 17, 987, 14, 12, 21, 1, 4, 2, 3, 93, 13}
	xs := []string{"random", "rainbow", "delights", "in", "torpedo", "summers", "under", "gallantry", "fragmented", "moons", "across", "magenta"}

	fmt.Println(xi)
	// sort xi
	sort.Ints(xi)
	fmt.Println(xi)

	fmt.Println(xs)
	// sort xs
	sort.Strings(xs)
	fmt.Println(xs)

	//Exercise 5:

	users2 = []user2{u4, u5, u6}
	fmt.Println(users2)

	//Sort by Age
	fmt.Println()
	fmt.Println("Sorting by Age ...")
	sort.Sort(ByAge(users2))
	for _,v := range users2{
		fmt.Println("First Name: ", v.First)
		fmt.Println("Last Name: ", v.Last)
		fmt.Println("Age: ", v.Age)
		fmt.Println("Sayings Unsorted: ", v.Sayings)
		fmt.Println("Sorting Sayings ...")
		sort.Strings(v.Sayings)
		fmt.Println("Sayings: ", v.Sayings)
		fmt.Println()
	}

	//Sort by Last
	fmt.Println()
	fmt.Println("Sorting by LastName ...")
	sort.Sort(ByLast(users2))
	for _,v := range users2{
		fmt.Println("First Name: ", v.First)
		fmt.Println("Last Name: ", v.Last)
		fmt.Println("Age: ", v.Age)
		fmt.Println("Sayings Unsorted: ", v.Sayings)
		fmt.Println("Sorting Sayings ...")
		sort.Strings(v.Sayings)
		fmt.Println("Sayings: ", v.Sayings)
		fmt.Println()
	}



}
