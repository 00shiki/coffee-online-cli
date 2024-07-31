package users

import (
	"coffee-online-cli/entity"
	"database/sql"
	"errors"
	"fmt"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) CreateUser(user entity.User) error {
	query, err := r.db.Prepare("INSERT INTO Users (Name, Email, Password, Location, RoleID) VALUES (?, ?, ?, ?, 1)")
	if err != nil {
		return fmt.Errorf("could not prepare query: %v", err)
	}
	defer query.Close()
	_, err = query.Exec(user.Name, user.Email, user.Password, user.Location)
	if err != nil {
		return fmt.Errorf("could not create user: %v", err)
	}
	return nil
}

func (r *Repository) CheckEmailExists(email string) error {
	query, err := r.db.Prepare("SELECT email FROM Users WHERE Email=?")
	if err != nil {
		return fmt.Errorf("could not prepare query: %v", err)
	}
	defer query.Close()
	rows, err := query.Query(email)
	if err != nil {
		return fmt.Errorf("could not query users: %v", err)
	}
	defer rows.Close()
	if rows.Next() {
		var email string
		err := rows.Scan(&email)
		if err != nil {
			return fmt.Errorf("could not scan row: %v", err)
		}
		if email != "" {
			return errors.New("email already exists")
		}
	}
	return nil
}

func (r *Repository) GetUserByEmail(email string) (*entity.User, error) {
	user := &entity.User{
		Email: email,
	}
	query, err := r.db.Prepare("SELECT UserID, Name, Password FROM Users WHERE Email=?")
	if err != nil {
		return nil, fmt.Errorf("could not prepare query: %v", err)
	}
	rows, err := query.Query(email)
	if err != nil {
		return nil, fmt.Errorf("could not query rows: %v", err)
	}
	defer rows.Close()
	if rows.Next() {
		err := rows.Scan(&user.ID, &user.Name, &user.Password)
		if err != nil {
			return nil, fmt.Errorf("could not scan row: %v", err)
		}
	}
	return user, nil
}

func (r *Repository) GetUserByID(id int) (*entity.User, error) {
	user := &entity.User{
		ID: id,
	}
	query, err := r.db.Prepare(`
	SELECT u.Name, u.Email, u.Location, r.RoleID, r.RoleName
	FROM Users u
			 JOIN Role r ON u.RoleID = r.RoleID
	WHERE u.UserID = ?
	`)
	if err != nil {
		return nil, fmt.Errorf("could not prepare query: %v", err)
	}
	rows, err := query.Query(id)
	if err != nil {
		return nil, fmt.Errorf("could not query rows: %v", err)
	}
	defer rows.Close()
	if !rows.Next() {
		return nil, errors.New("user not found")
	}
	err = rows.Scan(&user.Name, &user.Email, &user.Location, &user.Role.ID, &user.Role.Name)
	if err != nil {
		return nil, fmt.Errorf("could not scan row: %v", err)
	}
	return user, nil
}
