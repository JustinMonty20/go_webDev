package main

import "fmt"

//composite literals in go are go's take on the usual data types in programming languages.

//same way to write a multi line comment like in java.
/* you can create your own types in go. Struct is a similar to an object in an OOP*/

type person struct {
	fname string
	lname string
}

type footballer struct {
	person
	isgod bool
}

func (p person) speak() {
	fmt.Println(p.fname, p.lname, `says "whats goodie broo"`)
}

func (f footballer) speak() {
	fmt.Println(f.fname, f.lname, `says "whats good brah im a pro"`)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {

	p1 := person{
		"Justin",
		"Montgomery",
	}

	f1 := footballer{
		person{
			"James",
			"Rodgriguez",
		}, true,
	}

	saySomething(p1)
	saySomething(f1)

}
