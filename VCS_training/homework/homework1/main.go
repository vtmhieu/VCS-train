package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

type customer struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Sdt   int    `json:"sdt"`
	ID    int    `json:"id"`
}

// type customers struct {
// 	customers []customer `json:"customers"`
// }1

type room struct {
	Id_room    int    `json:"room_id"`
	Type_room  string `json:"room_type"`
	Price_room int    `json:"room_price"`
}

type book struct {
	Cus_name   string `json:"cus_name"`
	Cus_id     int    `json:"cus_id"`
	Cus_email  string `json:"cus_email"`
	Cus_sdt    int    `json:"cus_sdt"`
	Room_id    int    `json:"room_id"`
	Room_type  string `json:"room_type"`
	Room_price int    `json:"room_price"`
}

func nextLine() string {
	bio := bufio.NewReader(os.Stdin)
	line, _, _ := bio.ReadLine()
	return string(line)
}
func readFile(filename string) []byte {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}
	return content
}

func SaveRoom(filename string, rooms []room) {
	data, err := json.MarshalIndent(rooms, "", " ")
	if err != nil {
		fmt.Println(err)
	}

	err = ioutil.WriteFile(filename, data, 0666)
	if err != nil {
		fmt.Println(err)
	}
}

func LoadRoom(filename string) []room {

	var rooms []room
	data := readFile(filename)
	_ = json.Unmarshal(data, &rooms)
	return rooms
}

func SaveCus(filename string, customers []customer) {
	data, err := json.MarshalIndent(customers, "", " ")
	if err != nil {
		fmt.Println(err)
	}

	err = ioutil.WriteFile(filename, []byte(data), 0666)
	if err != nil {
		fmt.Println(err)
	}
}

func LoadCus(filename string) []customer {
	var cus []customer

	data := readFile(filename)
	_ = json.Unmarshal(data, &cus)
	return cus
}

func SaveBook(filename string, books []book) {
	data, err := json.MarshalIndent(books, "", " ")
	if err != nil {
		fmt.Println(err)
	}

	err = ioutil.WriteFile(filename, []byte(data), 0666)
	if err != nil {
		fmt.Println(err)
	}
}

func LoadBook(filename string) []book {
	var books []book

	data := readFile(filename)
	_ = json.Unmarshal(data, &books)
	return books
}

func InsertRoomID(Newroom *room, err error, rooms []room) int {
	fmt.Print("ID: ")
	Newroom.Id_room, err = strconv.Atoi(nextLine())
	if err != nil {
		fmt.Println("The input must be integer")
		fmt.Println("Please reinsert the room ID ")
		InsertRoomID(Newroom, err, rooms)
	}
	if Newroom.Id_room < 0 {
		fmt.Println("The input must be positive number")
		fmt.Println("Please reinsert the room ID ")
		InsertRoomID(Newroom, err, rooms)
	}
	n := 0
	for _, v := range rooms {
		if Newroom.Id_room == v.Id_room {
			n++
		}
	}
	if n > 0 {
		fmt.Println("The ID has been used. Please reinsert the room ID.")
		InsertRoomID(Newroom, err, rooms)
	}
	if n == 0 {
		return Newroom.Id_room
	}
	return Newroom.Id_room
}

func InsertRoomType(Newroom *room, err error) string {
	fmt.Print("1. Single  2. Double  3. President ? ")
	var k int
	fmt.Scanln(&k)

	switch k {
	case 1:
		Newroom.Type_room = "Single"

	case 2:
		Newroom.Type_room = "Double"

	case 3:
		Newroom.Type_room = "President"

	default:
		fmt.Println("Please enter room type in range 1 to 3")
		InsertRoomType(Newroom, err)
	}
	return Newroom.Type_room
}

func InsertRoomPrice(Newroom *room, err error) int {
	Newroom.Price_room, err = strconv.Atoi(nextLine())
	if err != nil {
		fmt.Println("Input must be number")
		InsertRoomPrice(Newroom, err)
	}
	if Newroom.Price_room < 0 {
		fmt.Println("Input must be positive number")
		InsertRoomPrice(Newroom, err)
	}
	return Newroom.Price_room
}

func InsertCusName(Newcustomer *customer) string {
	fmt.Println("Name: ")
	Newcustomer.Name = nextLine()
	if Newcustomer.Name == "" {
		fmt.Println("Please reinsert customer name.")
		InsertCusName(Newcustomer)
	}
	return Newcustomer.Name
}

