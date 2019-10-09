package main

import "fmt"

type MyInt struct {
	n int
}

func (myInt *MyInt) Increase() {
	myInt.n++
}

func (myInt *MyInt) Decrease() {
	myInt.n--
}

type Pet interface {
	Name(name string) string
	Age() uint8
	ShowPoint()
}

type Dog struct {
	name string
	age  uint8
}

func (dog Dog) Name(name string) string {
	dog.name = name
	return dog.name
}

func (dog Dog) Age() uint8 {
	return dog.age
}

func (dog Dog) ShowPoint() {
	fmt.Printf("%v\n", &(dog.name))
}

func main8() {
	mi := MyInt{}
	mi.Increase()
	mi.Increase()
	mi.Decrease()
	mi.Decrease()
	mi.Increase()
	fmt.Printf("%v\n", mi.n == 1)

	myDog := Dog{"Little D", 3}
	pet, ok2 := interface{}(myDog).(Pet)

	fmt.Printf("%v, %v\n", pet, ok2)

	pet.Name("new name")
	fmt.Printf("%v\n", pet)

	myDog.ShowPoint()
	pet.ShowPoint()

	fmt.Printf("%v, %v\n", &(myDog), &pet)

	ch1 := make(chan, int, 1)
	close
}
