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
	if name == "" {
		fmt.Println("Nama tidak boleh kosong")
		return
	}
	fmt.Print("Masukkan stok awal: ")
	var stock int
	_, err = fmt.Scan(&stock)
	if err != nil {
		log.Fatalf("Failed to read product stock: %v", err)
		return
	}
	if stock <= 0 {
		fmt.Println("Stok awal tidak boleh kosong atau negatif")
		return
	}
	fmt.Print("Masukkan harga kopi: ")
	var price float64
	_, err = fmt.Scan(&price)
	if err != nil {
		log.Fatalf("Failed to read product price: %v", err)
		return
	}
	if price <= 0 {
		fmt.Println("Harga tidak boleh kosong atau negatif")
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

func (h *Handler) ProductRestock() {
	for {
		products, err := h.productsRepo.FetchProducts()
		if err != nil {
			log.Fatal(err)
			return
		}
		fmt.Println("Daftar Produk Kopi:")
		for i, product := range products {
			fmt.Printf("%d. %s - Rp %s [%d]\n", i+1, product.Name, utils.PriceFormat(product.Price), product.Stock)
		}
		fmt.Printf("%d. Kembali\n", len(products)+1)
		fmt.Print("Masukkan kopi yang ingin distok ulang: ")
		var index int
		_, err = fmt.Scan(&index)
		if err != nil {
			log.Fatalf("Failed to read products index: %v", err)
			return
		}
		if index == len(products)+1 {
			return
		}
		product := products[index-1]
		fmt.Printf("Masukkan banyaknya stok baru (Stok: %d): ", product.Stock)
		var newStock int
		_, err = fmt.Scan(&newStock)
		if err != nil {
			log.Fatalf("Failed to read new stock: %v", err)
			return
		}
		err = h.productsRepo.ProductStockUpdate(product.ID, product.Stock+newStock)
		if err != nil {
			log.Fatal(err)
			return
		}
	}
}

func (h *Handler) ProductStock() {
		products, err := h.productsRepo.FetchProducts()
		if err != nil {
			log.Fatal(err)
			return
		}
		fmt.Println("Daftar Produk Kopi:")
		for i, product := range products {
			fmt.Printf("%d. %s -  %d\n", i+1, product.Name, product.Stock)
		}
		fmt.Printf("Kembali ke Menu Utama... \n \n")
}