func InsertCusEmail(Newcustomer *customer) string {
	fmt.Println("Email:")
	Newcustomer.Email = nextLine()
	if len(Newcustomer.Email) <= 10 || Newcustomer.Email[(len(Newcustomer.Email)-10):(len(Newcustomer.Email))] != ("@gmail.com") {
		fmt.Println("Please reinsert customer email.")
		fmt.Println("The email must have the tail of @gmail.com")
		InsertCusEmail(Newcustomer)
	}
	return Newcustomer.Email
}

func InsertCusPhone(Newcustomer *customer, err error) int {
	fmt.Println("Phone number:")
	Newcustomer.Sdt, err = strconv.Atoi(nextLine())
	if err != nil {
		fmt.Println("The input must be number.")
		fmt.Println("Please reinsert customer phone.")
		InsertCusPhone(Newcustomer, err)
	}
	if Newcustomer.Sdt < 0 || len(strconv.Itoa(Newcustomer.Sdt)) != 8 {
		fmt.Println("Wrong input. Input must have 8 numbers")
		fmt.Println("Please reinsert customer phone.")
		InsertCusPhone(Newcustomer, err)
	}
	return Newcustomer.Sdt
}

func InsertCusID(Newcustomer *customer, err error, customers []customer) int {
	fmt.Println("ID:")
	Newcustomer.ID, err = strconv.Atoi(nextLine())
	if err != nil {
		fmt.Println("The input must be integer")
		fmt.Println("Please reinsert the room ID.")
		InsertCusID(Newcustomer, err, customers)
	}
	if Newcustomer.ID < 0 {
		fmt.Println("The input must be positive number")
		fmt.Println("Please reinsert the room ID ")
		InsertCusID(Newcustomer, err, customers)
	}
	n := 0
	for _, v := range customers {
		if Newcustomer.ID == v.ID {
			n++
		}
	}
	if n > 0 {
		fmt.Println("The ID has been used. Please reinsert the room ID.")
		InsertCusID(Newcustomer, err, customers)
	}
	if n == 0 {
		return Newcustomer.ID
	}
	return Newcustomer.ID
}

// book:
// 1.find the customer by using id , if not found -> insert new customer
// 2. find the room by type and book.
func findCustomer(Newbook *book, Cus []customer, err error) (string, int, string, int) {
	fmt.Println("Please enter customer ID:")
	Newbook.Cus_id, err = strconv.Atoi(nextLine())
	if err != nil {
		fmt.Println("Please enter number")
		fmt.Println("Please reinsert the customer ID")
		findCustomer(Newbook, Cus, err)
	}
	if Newbook.Cus_id < 0 {
		fmt.Println("The input must be positive number")
		fmt.Println("Please reinsert the customer ID ")
		findCustomer(Newbook, Cus, err)
	}
	n := 0
	var k customer
	for _, v := range Cus {
		if Newbook.Cus_id == v.ID {
			k = v
			n++
		}
	}
	if n == 0 {
		fmt.Println("Please insert a new customer")

		var Newcus customer
		InsertCusName(&Newcus)

		InsertCusEmail(&Newcus)

		InsertCusPhone(&Newcus, err)

		InsertCusID(&Newcus, err, Cus)

		fmt.Println("Successful!")
		Cus = append(Cus, Newcus)
		SaveCus("customer.json", Cus)
		Cus = LoadCus("customer.json")
		fmt.Println("Please reinsert the customer ID to find the customer")
		findCustomer(Newbook, Cus, err)
	}
	if n == 1 {
		Newbook.Cus_name = k.Name
		Newbook.Cus_id = k.ID
		Newbook.Cus_email = k.Email
		Newbook.Cus_sdt = k.Sdt

	}

	return Newbook.Cus_name, Newbook.Cus_id, Newbook.Cus_email, Newbook.Cus_sdt
}

