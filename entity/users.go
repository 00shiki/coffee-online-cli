package entity

type User struct {
	ID			int
	Name		string
	Email		string
	Password	string
	Location	string
	CreatedAt	string
	UpdateAt	string
	Role
}

type Role struct {
	RoleID 		int
	RoleName	string
}