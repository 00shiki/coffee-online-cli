package entity

type User struct {
	UserID		int
	Name		string
	Email		string
	Password	string
	Location	string
	CreatedAt	string
	UpdateAt	string
	Role
}

type Role struct {
	RoleID 		string
	RoleName	string
}