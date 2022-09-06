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
	Name        string `json:"name"`
	Email       string `json:"email"`
	Sdt         int64  `json:"sdt"`
	Code        int64  `json:"code"`
	Type_room   string `json:"type_room"`
	Description string `json:"description"`
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

func InsertRoomID(Newroom *room, err error) int {
	fmt.Print("ID: ")
	Newroom.Id_room, err = strconv.Atoi(nextLine())
	if err != nil {
		fmt.Println("The input must be integer")
		fmt.Println("Please reinsert the room ID ")
		InsertRoomID(Newroom, err)
	}
	if Newroom.Id_room < 0 {
		fmt.Println("The input must be positive number")
		fmt.Println("Please reinsert the room ID ")
		InsertRoomID(Newroom, err)
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
func main() {

	var rooms []room
	var err error
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
			InsertRoomID(&Newroom, err)

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

			break
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
