package main

import (
	"lb4/handlers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlers.CatalogHandler)
	http.HandleFunc("/add-product", handlers.AddProductHandler)
	http.HandleFunc("/update-product", handlers.UpdateProductHandler)
	http.HandleFunc("/delete-product", handlers.DeleteProductHandler)
	http.HandleFunc("/update-cart-item", handlers.UpdateCartItemHandler)
	http.HandleFunc("/delete-cart-item", handlers.DeleteCartItemHandler)
	http.HandleFunc("/cart", handlers.CartHandler)
	http.HandleFunc("/payment", handlers.PaymentHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Println("Server running on :8080...")
	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
