package entity

type User struct {
	ID       int
	Name     string
	Email    string
	Password string
	Location string
	Role
}

type Role struct {
	ID   int
	Name string
}
