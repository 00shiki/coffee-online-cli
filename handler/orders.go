package handler

import (
	"coffee-online-cli/entity"
	"coffee-online-cli/utils"
	"fmt"
	"log"
)

func (h *Handler) CoffeeOrders(user *entity.User) {
	var cart []entity.OrderProduct
	var totalAmount float64
loop:
	for {
		fmt.Println("Menu kopi hari ini:")
		products, err := h.productsRepo.FetchProducts()
		if err != nil {
			log.Fatal(err)
			return
		}
		for i, product := range products {
			fmt.Printf("%d. %s - Rp %s [%d]\n", i+1, product.Name, utils.PriceFormat(product.Price), product.Stock)
		}
		fmt.Printf("%d. Kembali\n", len(products)+1)
		fmt.Print("Masukkan kopi yang ingin dibeli: ")
		var index int
		_, err = fmt.Scan(&index)
		if err != nil {
			log.Fatalf("Failed to read products index: %v", err)
			return
		}
		if index == len(products)+1 {
			break
		}
		product := products[index-1]
		fmt.Printf("Masukkan banyaknya kopi yang ingin dibeli (Stok: %d): ", product.Stock)
		var quantity int
		_, err = fmt.Scan(&quantity)
		if err != nil {
			log.Fatalf("Failed to read quantity: %v", err)
			return
		}
		cartIndex := utils.CheckProductExist(cart, product)
		if cartIndex != -1 {
			quantity += cart[cartIndex].Quantity
		}
		if quantity > product.Stock {
			log.Fatal("Quantity cannot be greater than Stock")
			return
		}
		var orderProduct entity.OrderProduct
		if cartIndex != -1 {
			orderProduct = cart[cartIndex]
			orderProduct.Quantity = quantity
		} else {
			orderProduct.Product = product
			orderProduct.Quantity = quantity
		}
		cart = append(cart, orderProduct)
		totalAmount += orderProduct.Product.Price * float64(orderProduct.Quantity)
		fmt.Println("List Pesanan: ")
		for i, cart := range cart {
			fmt.Printf("%d. %s [%d] - Rp %s\n", i+1, cart.Product.Name, cart.Quantity, utils.PriceFormat(cart.Product.Price))
		}
		fmt.Printf("Total Pesanan: Rp %s\n", utils.PriceFormat(totalAmount))
		fmt.Print("Ingin menambah pesanan? (y/n): ")
		var cnt string
		_, err = fmt.Scan(&cnt)
		if err != nil {
			log.Fatalf("Failed to read prompt: %v", err)
			return
		}
		switch cnt {
		case "y":
			continue loop
		case "n":
		default:
			fmt.Println("Mohon masukkan (y/n)")
		}
		fmt.Println("List Pesanan: ")
		for i, cart := range cart {
			fmt.Printf("%d. %s [%d] - Rp %s\n", i+1, cart.Product.Name, cart.Quantity, utils.PriceFormat(cart.Product.Price))
		}
		fmt.Printf("Total Pesanan: Rp %s\n", utils.PriceFormat(totalAmount))
		fmt.Println("Ongkos Kirim: Rp 9.000,00")
		fmt.Println("Total Pembayaran: Rp " + utils.PriceFormat(totalAmount+9000))
		fmt.Print("Lanjutkan ke pembayaran? (y/n): ")
		var pay string
		_, err = fmt.Scan(&pay)
		if err != nil {
			log.Fatalf("Failed to read prompt: %v", err)
			return
		}
		switch pay {
		case "y":
			break loop
		case "n":
			continue loop
		default:
			fmt.Println("Mohon masukkan (y/n)")
		}
	}
	order := &entity.Order{
		OrderProduct: cart,
		User:         *user,
	}
	err := h.ordersRepo.OrderPayment(order)
	if err != nil {
		log.Fatal(err)
		return
	}
	err = h.ordersRepo.CreateOrder(order)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("Pesanan telah dibuat. Pesanan akan segera datang. Selamat menikmati kopi anda!")
}
