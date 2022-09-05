package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

type Room struct {
	ID    string `json: "room_id"`
	Type  int    `json: "room_type"`
	Price int    `json: "room_price"`
}

// Read from stdin until line break, return string without line break
// Refer from https://stackoverflow.com/questions/36356573/go-lang-scan-doesent-scan-for-next-line
func nextLine() string {
	bio := bufio.NewReader(os.Stdin)
	line, _, _ := bio.ReadLine()
	return string(line)
}

// Read data from file in filePath and return data in []byte
func readFile(filePath string) []byte {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
	}
	return data
}

// Save room to file path
func saveRooms(file string, Rooms []Room) {
	data, err := json.MarshalIndent(Rooms, "", " ")
	if err != nil {
		fmt.Println(err)
	}

	err = ioutil.WriteFile(file, []byte(data), 0666)
	if err != nil {
		fmt.Println(err)
	}
}

// Load room from filePath and return existing room as a slice
func loadRoom(filePath string) []Room {
	var rooms []Room
	data := readFile(filePath)
	_ = json.Unmarshal(data, &rooms)
	return rooms
}

func main() {

	var rooms []Room
	var err error

	rooms = loadRoom("PHONG.json")
	for {
		// Print menu
		fmt.Println("Features:")
		fmt.Println("1. Enter new room")
		fmt.Println("2. Enter new customer")
		fmt.Println("3. Book")
		fmt.Println("4. Sort room booking by room type")
		fmt.Println("5. Create bill for customer")
		fmt.Println("6. Exit")

		fmt.Print("Your choice? ")
		var n int
		fmt.Scanln(&n)

		// Choose and take action base on choice
		switch n {
		case 1: // NewRoom
			fmt.Println("Enter new room:")
			var newRoom Room
			fmt.Print("ID: ")
			newRoom.ID = nextLine()
			fmt.Print("Type: 1. Single 2. Double 3. VIP? ")
			newRoom.Type, err = strconv.Atoi(nextLine())
			if err != nil {
				fmt.Println("Type must be an integer.\n")
				break
			}
			if newRoom.Type < 1 || newRoom.Type > 3 {
				fmt.Println("Type must be an integer in range [1, 3].\n")
				break
			}
			fmt.Print("Price: ")
			newRoom.Price, err = strconv.Atoi(nextLine())
			if err != nil {
				fmt.Println("Price must be an integer.\n")
				break
			}
			fmt.Println("Successful!")
			rooms = append(rooms, newRoom)
			saveRooms("PHONG.json", rooms)

			fmt.Printf("%10s%10s%10s\n", "ID", "Type", "Price")
			fmt.Printf("%10s%10d%10d\n", newRoom.ID, newRoom.Type, newRoom.Price)
		case 6:
			return
		default:
			fmt.Println("Invalid.\n")
		}
	}
}
