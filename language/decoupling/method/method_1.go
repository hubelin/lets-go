package main

import "fmt"

// user defines a user in the program

type user struct {
	name  string
	email string
}

//method notify with value receiver
// operates on it's on copy of u
func (u user) notify() {
	fmt.Printf("Sending user email to %s", u.name)
}

//method changeEmail with pointer receiver
// method acts on shared access of u
func (u *user) changeEmail() {
	u.email = "test@gmail.co"
}

func main() {
	// Methods implemented with value receivers can be used on value and pointer types
	// &user.notify() will be turned to (*user),notify()
	// Go will copy the value where &user points to and call notify with it
	users := []*user{
		{"hubert", "hi@gmail.com"},
	}

	for _, u := range users {
		u.changeEmail()
		fmt.Println(*u)
	}
	fmt.Println(users)
}
