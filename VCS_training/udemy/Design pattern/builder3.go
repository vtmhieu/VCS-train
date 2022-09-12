package main

import "fmt"

type Person struct {
	StreetAddress, Postcode, City string

	CompanyName, Position string
	AnnualIncome          int
}

type PersonBuilder struct {
	person *Person
}

func NewpersonBuilder() *PersonBuilder {
	return &PersonBuilder{&Person{}}

}

type PersonAddressBuilder struct {
	PersonBuilder
}

type PersonJobBuilder struct {
	PersonBuilder
}

func (b *PersonBuilder) Live() *PersonAddressBuilder {
	return &PersonAddressBuilder{*b}
}

func (b *PersonBuilder) Work() *PersonJobBuilder {
	return &PersonJobBuilder{*b}
}

func (it *PersonAddressBuilder) Atstr(StreetAddress string) *PersonAddressBuilder {
	it.person.StreetAddress = StreetAddress
	return it
}

func (it *PersonAddressBuilder) In(city string) *PersonAddressBuilder {
	it.person.City = city
	return it
}

func (it *PersonAddressBuilder) With(postcode string) *PersonAddressBuilder {
	it.person.Postcode = postcode
	return it
}

func (it *PersonJobBuilder) As(Position string) *PersonJobBuilder {
	it.person.Position = Position
	return it
}

func (it *PersonJobBuilder) Atcomo(company string) *PersonJobBuilder {
	it.person.CompanyName = company
	return it
}

func (it *PersonJobBuilder) Sal(salary int) *PersonJobBuilder {
	it.person.AnnualIncome = salary
	return it
}

func (b *PersonBuilder) Build() *Person {
	return b.person
}
func main() {
	pb := NewpersonBuilder()

	pb.
		Live().
		Atstr("Hoang Quoc Viet").
		In("Hanoi").
		With("0012312j").
		Work().
		As("Software Engineer").
		Atcomo("Viettel cybersecurity").
		Sal(3000000)
	person := pb.Build()
	fmt.Println(person)
}

// BUILDER FACETS
