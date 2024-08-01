package handler

import (
	"coffee-online-cli/entity"
	"coffee-online-cli/utils"
	"fmt"
	"log"
)

func (h *Handler) PopularProduct() {
	fmt.Println("Daftar Produk Populer Coffe Hunter")
	popular, err := h.productsRepo.PopularProduct()
	if err != nil {
		log.Fatalf("error fetching report: %v", err)
		return
	}
	utils.PopularProductTable(popular)
	fmt.Println("\nTekan tombol ENTER untuk melanjutkan ke menu...")
	fmt.Scanf("\n")
}

func (h *Handler) CreateProduct() {
	fmt.Print("Masukkan nama kopi: ")
	var name string
	_, err := fmt.Scanln(&name)
	if err != nil {
		log.Fatalf("Failed to read product name: %v", err)
		return
	}
	fmt.Print("Masukkan stok awal: ")
	var stock int
	_, err = fmt.Scan(&stock)
	if err != nil {
		log.Fatalf("Failed to read product stock: %v", err)
		return
	}
	fmt.Print("Masukkan harga kopi: ")
	var price float64
	_, err = fmt.Scan(&price)
	if err != nil {
		log.Fatalf("Failed to read product price: %v", err)
		return
	}
	product := entity.Product{
		Name:  name,
		Stock: stock,
		Price: price,
	}
	err = h.productsRepo.CreateProduct(product)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("Menu kopi baru berhasil dibuat!")
}
