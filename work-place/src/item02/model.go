package item02

import "fmt"

type User struct {

	id int
	name string
	address string
	email string
}

func UserFactory(id int, name string, address string, email string) User {

	return User{
		id:      id,
		name:    name,
		address: address,
		email:   email,
	}
}

func (user *User) GetInfo() {

	fmt.Printf("%v\t%v\t%v\t%v\t\n",user.id,user.name,user.address,user.email)
}

func (user *User) GetName() string{

	return user.name
}
func (user *User) GetId() int{

	return user.id
}
func (user *User) GetEmail() string{

	return user.email
}

func (user *User) GetAddress() string{

	return user.address
}