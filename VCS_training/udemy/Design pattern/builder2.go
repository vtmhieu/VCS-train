package main

import (
	"fmt"
	"strings"
)

type Html struct {
	name, text string
	//name = <p>, type of paragraph (TAG)
	//text = is what go in between tags
	elements []Html
	//in an HTML tag can have many other elements tag in it

}

const (
	indentSize = 2
	//thụt lề
)

func (e *Html) String() string {
	return e.string(0)
}

func (e *Html) string(indent int) string {
	sb := strings.Builder{}
	i := strings.Repeat(" ", indentSize*indent)
	sb.WriteString(fmt.Sprintf("%s</%s>\n", i, e.name))
	if len(e.text) > 0 {
		sb.WriteString(strings.Repeat(" ", indentSize*(indent+1)))
		sb.WriteString(e.text)
		sb.WriteString("\n")
	}

	for _, el := range e.elements {
		sb.WriteString(el.string(indent + 1))
	}
	sb.WriteString(fmt.Sprintf("%s</%s>\n", i, e.name))
	return sb.String()
}

type HtmlBuilder struct {
	rootName string
	root     Html
}

func NewHtmlBuilder(rootName string) *HtmlBuilder {

	return &HtmlBuilder{rootName,
		Html{rootName, "", []Html{}}}
}

func (b *HtmlBuilder) String() string {
	return b.root.String()
}

func (b *HtmlBuilder) Addchild(childName, childText string) {
	e := Html{childName, childText, []Html{}}
	b.root.elements = append(b.root.elements, e)
}
func main() {
	//change a text into a paragraph
	hello := "hello"
	sb := strings.Builder{}
	sb.WriteString("<p>")
	sb.WriteString(hello)
	sb.WriteString("</p>")
	fmt.Println(sb.String())

	//suppose have a list of words and you want to merge them as a list
	words := []string{"Hello", "world"}
	sb.Reset()
	sb.WriteString("<ul>")
	for _, v := range words {
		sb.WriteString("<li>")
		sb.WriteString(v)
		sb.WriteString("</li>")

	}
	sb.WriteString("</ul>")
	fmt.Println(sb.String())

	b := NewHtmlBuilder("ul")
	b.Addchild("li", "Hello")
	b.Addchild("li", "world")
	fmt.Println(b.String())
}
