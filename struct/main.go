package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	// declares and initializes the struct person in the struct order
	// alex := person{"Alex", "Hales"}

	// alex := person{firstName: "Alex", lastName: "Hales"}

	// alex.firstName = "John"
	// alex.lastName = "Wick"

	// fmt.Println(alex)

	john := person{
		firstName: "John",
		lastName:  "Wick",
		contact: contactInfo{
			email:   "johnwick@gmail.com",
			zipCode: 293525,
		},
	}

	fmt.Printf("%+v", john)
}
