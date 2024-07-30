package handler

import (
	"coffee-online-cli/repository/orders"
	"coffee-online-cli/repository/products"
	"coffee-online-cli/repository/users"
)

type Handler struct {
	usersRepo    *users.Repository
	productsRepo *products.Repository
	ordersRepo   *orders.Repository
}

func NewHandler(usersRepo *users.Repository, productsRepo *products.Repository, ordersRepo *orders.Repository) *Handler {
	return &Handler{
		usersRepo:    usersRepo,
		productsRepo: productsRepo,
		ordersRepo:   ordersRepo,
	}
}
