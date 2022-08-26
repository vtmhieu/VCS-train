package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// the single responsibility principle states that a type should have one primary responsibility
// an as a result, it should have one reason to change that reason being somehow related to its primary responsibility

var entriesCount = 0

type Journal struct {
	entries []string
}

func (j *Journal) String() string {
	return strings.Join(j.entries, "\n")
}
func (j *Journal) addEntries(text string) int {
	entriesCount++
	entry := fmt.Sprintf("%d: %s", entriesCount, text)

	j.entries = append(j.entries, entry)

	return entriesCount
}

func (j *Journal) removeEntries(text string) int {

	entriesCount--

	return entriesCount
}

//separation of concerns

func (j *Journal) Save(filename string) error {
	return ioutil.WriteFile(filename, []byte(j.String()), 0644)

}

func (j *Journal) Load(filename string) {

}

func main() {

}
