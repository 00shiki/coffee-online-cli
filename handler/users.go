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

func (h *Handler) LoginUsers() {
	fmt.Print("Masukkan email: ")
	var email string
	_, err := fmt.Scan(&email)
	if err != nil {
		log.Fatalf("Failed to read email: %v", err)
		return
	}
	fmt.Print("Masukkan password: ")
	password, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		log.Fatalf("Failed to read password: %v", err)
		return
	}
	fmt.Println()

	user, err := h.usersRepo.GetUserByEmail(email)
	if err != nil {
		log.Fatalf("Failed to get user: %v", err)
		return
	}

	if !utils.ComparePassword([]byte(user.Password), password) {
		log.Fatal("Invalid password")
		return
	}

	fmt.Println("Login berhasil!")

	loggedUser, err := h.usersRepo.GetUserByID(user.ID)
	if err != nil {
		log.Fatalf("Failed to get user: %v", err)
		return
	}

	switch loggedUser.Role.ID {
	case 1:
		h.CustomerMenu(loggedUser)
	case 2:
		h.AdminMenu(loggedUser)
	}
}

func (h *Handler) CustomerMenu(user *entity.User) {
loop:
	for {
		fmt.Printf("Halo %s, ingin kopi apa hari ini?\n", user.Name)
		fmt.Println("1. Pesan Kopi")
		fmt.Println("2. Melihat Pesanan")
		fmt.Println("3. Edit User")
		fmt.Println("4. Logout")
		fmt.Print("Masukkan pilihan: ")
		var option int
		_, err := fmt.Scan(&option)
		if err != nil {
			log.Fatalf("Failed to read option: %v", err)
		}
		switch option {
		case 1:
		case 2:
		case 3:
		case 4:
			break loop
		default:
			fmt.Println("Mohon masukkan pilihan (1/2/3/4/5/6)...")
		}
	}
}

func (h *Handler) AdminMenu(user *entity.User) {
loop:
	for {
		fmt.Println("***ADMIN***")
		fmt.Println("1. Tambahkan Produk")
		fmt.Println("2. Restock Produk")
		fmt.Println("3. Mengubah Status Pengiriman")
		fmt.Println("4. Laporan Produk Populer")
		fmt.Println("5. Laporan Stok Produk")
		fmt.Println("6. Laporan Pengeluaran Customer")
		fmt.Println("7. Logout")
		fmt.Print("Masukkan pilihan: ")
		var option int
		_, err := fmt.Scan(&option)
		if err != nil {
			log.Fatalf("Failed to read option: %v", err)
		}
		switch option {
		case 1:
		case 2:
		case 3:
		case 4:
		case 5:
		case 6:
		case 7:
			break loop
		default:
			fmt.Println("Mohon masukkan pilihan (1/2/3/4/5/6)...")
		}
	}
}
