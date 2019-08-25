/*
Everything in Go is passed by value

Pointers share value across program boundaries.package pointer
For example - function calls or goroutines

When program starts, runtime creates a goroutine.
Every goroutine is a separate path of execution.
Can think of it as a lightweight thread.

Every goroutine is given a stack of 2K that can change.
Stack grows "downward"

Every func given a stack frame. Size of stack frame known
at compile time. No value can be placed unless compiler
knows its size ahead of time. If we don't know the size of something at compile time, it has to be on the heap

Zero value enables us to initialize every stack frame we take.
*/

package main

import "fmt"

//user defined type of user
type user struct {
	name  string
	email string
}

func main() {
	// PBV

	count := 10
	increment1(count)

	// Pass address of count by value
	increment2(&count)

	// Escape Analysis
	escapeToHeap()
	fmt.Println(count)
}

func increment1(inc int) {
	inc++
}

func increment2(inc *int) {
	// increment the value of count that pointer inc points to
	*inc++
}

// return value cannot be placed inside stack frame of main s it will escape to heap. main will have a pointer to the heap
func escapeToHeap() *user {
	u := user{
		name:  "Hubert",
		email: "test@gmail.com",
	}
	return &u
}
