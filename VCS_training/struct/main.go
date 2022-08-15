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

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
	//*pointerToPerson is an operator- it means we want to manipulate the value the pointer is referencing
	//*person is the type description, it means we are working with a pointer to a person
}

//  &variable give me the memory address of the value this variable is pointing at
// *pointer give me the value this memory address is pointing at
func (p person) print() {
	fmt.Printf("%+v", p)
}

func main() {
	//tell go what different fields of a person/struct should have
	//create a new value of type person

	//1st way

	// alex := person{"Hieu", "Vu"}
	// fmt.Println(alex)

	//2nd way
	// var alex person
	// alex.firstName = "Hieu"
	// alex.lastName = "Vu"
	// fmt.Println(alex)
	// fmt.Printf("%+v", alex)

	jim := person{
		firstName: "Jim",
		lastName:  "Hieu",
		contact: contactInfo{
			email:   "hieu@gmail.com",
			zipCode: 23323,
		},
	}

	jimPointer := &jim
	// take the address of jim
	jimPointer.updateName("Hie")
	jim.print()
}

// turn address into value of that address = *address
// turn value into the address of that value = &value
