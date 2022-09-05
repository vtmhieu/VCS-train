package main

import "fmt"

type customer struct {
	name        string
	email       string
	sdt         int64
	code        int64
	type_room   string
	description string
}

type room struct {
	code_room  int64
	type_room  string
	price_room int64
}

type customers []customer

func getInfomation() customer {
	var cus customer

	fmt.Println("Please fill in this form")
	fmt.Print("Name? ")
	fmt.Scan(&cus.name)
	fmt.Print("email? ")
	fmt.Scan(&cus.email)
	fmt.Print("sdt? ")
	fmt.Scan(&cus.sdt)
	fmt.Print("type_room? ")
	fmt.Scan(&cus.type_room)
	fmt.Print("description? ")
	fmt.Scan(&cus.description)

	return cus
}

func (c customer) PrintInfomation() {
	fmt.Println("Name: ", c.name)
	fmt.Println(c.name, "'s email: ", c.email)
	fmt.Println(c.name, "'s phone number: ", c.sdt)
	fmt.Println(c.name, "'s type room: ", c.type_room)
	fmt.Println("Description: ", c.description)
}

func main() {
	Customers := []customer{}

	Customers = append(Customers, getInfomation())
	for _, c := range Customers {
		fmt.Println("This is the informations of customer ", c.name)
		c.PrintInfomation()
	}
}
