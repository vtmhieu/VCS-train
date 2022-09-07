package main

import (
	"fmt"
	"strings"
)

const (
	indentSize = 2
)

type HtmlElement struct {
	name, text string
	elements   []HtmlElement
}

func main() {
	hello := "hello"
	sb := strings.Builder{} // create a paragraph
	sb.WriteString("<p>")
	sb.WriteString(hello)
	sb.WriteString("</p>")
	fmt.Println(sb.String())

	st := []string{"hello", "world"}

	sb.Reset()
	sb.WriteString("<ul")
	for _, v := range st {
		sb.WriteString("<li>")
		sb.WriteString(v)
		sb.WriteString("</li>")

	}
	sb.WriteString("</ul>")
	fmt.Println(sb.String())
	// to minimize the function we could put it into structure which are more
	// flexible to use
	// this is why we have builder at the first place, the idea is that you have some sort of object that you want to build up in steps

}
