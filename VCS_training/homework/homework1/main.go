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

func main() {
	var customers []customer
	var rooms []room
	var err error
	customers = LoadCus("customer.json")
	rooms = LoadRoom("room.json")
	for {
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

		case 4:
		case 5:
		case 6:
			return
		default:
			break
		}
	}

}
