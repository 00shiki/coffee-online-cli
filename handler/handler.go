package handler

import (
	"coffee-online-cli/repository/orders"
	"coffee-online-cli/repository/products"
	"coffee-online-cli/repository/users"
)

type Handler struct {
	usersRepo    users.Repo
	productsRepo products.Repo
	ordersRepo   orders.Repo
}

func NewHandler(usersRepo users.Repo, productsRepo products.Repo, ordersRepo *orders.Repository) *Handler {
	return &Handler{
		usersRepo:    usersRepo,
		productsRepo: productsRepo,
		ordersRepo:   ordersRepo,
	}
}
