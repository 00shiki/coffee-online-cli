package cli

import (
	"coffee-online-cli/handler"
	"fmt"
	"log"
)

type Cli struct {
	Handler *handler.Handler
}

func New(handler *handler.Handler) *Cli {
	return &Cli{
		Handler: handler,
	}
}

func (c *Cli) Run() {
loop:
	for {
		fmt.Println("Selamat datang di Coffee Line!")
		fmt.Println("1. Login")
		fmt.Println("2. Register")
		fmt.Println("Masukkan pilihan: ")
		var option int
		_, err := fmt.Scan(&option)
		if err != nil {
			log.Fatalf("Failed to read option: %v", err)
			return
		}
		switch option {
		case 1:
		case 2:
			c.Handler.RegisterUsers()
		case 3:
			break loop
		default:
			fmt.Println("Mohon masukkan pilihan (1/2)...")
		}
	}
}
