package handler

import (
	"coffee-online-cli/entity"
	"coffee-online-cli/utils"
	"fmt"
	"golang.org/x/term"
	"log"
	"syscall"
)

func (h *Handler) RegisterUsers() {
	fmt.Print("Masukkan nama: ")
	var name string
	_, err := fmt.Scanln(&name)
	if err != nil {
		log.Fatalf("Failed to read name: %v", err)
		return
	}

	fmt.Print("Masukkan email: ")
	var email string
	_, err = fmt.Scan(&email)
	if err != nil {
		log.Fatalf("Failed to read email: %v", err)
		return
	}
	err = h.usersRepo.CheckEmailExists(email)
	if err != nil {
		log.Fatalf("Failed to check user email: %v", err)
		return
	}

	fmt.Print("Masukkan password: ")
	password, err := term.ReadPassword(syscall.Stdin)
	if err != nil {
		log.Fatalf("Failed to read password: %v", err)
		return
	}
	hashedPassword := utils.HashPassword(password)

	fmt.Print("\nMasukkan lokasi: ")
	var location string
	_, err = fmt.Scanln(&location)
	if err != nil {
		log.Fatalf("Failed to read location: %v", err)
		return
	}

	user := entity.User{
		Name:     name,
		Email:    email,
		Password: hashedPassword,
		Location: location,
	}
	err = h.usersRepo.CreateUser(user)
	if err != nil {
		log.Fatalf("Failed to create user: %v", err)
		return
	}
	fmt.Println("User berhasil dibuat!")
}
