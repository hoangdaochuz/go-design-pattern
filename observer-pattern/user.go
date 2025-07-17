package observerpattern

import "fmt"

type User struct {
	name string
	id   string
}

func NewUser(name, id string) *User {
	return &User{
		name: name,
		id:   id,
	}
}

func (u *User) Update(msg string) {
	fmt.Println(u.name + " : " + msg)
}

func (u *User) GetId() string {
	return u.id
}
