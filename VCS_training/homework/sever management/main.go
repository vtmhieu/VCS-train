package main

import (
	"log"
	"test3/go/pkg/mod/github.com/xuri/excelize@v1.4.1"
)

func connect() {
	f, err := excelize.OpenFile("simple.xlsx")

	if err != nil {
		log.Fatal(err)
	}
}

func main() {

}
