package handlers

import (
	"html/template"
	"lb4/models"
	"net/http"
)

func CartHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		productID := r.FormValue("product_id")
		quantity := r.FormValue("quantity")

		if err := models.AddToCart(productID, quantity); err != nil {
			renderError(w, err)
			return
		}
		http.Redirect(w, r, "/cart", http.StatusSeeOther)
		return
	}
	cart, err := models.GetCart()
	if err != nil {
		renderError(w, err)
		return
	}
	t := template.Must(template.ParseFiles("templates/layout.html", "templates/cart.html"))
	t.Execute(w, cart)
}
