package main

import (
	"log"

	"github.com/xuri/excelize/v2"
)

func connect(filename string) {
	f, err := excelize.OpenFile(filename)

	if err != nil {
		log.Fatal(err)
		return
	}
}

func main() {

}
