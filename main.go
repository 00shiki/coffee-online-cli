package main

import (
	"coffee-online-cli/cli"
	"coffee-online-cli/config"
	"coffee-online-cli/handler"
	ORDERS_REPO "coffee-online-cli/repository/orders"
	PRODUCTS_REPO "coffee-online-cli/repository/products"
	USERS_REPO "coffee-online-cli/repository/users"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}

	db, err := sql.Open("mysql", config.DatabaseConfig())
	if err != nil {
		log.Fatal(err)
		return
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
		return
	}

	usersRepo := USERS_REPO.NewRepository(db)
	productsRepo := PRODUCTS_REPO.NewRepository(db)
	ordersRepo := ORDERS_REPO.NewRepository(db)

	h := handler.NewHandler(usersRepo, productsRepo, ordersRepo)
	c := cli.New(h)
	c.Run()
}
