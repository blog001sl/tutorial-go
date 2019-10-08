package main

import (
	"fmt"
)

// Person ssssssss
type Person struct {
	Name    string
	Gender  string
	Age     uint8
	Address string
}

type Animal interface {
	Grow()
	Move(addr string) string
}

func (person *Person) Move(address string) string {
	oldAddr := person.Address
	person.Address = address
	return oldAddr

}

func (person *Person) Grow() {
	fmt.Printf("   ")
	fmt.Printf("growing.....")
}

func main6() {
	p := Person{"Robert", "Male", 33, "Beijing"}

	oladAddress := p.Move("San fransisco")
	fmt.Printf("111111 %v, %v \n", p, oladAddress)

	animal, ok := interface{}(&p).(Animal)
	if ok {
		fmt.Printf("%v\n", animal.Move("Shanghai"))
		animal.Grow()
	} else {
		fmt.Printf("transfer to animal failed. \n")
	}

}
