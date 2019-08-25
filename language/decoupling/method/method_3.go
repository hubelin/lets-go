/*
Methods are just syntactic sugar to pretend data has certain behavior. Similar to OOP.package method
*/

package main

import "fmt"

type data struct {
	name string
	age  int
}

func (d data) displayName() {
	fmt.Printf("My name is %s", d.name)
}

func (d *data) setAge(age int) {
	d.age = age
}

func main() {
	d := data{
		name: "Hubert",
	}

	// data.displayName(d)
	d.displayName()
	// (*data).setAge(&d, 22)
	d.setAge(22)
}
