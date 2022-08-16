package main

import "fmt"

func main() {
	//create a map
	//var colors map[string]string
	//same as
	//colors := make(map[string]string)
	// could use map[int]string then below must be colors[12] = "adskd"

	//1st way
	colors := map[string]string{
		"red":   "#ff000", //add ,
		"green": "#28377",
	}

	colors["white"] = "#ffff"
	//white = string in []

	delete(colors, "white")
	fmt.Println(colors)
	printMap(colors)
}

// func printMap(c map[string]string){
// //c= argument name, [string] = type of map
//	for color,hex := range c{
// //color key for this step through the loop, hex= value for this step through the loop
//}
//}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
