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
	name        string `json:"name"`
	email       string `json:"email"`
	sdt         int64  `json:"sdt"`
	code        int64  `json:"code"`
	type_room   string `json:"type_room"`
	description string `json:"description"`
}

// type customers struct {
// 	customers []customer `json:"customers"`
// }

type room struct {
	id_room    int    `json:"room_id"`
	type_room  string `json:"room_type"`
	price_room int    `json:"room_price"`
}

func nextLine() string {
	bio := bufio.NewReader(os.Stdin)
	line, _, _ := bio.ReadLine()
	return string(line)
}
func ReadFile(filename string) []byte {
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

	err = ioutil.WriteFile(filename, []byte(data), 0666)
	if err != nil {
		fmt.Println(err)
	}

}

func LoadRoom(filename string) []room {

	var rooms []room
	data := ReadFile(filename)
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

	data := ReadFile(filename)
	_ = json.Unmarshal(data, &cus)
	return cus
}

func InsertRoomID(Newroom *room, err error) int {
	fmt.Print("ID: ")
	Newroom.id_room, err = strconv.Atoi(nextLine())
	if err != nil {
		fmt.Println("The input must be integer")
		fmt.Println("Please reinsert the room ID ")
		InsertRoomID(Newroom, err)
	}
	if Newroom.id_room < 0 {
		fmt.Println("The input must be positive number")
		fmt.Println("Please reinsert the room ID ")
		InsertRoomID(Newroom, err)
	}
	return Newroom.id_room
}

func InsertRoomType(Newroom *room, err error) string {
	fmt.Print("1. Single  2. Double  3. President ? ")
	var k int
	fmt.Scanln(&k)

	switch k {
	case 1:
		Newroom.type_room = "Single"

	case 2:
		Newroom.type_room = "Double"

	case 3:
		Newroom.type_room = "President"

	default:
		fmt.Println("Please enter room type in range 1 to 3")
		InsertRoomType(Newroom, err)
	}
	return Newroom.type_room
}
func main() {

	var rooms []room
	var err error
	rooms = LoadRoom("Phong.json")
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

			fmt.Print("Pleas enter room infomation:\n")
			var Room room
			InsertRoomID(&Room, err)

			fmt.Println("Type of room: ")
			InsertRoomType(&Room, err)

			fmt.Println("Price of room: ")
			Room.price_room, err = strconv.Atoi(nextLine())
			if err != nil {
				fmt.Println("Input must be number")
				break
			}
			if Room.price_room < 0 {
				fmt.Println("Input must be positive number")
				break
			}

			rooms = append(rooms, Room)
			SaveRoom("Phong.json", rooms)
			fmt.Printf("%10s%10s%10s\n", "ID", "Type", "Price")
			fmt.Printf("%10d%10s%10d\n", *&Room.id_room, *&Room.type_room, *&Room.price_room)

		case 2:

			break
		case 3:
		case 4:
		case 5:
		case 6:
			return
		default:

		}
	}

}
