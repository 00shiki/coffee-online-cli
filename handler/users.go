package handler

import (
	"bufio"
	"coffee-online-cli/entity"
	"coffee-online-cli/utils"
	"fmt"
	"log"
	"os"
	"strings"
	"syscall"

	"golang.org/x/term"
)

func (h *Handler) RegisterUsers() {
	fmt.Print("Masukkan nama: ")
	var name string
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')
	name = strings.TrimSpace(name)
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
	password, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		log.Fatalf("Failed to read password: %v", err)
		return
	}
	hashedPassword := utils.HashPassword(password)

	fmt.Print("\nMasukkan lokasi: ")
	var location string
	location, err = reader.ReadString('\n')
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
	loggedUser.Password = user.Password

	switch loggedUser.Role.ID {
	case 1:
		h.CustomerMenu(loggedUser)
	case 2:
		h.AdminMenu()
	}
}

func (h *Handler) ReportLoyal() {
	fmt.Println("Daftar Loyal Customer Coffee Hunter")
	loyals, err := h.usersRepo.LoyalCustomer()
	if err != nil {
		log.Fatalf("error fetching report: %v", err)
		return
	}
	utils.LoyalTable(loyals)
	fmt.Println("\nTekan tombol ENTER untuk melanjutkan ke menu...")
	fmt.Scanf("\n")
}

func (h *Handler) UserUpdate(user *entity.User) {
	newUser := entity.User{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		Location: user.Location,
	}
	reader := bufio.NewReader(os.Stdin)
loop:
	for {
		fmt.Println("Daftar kolom user:")
		fmt.Printf("1. Nama (%s)\n", newUser.Name)
		fmt.Printf("2. Email (%s)\n", newUser.Email)
		fmt.Println("3. Password")
		fmt.Printf("4. Lokasi (%s)\n", newUser.Location)
		fmt.Println("5. Simpan Perubahan")
		fmt.Println("6. Kembali")
		fmt.Print("Masukkan kolom yang ingin diubah: ")
		var column int
		_, err := fmt.Scan(&column)
		if err != nil {
			log.Fatalf("Failed to read column: %v", err)
			return
		}
		switch column {
		case 1:
			fmt.Print("Masukkan nama baru: ")
			newUser.Name, err = reader.ReadString('\n')
			newUser.Name = strings.TrimSpace(newUser.Name)
			if err != nil {
				log.Fatalf("Failed to read name: %v", err)
				return
			}
		case 2:
			fmt.Print("Masukkan email baru: ")
			_, err := fmt.Scan(&newUser.Email)
			if err != nil {
				log.Fatalf("Failed to read email: %v", err)
				return
			}
		case 3:
			fmt.Print("Masukkan password baru: ")
			newPassword, err := term.ReadPassword(int(syscall.Stdin))
			if err != nil {
				log.Fatalf("Failed to read password: %v", err)
				return
			}
			newUser.Password = utils.HashPassword(newPassword)
		case 4:
			fmt.Print("Masukkan lokasi baru: ")
			newUser.Location, err = reader.ReadString('\n')
			newUser.Location = strings.TrimSpace(newUser.Location)
			if err != nil {
				log.Fatalf("Failed to read location: %v", err)
				return
			}
		case 5:
			break loop
		case 6:
			return
		default:
			fmt.Println("Mohon masukkan (1/2/3/4/5/6)")
		}
	}
	err := h.usersRepo.EditUser(newUser)
	if err != nil {
		log.Fatalf("Failed to update user: %v", err)
		return
	}
	user.Name = newUser.Name
	user.Email = newUser.Email
	user.Password = newUser.Password
	user.Location = newUser.Location
	fmt.Println("Perubahan berhasil dibuat")
}

func (h *Handler) CustomerMenu(user *entity.User) {
loop:
	for {
		fmt.Printf("Halo %s, ingin kopi apa hari ini?\n", user.Name)
		fmt.Println("1. Pesan Kopi")
		fmt.Println("2. Melihat Status Pesanan")
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
			h.CoffeeOrders(user)
		case 2:
			h.UserOrders(user)
		case 3:
			h.UserUpdate(user)
		case 4:
			break loop
		default:
			fmt.Println("Mohon masukkan pilihan (1/2/3/4/5/6)...")
		}
	}
}

func (h *Handler) AdminMenu() {
loop:
	for {
		fmt.Println("***ADMIN***")
		fmt.Println("1. Tambahkan Produk")
		fmt.Println("2. Restock Produk")
		fmt.Println("3. Kirim Pesanan")
		fmt.Println("4. Laporan Produk Populer")
		fmt.Println("5. Laporan Stok Produk")
		fmt.Println("6. Laporan Loyal Customer")
		fmt.Println("7. Logout")
		fmt.Print("Masukkan pilihan: ")
		var option int
		_, err := fmt.Scan(&option)
		if err != nil {
			log.Fatalf("Failed to read option: %v", err)
		}
		switch option {
		case 1:
			h.CreateProduct()
		case 2:
			h.ProductRestock()
		case 3:
			h.ShipOrders()
		case 4:
			h.PopularProduct()
		case 5:
			h.ProductStock()
		case 6:
			h.ReportLoyal()
		case 7:
			break loop
		default:
			fmt.Println("Mohon masukkan pilihan (1/2/3/4/5/6)...")
		}
	}
}
