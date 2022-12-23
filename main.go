package main

import "fmt"

type person struct {
	first string
}

func (p person) speak() {
	fmt.Println("from a person - the name is, ", p.first)
}

type secreteAgent struct {
	person
	ltk bool
}

func (sa secreteAgent) speak() {
	fmt.Println("i am a secrete agend, do i have the license to kill? ", sa.ltk)
}

// person type and secreteAgent type both are concrete types
// human is an abstract type
type human interface {
	speak()
}

func main() {

	p1 := person{
		first: "Mr. KKK",
	}

	sa1 := secreteAgent{
		person: person{
			first: "James",
		},
		ltk: true,
	}

	var x, y human
	x = p1
	y = sa1

	x.speak()
	y.speak()

}
