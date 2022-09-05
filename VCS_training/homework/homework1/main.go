package main

import "fmt"

type customer struct {
	name        string `json:"name"`
	email       string `json:"email"`
	sdt         int64  `json:"sdt"`
	code        int64  `json:"code"`
	type_room   string `json:"type_room"`
	description string `json:"description"`
}

type customers struct {
	customers []customer `json:"customers"`
}

type room struct {
	code_room  int64
	type_room  string
	price_room int64
}

func get_cus_Infomation() customer {
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

func (c customer) Print_cus_Infomation() {
	fmt.Println("Name: ", c.name)
	fmt.Println(c.name, "'s email: ", c.email)
	fmt.Println(c.name, "'s phone number: ", c.sdt)
	fmt.Println(c.name, "'s type room: ", c.type_room)
	fmt.Println("Description: ", c.description)
}

func SavetoFile() {

}

func main() {
	Customers := []customer{}

	var n int
	fmt.Println("----------------------------------------")
	fmt.Println("MENU")
	fmt.Println("1.Enter new room")
	fmt.Println("2.Enter new customer")
	fmt.Println("3.Book")
	fmt.Println("4.Sort room booking by room type")
	fmt.Println("5.Create bill for customer")
	fmt.Println("6.Out program")
	fmt.Print("Your choice? ")
	fmt.Scan(&n)

	switch n {
	case 1:
	case 2:
		Customers = append(Customers, get_cus_Infomation())

		main()
	case 3:
	case 4:
	case 5:
	case 6:
		return
	default:
		main()
	}
}
