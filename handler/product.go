package handler

import (
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
