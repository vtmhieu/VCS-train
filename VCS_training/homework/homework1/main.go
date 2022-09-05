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

func (Customers customer) getInfomation() []customer {
	cus := []customer{}
	fmt.Println("Please fill in this form")
	fmt.Print("Name? ")
	fmt.Scan(&cus[len(cus)].name)
	fmt.Print("email? ")
	fmt.Scan(&cus[len(cus)].email)
	fmt.Print("sdt? ")
	fmt.Scan(&cus[len(cus)].sdt)
	fmt.Print("type_room? ")
	fmt.Scan(&cus[len(cus)].type_room)
	fmt.Print("description? ")
	fmt.Scan(&cus[len(cus)].description)

	return cus
}

func (Customers customer) PrintInfomation() {

}

func main() {
	Customers := []customer{}
	fmt.Println("Welcome to our hotel!")
	fmt.Println("Please enter your name and email address:")

	for _, v := range Customers {
		Customers = append(Customers, Customers.getInfomation())
	}
}
