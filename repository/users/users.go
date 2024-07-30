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
