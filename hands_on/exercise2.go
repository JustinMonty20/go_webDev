package hands_on

import "fmt"

type Person struct {
	Fname, Lname string
}

type SecretAgent struct {
	Person
	Licensetokill bool
	Drinks        int
}

type Human interface {
	Speak()
}

func SayIt(h Human) {
	h.Speak()
}

func (p Person) Speak() {
	fmt.Println(p.Fname, p.Lname, `says "Stoopid"`)
}

func (sa SecretAgent) Speak() {
	fmt.Println(sa.Fname, sa.Lname, `says "Gimmie dat shit"`)
}
