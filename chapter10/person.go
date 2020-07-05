package main

import (
	"fmt"
	"strings"
)

type Person struct {
	fistName string
	lastName string
}

func upPerson(p *Person) {
	p.fistName = strings.ToUpper(p.fistName)
	p.lastName = strings.ToUpper(p.lastName)
}

func main() {
	var pers1 Person
	pers1.fistName = "zhang"
	pers1.lastName = "zhenqiang"
	fmt.Printf("%v\n", pers1)
	upPerson(&pers1)
	fmt.Printf("%v\n", pers1)

	pers2 := new(Person)
	pers2.fistName = "zhang"
	pers2.lastName = "wanxia"
	(*pers2).lastName = "bendan"
	upPerson(pers2)

	fmt.Printf("%v\n", pers2);

	pers3 := &Person{"zhang", "bendan"}
	upPerson(pers3)
	fmt.Printf("%v\n", pers3)
}