func findRoom(Newbook *book, rooms []room, err error) (int, string, int) {

	fmt.Println("Please choose room type.")
	fmt.Print("1. Single    2. Double    3.President?")
	var k int
	fmt.Scanln(&k)
	switch k {
	case 1:
		Newbook.Room_type = "Single"

	case 2:
		Newbook.Room_type = "Double"

	case 3:
		Newbook.Room_type = "President"

	default:
		fmt.Println("Please enter room type in range 1 to 3")
		findRoom(Newbook, rooms, err)
	}
	var r []room
	for _, v := range rooms {
		if v.Type_room == Newbook.Room_type {
			r = append(r, v)
		}
	}
	fmt.Println("This is the list of rooms you want to book")
	for _, v := range r {
		fmt.Printf("Room ID:%d    Type: %s   Price: %d \n", v.Id_room, v.Type_room, v.Price_room)
	}

	fmt.Println("\nPlease enter the room ID that you want to book: ")
	Newbook.Room_id, err = strconv.Atoi(nextLine())
	if err != nil {
		fmt.Println("Please enter number")
		fmt.Println("Please reinsert the customer ID")
		findRoom(Newbook, rooms, err)
	}
	if Newbook.Cus_id < 0 {
		fmt.Println("The input must be positive number")
		fmt.Println("Please reinsert the customer ID ")
		findRoom(Newbook, rooms, err)
	}
	n := 0
	var h room
	for _, v := range r {
		if Newbook.Room_id == v.Id_room {
			h = v
			n++
		}
	}
	if n == 0 {
		fmt.Println("Please reinsert the room ID you want.")

	}
	if n == 1 {
		Newbook.Room_id = h.Id_room
		Newbook.Room_type = h.Type_room
		Newbook.Room_price = h.Price_room
	}

	return Newbook.Room_id, Newbook.Room_type, Newbook.Room_price
}

func sortRoom(rooms []room) {
	var r []room
	var t string
	fmt.Println("Please choose room type.")
	fmt.Print("1. Single    2. Double    3.President?  ")
	var k int
	fmt.Scanln(&k)
	switch k {
	case 1:
		t = "Single"

	case 2:
		t = "Double"

	case 3:
		t = "President"
	default:
		fmt.Println("The input must be between 1 and 3.")
		sortRoom(rooms)
	}
	for _, v := range rooms {
		if v.Type_room == t {
			r = append(r, v)
		}
	}
	fmt.Println("This is the list of rooms you want to sort")
	for _, v := range r {
		fmt.Printf("Room ID:%d    Type: %s   Price: %d \n", v.Id_room, v.Type_room, v.Price_room)
	}

}

func createBill(books []book) {
	total := 0
	var k int

	fmt.Print("Please insert the customer id: ")
	fmt.Scanln(&k)
	for _, v := range books {
		if v.Cus_id == k {
			total += v.Room_price
		}
	}

	fmt.Printf("Total price: %d\n", total)
}
func main() {
	var customers []customer
	var rooms []room
	var books []book
	var err error

	for {
		customers = LoadCus("customer.json")
		rooms = LoadRoom("room.json")
		books = LoadBook("book.json")
		var n int
		fmt.Println("----------------------------------------")
		fmt.Println("MENU")
		fmt.Println("1. Enter new room")
		fmt.Println("2. Enter new customer")
		fmt.Println("3. Book")
		fmt.Println("4. Sort room booking by room type")
		fmt.Println("5. Create bill for customer")
		fmt.Println("6. Out program")
		fmt.Println("Your choice? ")
		fmt.Scanln(&n)

		switch n {
		case 1:

			fmt.Print("Please enter room infomation:\n")
			var Newroom room
			InsertRoomID(&Newroom, err, rooms)

			fmt.Println("Type of room: ")
			InsertRoomType(&Newroom, err)

			fmt.Println("Price of room: ")
			InsertRoomPrice(&Newroom, err)

			fmt.Println("Successful!")
			rooms = append(rooms, Newroom)
			SaveRoom("room.json", rooms)
			fmt.Printf("%10s%10s%10s\n", "ID", "Type", "Price")
			fmt.Printf("%10d%10s%10d\n", Newroom.Id_room, Newroom.Type_room, Newroom.Price_room)

		case 2:

			fmt.Print("Please enter customer infomation:\n")
			var Newcus customer
			InsertCusName(&Newcus)

			InsertCusEmail(&Newcus)

			InsertCusPhone(&Newcus, err)

			InsertCusID(&Newcus, err, customers)

			fmt.Println("Successful!")
			customers = append(customers, Newcus)
			SaveCus("customer.json", customers)

		case 3:
			var Newbook book
			findCustomer(&Newbook, customers, err)
			findRoom(&Newbook, rooms, err)

			fmt.Println("Successful!")
			books = append(books, Newbook)
			SaveBook("book.json", books)

		case 4:
			sortRoom(rooms)

		case 5:
			createBill(books)
		case 6:
			return
		default:
			break
		}
	}

}
