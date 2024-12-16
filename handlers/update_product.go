package handlers

import (
	"errors"
	"lb4/models"
	"net/http"
	"strconv"
	"text/template"
)

func UpdateProductHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tmpl, _ := template.ParseFiles("templates/layout.html", "templates/update_product.html")
		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			renderError(w, err)
			return
		}
		products, err := models.GetAllProducts()
		if err != nil {
			renderError(w, err)
			return
		}
		var product *models.Product
		for _, p := range products {
			if p.ID == id {
				product = &p
				break
			}
		}
		if product == nil {
			renderError(w, errors.New("unknown product"))
			return
		}
		tmpl.Execute(w, product)
		return
	}

	if r.Method == http.MethodPost {
		productID := r.FormValue("product_id")
		name := r.FormValue("name")
		price := r.FormValue("price")

		err := models.UpdateProduct(productID, name, price)
		if err != nil {
			renderError(w, err)
			return
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
}
