package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
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
		contactInfo: contactInfo{
			email:   "johnwick@gmail.com",
			zipCode: 293525,
		},
	}

	john.updatePersonInfo()

	john.print()
}

func (pointerToPerson *person) updatePersonInfo() {

	var fn, ln, e string
	var z int

	fmt.Printf("Write the updated information in the following order:\nFirst_Name Last_Name Email Zip_Code\n")
	fmt.Scanf("%v %v %v %v", &fn, &ln, &e, &z)

	(*pointerToPerson).firstName = fn
	(*pointerToPerson).lastName = ln
	(*pointerToPerson).contactInfo.email = e
	(*pointerToPerson).contactInfo.zipCode = z

}

func (p person) print() {
	fmt.Printf("First Name: %v\nLast Name: %v\nEmail: %v\nZip Code: %v", p.firstName,
		p.lastName, p.contactInfo.email, p.contactInfo.zipCode)
}
