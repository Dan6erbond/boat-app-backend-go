package models

type User struct {
	Base
	Username  string
	Password  string
	FirstName string
	LastName  string
}
